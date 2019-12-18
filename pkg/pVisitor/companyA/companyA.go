package companyA

import (
	cR "github.com/solympe/Golang_Training/pkg/pVisitor/companyRegulator"
	iI "github.com/solympe/Golang_Training/pkg/pVisitor/inspectorInterface"
)

type companyA struct {
	inspector iI.Inspector
	report string
}

func (c *companyA) MakeReport(report string) {
	c.report = report
}

func (c *companyA) SendReport() string {
	return c.report
}

func (c *companyA) AcceptInspector() string{
	return c.inspector.VisitCompany(c)
}

func (c *companyA) AddInspector(inspector iI.Inspector) {
	c.inspector = inspector
}

func NewCompanyA() cR.CompanyRegulator {
	return &companyA{}
}
