package master

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.com/upn-belajar-go/infras"
	"gitlab.com/upn-belajar-go/shared/logger"
)

var (
	kelasSiswaQuery = struct {
		Insert string
	}{
		Insert: `INSERT INTO kelas_siswa (id, id_kelas, tahun_ajaran, created_by, created_at) values(:id, :id_kelas, :tahun_ajaran, :created_by, :created_at) `,
	}
)

var (
	kelasSiswaDetailQuery = struct {
		InsertBulk            string
		InsertBulkPlaceholder string
		Exist                 string
	}{
		InsertBulk:            `INSERT INTO public.kelas_siswa_detail(id, id_kelas_siswa, id_siswa, created_at, created_by) values `,
		InsertBulkPlaceholder: ` (:id, :id_kelas_siswa, :id_siswa, :created_at, :created_by) `,
		Exist: ` select count(id_siswa)>0 from kelas_siswa_detail s
			left join kelas_siswa ks on ks.id = s.id_kelas_siswa
		`,
	}
)

type KelasSiswaRepository interface {
	Create(data KelasSiswa) error
	ExistByIdSiswa(idSiswa string, idKelasSiswa string) (bool, error)
}

type KelasSiswaRepositoryPostgreSQL struct {
	DB *infras.PostgresqlConn
}

func ProvideKelasSiswaRepositoryPostgreSQL(db *infras.PostgresqlConn) *KelasSiswaRepositoryPostgreSQL {
	s := new(KelasSiswaRepositoryPostgreSQL)
	s.DB = db
	return s
}

func (r *KelasSiswaRepositoryPostgreSQL) Create(data KelasSiswa) error {
	return r.DB.WithTransaction(func(tx *sqlx.Tx, e chan error) {
		if err := r.CreateTxKelasSiswa(tx, data); err != nil {
			e <- err
			return
		}

		if err := txCreateKelasSiswaDetail(tx, data.Detail); err != nil {
			e <- err
			return
		}
		e <- nil
	})
}

func (r *KelasSiswaRepositoryPostgreSQL) CreateTxKelasSiswa(tx *sqlx.Tx, data KelasSiswa) error {
	stmt, err := tx.PrepareNamed(kelasSiswaQuery.Insert)
	if err != nil {
		logger.ErrorWithStack(err)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(data)
	if err != nil {
		logger.ErrorWithStack(err)
		return err
	}
	return nil
}

func txCreateKelasSiswaDetail(tx *sqlx.Tx, details []KelasSiswaDetail) (err error) {
	if len(details) == 0 {
		return
	}
	query, args, err := composeBulkUpsertKelasSiswaDetailQuery(details)
	if err != nil {
		return
	}

	query = tx.Rebind(query)
	stmt, err := tx.Preparex(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Stmt.Exec(args...)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	return
}

func composeBulkUpsertKelasSiswaDetailQuery(details []KelasSiswaDetail) (qResult string, params []interface{}, err error) {
	values := []string{}
	for _, d := range details {
		param := map[string]interface{}{
			"id":             d.ID,
			"id_kelas_siswa": d.IdKelasSiswa,
			"id_siswa":       d.IdSiswa,
			"created_at":     d.CreatedAt,
			"created_by":     d.CreatedBy,
		}
		q, args, err := sqlx.Named(kelasSiswaDetailQuery.InsertBulkPlaceholder, param)
		if err != nil {
			return qResult, params, err
		}
		values = append(values, q)
		params = append(params, args...)
	}
	qResult = fmt.Sprintf(`%v %v 
						ON CONFLICT (id) 
						DO UPDATE SET id_kelas_siswa=EXCLUDED.id_kelas_siswa, id_siswa=EXCLUDED.id_siswa `, kelasSiswaDetailQuery.InsertBulk, strings.Join(values, ","))

	fmt.Println("tes", qResult)
	return
}

func (r *KelasSiswaRepositoryPostgreSQL) ExistByIdSiswa(idSiswa string, idKelasSiswa string) (bool, error) {
	var exist bool
	err := r.DB.Read.Get(&exist, kelasSiswaDetailQuery.Exist+" where coalesce(s.is_deleted) = false and s.id_siswa = $1 and ks.id = $2 ", idSiswa, idKelasSiswa)
	if err != nil {
		logger.ErrorWithStack(err)
	}
	return exist, err
}
