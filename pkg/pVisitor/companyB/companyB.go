package companyB

import (
	cR "github.com/solympe/Golang_Training/pkg/pVisitor/companyRegulator"
	iI "github.com/solympe/Golang_Training/pkg/pVisitor/inspectorInterface"
)

type companyB struct {
	inspector iI.Inspector
	report string
}

func (c *companyB) MakeReport(report string) {
	c.report = report
}

func (c *companyB) SendReport() string {
	return c.report
}

func (c *companyB) AcceptInspector() string{
	return c.inspector.VisitCompany(c)
}

func (c *companyB) AddInspector(inspector iI.Inspector) {
	c.inspector = inspector
}

func NewCompanyB() cR.CompanyRegulator {
	return &companyB{}
}
