package patients

import (
	"Hospital-Management-System/business"
)

type servicePatient struct {
	patientRepository Repository
}

func NewServicePatient(repoPatient Repository) Service {
	return &servicePatient{
		patientRepository: repoPatient,
	}
}
func (serv *servicePatient) AllPatient() ([]Domain, error) {

	result, err := serv.patientRepository.AllPatient()

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}
func (serv *servicePatient) Register(domain *Domain) (Domain, error) {

	result, err := serv.patientRepository.Register(domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *servicePatient) Update(patID int, domain *Domain) (Domain, error) {

	result, err := serv.patientRepository.Update(patID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (serv *servicePatient) PatientByID(id int) (Domain, error) {

	result, err := serv.patientRepository.PatientByID(id)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (serv *servicePatient) Delete(id int) (string, error) {

	result, err := serv.patientRepository.Delete(id)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}
