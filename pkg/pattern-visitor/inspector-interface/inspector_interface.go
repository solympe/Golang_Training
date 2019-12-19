package inspector_interface

import cR "github.com/solympe/Golang_Training/pkg/pattern-visitor/company-regulator"

type Inspector interface {
	VisitCompany(company cR.CompanyRegulator) string
}
