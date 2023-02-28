package master

import (
	"time"

	"github.com/gofrs/uuid"
	"gitlab.com/upn-belajar-go/shared"
)

type Siswa struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	Nama      string     `db:"nama" json:"nama" validate:"required"`
	Kelas     string     `db:"kelas" json:"kelas"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	CreatedBy *string    `db:"created_by" json:"createdBy"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	UpdatedBy *string    `db:"updated_by" json:"updatedBy"`
	IsDeleted bool       `db:"is_deleted" json:"isDeleted"`
}

type RequestSiswaFormat struct {
	ID    string `db:"id" json:"id"`
	Nama  string `db:"nama" json:"nama" validate:"required"`
	Kelas string `db:"kelas" json:"kelas" validate:"required"`
}

// Validate validates the entity.
func (f *Siswa) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(f)
}

func (s *Siswa) NewSiswaFormat(reqFormat RequestSiswaFormat, userID string) (newSiswa Siswa, err error) {
	newID, _ := uuid.NewV4()
	now := time.Now()

	if reqFormat.ID == "" {
		newSiswa = Siswa{
			ID:        newID,
			Nama:      reqFormat.Nama,
			Kelas:     reqFormat.Kelas,
			CreatedAt: time.Now(),
			CreatedBy: &userID,
		}
	} else {
		id, _ := uuid.FromString(reqFormat.ID)
		newSiswa = Siswa{
			ID:        id,
			Nama:      reqFormat.Nama,
			Kelas:     reqFormat.Kelas,
			UpdatedAt: &now,
			UpdatedBy: &userID,
		}
	}
	err = newSiswa.Validate()
	return
}

var ColumnMappSiswa = map[string]interface{}{
	"id":        "id",
	"nama":      "nama",
	"kelas":     "kelas",
	"createdBy": "created_by",
	"createdAt": "created_at",
	"updatedBy": "updated_by",
	"updatedAt": "updated_at",
	"isDeleted": "is_deleted",
}

func (siswa *Siswa) SoftDelete(userId string) {
	now := time.Now()
	siswa.IsDeleted = true
	siswa.UpdatedBy = &userId
	siswa.UpdatedAt = &now
}
