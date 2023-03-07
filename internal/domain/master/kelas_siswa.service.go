package master

import (
	"gitlab.com/upn-belajar-go/configs"
)

type KelasSiswaService interface {
	Create(reqFormat KelasSiswaRequest, userID string) (newKelas KelasSiswa, err error)
	ExistByIdSiswa(idSiswa string, idKelasSiswa string) (exist bool, err error)
}

type KelasSiswaServiceImpl struct {
	KelasSiswaRepository KelasSiswaRepository
	Config               *configs.Config
}

func ProvideKelasSiswaServiceImpl(repository KelasSiswaRepository, config *configs.Config) *KelasSiswaServiceImpl {
	s := new(KelasSiswaServiceImpl)
	s.KelasSiswaRepository = repository
	s.Config = config
	return s
}

func (s *KelasSiswaServiceImpl) Create(reqFormat KelasSiswaRequest, userID string) (newKelas KelasSiswa, err error) {
	newKelas, _ = newKelas.NewKelasSiswaFormat(reqFormat, userID)
	err = s.KelasSiswaRepository.Create(newKelas)
	if err != nil {
		return KelasSiswa{}, err
	}
	return newKelas, nil
}

func (s *KelasSiswaServiceImpl) ExistByIdSiswa(idSiswa string, idKelasSiswa string) (exist bool, err error) {
	exist, err = s.KelasSiswaRepository.ExistByIdSiswa(idSiswa, idKelasSiswa)
	if err != nil {
		return false, err
	}
	return
}
