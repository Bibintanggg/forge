package service

import "github.com/Bibintanggg/forge/internal/doctor"

type DoctorService struct{}

func NewDoctorService() *DoctorService {
	return &DoctorService{}
}

func (s *DoctorService) Run() doctor.Result {
	return doctor.Run()
}
