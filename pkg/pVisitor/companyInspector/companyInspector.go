package companyInspector

import (
	cR "github.com/solympe/Golang_Training/pkg/pVisitor/companyRegulator"
	iI "github.com/solympe/Golang_Training/pkg/pVisitor/inspectorInterface"
)

type inspector struct {
}

func (i *inspector) VisitCompany(company cR.CompanyRegulator) string{
	return company.SendReport()
}

func NewInspector() iI.Inspector {
	return &inspector{}
}