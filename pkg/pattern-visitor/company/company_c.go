package company

import (
	cR "github.com/solympe/Golang_Training/pkg/pattern-visitor/company-regulator"
	iI "github.com/solympe/Golang_Training/pkg/pattern-visitor/inspector-interface"
)

type companyC struct {
	inspector iI.Inspector
	report string
}

func (c *companyC) MakeReport(report string) {
	c.report = report
}

func (c *companyC) SendReport() string {
	return c.report
}

func (c *companyC) AcceptInspector() string{
	return c.inspector.VisitCompany(c)
}

func (c *companyC) AddInspector(inspector iI.Inspector) {
	c.inspector = inspector
}

func NewCompanyC() cR.CompanyRegulator {
	return &companyC{}
}
