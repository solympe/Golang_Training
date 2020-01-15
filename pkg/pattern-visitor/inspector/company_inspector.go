package inspector

import "github.com/solympe/Golang_Training/pkg/pattern-visitor/company"

type inspector struct{}

// VisitCompany ...
func (i *inspector) VisitCompany(company company.CompanyRegulator) string {
	return company.SendReport()
}

// NewInspectorVisitor ...
func NewInspectorVisitor() InspectorVisitor {
	return &inspector{}
}
