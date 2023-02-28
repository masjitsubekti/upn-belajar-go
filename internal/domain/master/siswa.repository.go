package master

import (
	"gitlab.com/upn-belajar-go/infras"
	"gitlab.com/upn-belajar-go/shared/logger"
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
				id=:id, 
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
