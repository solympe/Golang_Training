package company_inspector

import (
	cR "github.com/solympe/Golang_Training/pkg/pattern-visitor/company-regulator"
	iI "github.com/solympe/Golang_Training/pkg/pattern-visitor/inspector-interface"
)

type inspector struct {
}

func (i *inspector) VisitCompany(company cR.CompanyRegulator) string{
	return company.SendReport()
}

func NewInspector() iI.Inspector {
	return &inspector{}
}