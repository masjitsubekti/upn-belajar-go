package master

import (
	"bytes"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.com/upn-belajar-go/infras"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/shared/logger"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/shared/pagination"
)

var (
	siswaQuery = struct {
		Select string
		Insert string
		Update string
		Delete string
		Exist  string
		Count  string
	}{
		Select: `select id, nama, kelas, updated_at, created_at, created_by, updated_by, is_deleted from m_siswa `,
		Insert: `INSERT INTO m_siswa (id, nama, kelas, created_by, created_at) values(:id, :nama, :kelas, :created_by, :created_at) `,
		Update: `UPDATE m_siswa SET 
				
				nama=:nama,
				kelas=:kelas,
				updated_at=:updated_at,
				updated_by=:updated_by, 
				is_deleted=:is_deleted`,
		Delete: `delete from m_siswa `,
		Exist:  `select count(id)>0 from m_siswa `,
		Count:  `select count(id) from m_siswa `,
	}
)

type SiswaRepository interface {
	Create(data Siswa) error
	GetAllData() (data []Siswa, errr error)
	ResolveAll(req model.StandardRequest) (data pagination.Response, err error)
	ResolveByID(id string) (data Siswa, err error)
	Update(data Siswa) error
}

type SiswaRepositoryPostgreSQL struct {
	DB *infras.PostgresqlConn
}

func ProvideSiswaRepositoryPostgreSQL(db *infras.PostgresqlConn) *SiswaRepositoryPostgreSQL {
	s := new(SiswaRepositoryPostgreSQL)
	s.DB = db
	return s
}

func (r *SiswaRepositoryPostgreSQL) Create(data Siswa) error {
	stmt, err := r.DB.Read.PrepareNamed(siswaQuery.Insert)
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

func (r *SiswaRepositoryPostgreSQL) Update(data Siswa) error {
	return r.DB.WithTransaction(func(tx *sqlx.Tx, e chan error) {
		if err := txUpdateSiswa(tx, data); err != nil {
			e <- err
			return
		}
		e <- nil
	})
}

func txUpdateSiswa(tx *sqlx.Tx, data Siswa) (err error) {
	stmt, err := tx.PrepareNamed(siswaQuery.Update + " WHERE id=:id")
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(data)
	if err != nil {
		logger.ErrorWithStack(err)
	}
	return
}

func (r *SiswaRepositoryPostgreSQL) GetAllData() (data []Siswa, errr error) {
	rows, err := r.DB.Read.Queryx(siswaQuery.Select + " where is_deleted = false ")
	if err == sql.ErrNoRows {
		_ = failure.NotFound("Siswa Not Found")
		return
	}

	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	for rows.Next() {
		var master Siswa
		err = rows.StructScan(&master)

		if err != nil {
			return
		}

		data = append(data, master)
	}
	return
}

func (r *SiswaRepositoryPostgreSQL) ResolveAll(req model.StandardRequest) (data pagination.Response, err error) {
	var searchParams []interface{}
	var searchRoleBuff bytes.Buffer
	searchRoleBuff.WriteString(" WHERE coalesce(is_deleted, false) = false ")

	if req.Keyword != "" {
		searchRoleBuff.WriteString(" AND ")
		searchRoleBuff.WriteString(" concat(nama, kelas) ilike ? ")
		searchParams = append(searchParams, "%"+req.Keyword+"%")
	}

	query := r.DB.Read.Rebind(siswaQuery.Count + searchRoleBuff.String())
	fmt.Println("query", searchRoleBuff.String())

	var totalData int
	err = r.DB.Read.QueryRow(query, searchParams...).Scan(&totalData)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	fmt.Println(totalData)

	if totalData < 1 {
		data.Items = make([]interface{}, 0)
		return
	}

	searchRoleBuff.WriteString("order by " + ColumnMappSiswa[req.SortBy].(string) + " " + req.SortType + " ")

	offset := (req.PageNumber - 1) * req.PageSize
	searchRoleBuff.WriteString("limit ? offset ? ")
	searchParams = append(searchParams, req.PageSize)
	searchParams = append(searchParams, offset)

	searchSiswaQuery := searchRoleBuff.String()
	searchSiswaQuery = r.DB.Read.Rebind(siswaQuery.Select + searchSiswaQuery)
	fmt.Println("query", searchSiswaQuery)
	rows, err := r.DB.Read.Queryx(searchSiswaQuery, searchParams...)
	if err != nil {
		return
	}
	for rows.Next() {
		var siswa Siswa
		err = rows.StructScan(&siswa)
		if err != nil {
			return
		}

		data.Items = append(data.Items, siswa)
	}

	data.Meta = pagination.CreateMeta(totalData, req.PageSize, req.PageNumber)

	return
}

func (r *SiswaRepositoryPostgreSQL) ResolveByID(id string) (data Siswa, err error) {
	err = r.DB.Read.Get(&data, siswaQuery.Select+" WHERE id=$1 ", id)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	return
}
