package request

import (
	"Hospital-Management-System/business/patients"
)

type Patients struct {
	Fullname string `json:"fullname"`
	Gender   string `json:"gender"`
	NIK      int    `json:"nik"`
	Address  string `json:"address"`
	DOB      string `json:"dob"`
	No_Rm    string `json:"no_rm"`
}

func (req *Patients) ToDomain() *patients.Domain {
	return &patients.Domain{

		Fullname: req.Fullname,
		Gender:   req.Gender,
		NIK:      req.NIK,
		Address:  req.Address,
		DOB:      req.DOB,
		No_Rm:    req.No_Rm,
	}
}
