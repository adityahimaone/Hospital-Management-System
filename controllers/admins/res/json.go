package res

import (
	"Hospital-Management-System/business/admins"
)

type AdminLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func FromDomainLogin(domain admins.Domain) AdminLoginResponse {
	return AdminLoginResponse{
		Message: "Admin Login Success",
		Token:   domain.Token,
	}
}
