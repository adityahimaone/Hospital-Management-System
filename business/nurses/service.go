package nurses

import (
	"Hospital-Management-System/app/middlewares"
	"Hospital-Management-System/business"
	"Hospital-Management-System/helpers/encrypt"
	"time"
)

type serviceNurse struct {
	nurseRepository Repository
	contextTimeout   time.Duration
	jwtAuth          *middlewares.ConfigJWT
}

func NewServiceNurse(repoNurse Repository, timeout time.Duration, jwtauth *middlewares.ConfigJWT) Service {
	return &serviceNurse{
		nurseRepository: repoNurse,
		contextTimeout:   timeout,
		jwtAuth:          jwtauth,
	}
}

func (serv *serviceNurse) AllNurse() ([]Domain, error) {

	result, err := serv.nurseRepository.AllNurse()

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (serv *serviceNurse) Register(domain *Domain) (Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	domain.Password = hashedPassword

	result, err := serv.nurseRepository.Register(domain)

	if result == (Domain{}) {
		return Domain{}, business.ErrDuplicateData
	}

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return result, nil
}

func (serv *serviceNurse) Login(username, password string) (Domain, error) {

	result, err := serv.nurseRepository.Login(username, password)

	if err != nil {
		return Domain{}, business.ErrEmailorPass
	}

	checkPass := encrypt.CheckPasswordHash(password, result.Password)

	if !checkPass {
		return Domain{}, business.ErrEmailorPass
	}

	result.Token = serv.jwtAuth.GenerateToken(result.ID, "nurse")

	return result, nil
}

func (serv *serviceNurse) Update(nurseID int, domain *Domain) (Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)
	domain.Password = hashedPassword
	result, err := serv.nurseRepository.Update(nurseID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceNurse) NurseByID(id int) (Domain, error) {

	result, err := serv.nurseRepository.NurseByID(id)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceNurse) Delete(id int) (string, error) {

	result, err := serv.nurseRepository.Delete(id)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}