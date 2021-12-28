package doctors

import (
	"Hospital-Management-System/app/middlewares"
	"Hospital-Management-System/business"
	"Hospital-Management-System/helpers/encrypt"
	"time"
)

type serviceDoctor struct {
	doctorRepository Repository
	contextTimeout   time.Duration
	jwtAuth          *middlewares.ConfigJWT
}

func NewServiceDoctor(repoDoctor Repository, timeout time.Duration, jwtauth *middlewares.ConfigJWT) Service {
	return &serviceDoctor{
		doctorRepository: repoDoctor,
		contextTimeout:   timeout,
		jwtAuth:          jwtauth,
	}
}
func (serv *serviceDoctor) AllDoctor() ([]Domain, error) {

	result, err := serv.doctorRepository.AllDoctor()

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}
func (serv *serviceDoctor) Register(domain *Domain) (Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	domain.Password = hashedPassword

	result, err := serv.doctorRepository.Register(domain)

	if result == (Domain{}) {
		return Domain{}, business.ErrDuplicateData
	}

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return result, nil
}

func (serv *serviceDoctor) Login(username, password string) (Domain, error) {

	result, err := serv.doctorRepository.Login(username, password)

	if err != nil {
		return Domain{}, business.ErrEmailorPass
	}

	checkPass := encrypt.CheckPasswordHash(password, result.Password)

	if !checkPass {
		return Domain{}, business.ErrEmailorPass
	}

	result.Token = serv.jwtAuth.GenerateToken(result.ID, "doctor")

	return result, nil
}
func (serv *serviceDoctor) Update(docID int, domain *Domain) (Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)
	domain.Password = hashedPassword
	result, err := serv.doctorRepository.Update(docID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (serv *serviceDoctor) DoctorByID(id int) (Domain, error) {

	result, err := serv.doctorRepository.DoctorByID(id)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (serv *serviceDoctor) Delete(id int) (string, error) {

	result, err := serv.doctorRepository.Delete(id)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}
