package company

import (
	cR "github.com/solympe/Golang_Training/pkg/pattern-visitor/company-regulator"
	iI "github.com/solympe/Golang_Training/pkg/pattern-visitor/inspector-interface"
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
