package master

import (
	"errors"

	"github.com/rs/zerolog/log"
	"gitlab.com/upn-belajar-go/configs"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/shared/pagination"
)

type SiswaService interface {
	Create(reqFormat RequestSiswaFormat, userID string) (newSiswa Siswa, err error)
	GetAllData() (data []Siswa, err error)
	ResolveAll(req model.StandardRequest) (data pagination.Response, err error)
	ResolveByID(id string) (data Siswa, err error)
	Update(reqFormat RequestSiswaFormat, userID string) (data Siswa, err error)
	DeleteByID(id string, userID string) error
}

type SiswaServiceImpl struct {
	SiswaRepository SiswaRepository
	Config          *configs.Config
}

func ProvideSiswaServiceImpl(repository SiswaRepository, config *configs.Config) *SiswaServiceImpl {
	s := new(SiswaServiceImpl)
	s.SiswaRepository = repository
	s.Config = config
	return s
}

func (s *SiswaServiceImpl) Create(reqFormat RequestSiswaFormat, userID string) (newSiswa Siswa, err error) {
	newSiswa, _ = newSiswa.NewSiswaFormat(reqFormat, userID)
	err = s.SiswaRepository.Create(newSiswa)
	if err != nil {
		return Siswa{}, err
	}
	return newSiswa, nil
}

func (s *SiswaServiceImpl) GetAllData() (data []Siswa, err error) {
	return s.SiswaRepository.GetAllData()
}

func (s *SiswaServiceImpl) ResolveAll(req model.StandardRequest) (data pagination.Response, err error) {
	return s.SiswaRepository.ResolveAll(req)
}

func (s *SiswaServiceImpl) ResolveByID(id string) (data Siswa, err error) {
	return s.SiswaRepository.ResolveByID(id)
}

func (s *SiswaServiceImpl) Update(reqFormat RequestSiswaFormat, userID string) (data Siswa, err error) {
	siswa, _ := data.NewSiswaFormat(reqFormat, userID)
	err = s.SiswaRepository.Update(siswa)
	if err != nil {
		log.Error().Msgf("service.UpdateSiswa error", err)
	}
	return siswa, nil
}

func (s *SiswaServiceImpl) DeleteByID(id string, userID string) error {
	siswa, err := s.SiswaRepository.ResolveByID(id)

	if err != nil || (Siswa{}) == siswa {
		return errors.New("Data siswa dengan ID :" + id + " tidak ditemukan")
	}

	siswa.SoftDelete(userID)
	err = s.SiswaRepository.Update(siswa)
	if err != nil {
		return errors.New("Ada kesalahan dalam menghapus data siswa dengan ID: " + id)
	}
	return nil
}
