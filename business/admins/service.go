package admins

import (
	middleware "Hospital-Management-System/app/middlewares"
	"Hospital-Management-System/business"

	"time"
)

type serviceAdmin struct {
	adminRepository Repository
	contextTimeout  time.Duration
	jwtAuth         *middleware.ConfigJWT
}

func NewServiceAdmin(repoAdmin Repository, timeout time.Duration, jwtauth *middleware.ConfigJWT) Service {
	return &serviceAdmin{
		adminRepository: repoAdmin,
		contextTimeout:  timeout,
		jwtAuth:         jwtauth,
	}
}

func (serv *serviceAdmin) Login(username, password string) (Domain, error) {

	result, err := serv.adminRepository.Login(username, password)

	if err != nil {
		return Domain{}, business.ErrEmailorPass
	}

	result.Token = serv.jwtAuth.GenerateToken(result.ID, "admin")

	return result, nil
}
