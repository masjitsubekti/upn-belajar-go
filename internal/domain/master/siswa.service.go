package master

import (
	"gitlab.com/upn-belajar-go/configs"
)

type SiswaService interface {
	Create(reqFormat RequestSiswaFormat, userID string) (newSiswa Siswa, err error)
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
