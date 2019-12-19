package cooperative

import (
	cC "github.com/solympe/Golang_Training/pkg/pattern-visitor/company-regulator"
	iI "github.com/solympe/Golang_Training/pkg/pattern-visitor/inspector-interface"
)

type Corporation interface {
	AddCompany(company cC.CompanyRegulator)
	VisitCompanies(i iI.Inspector) string
}

type corporationOwner struct {
	companies [] cC.CompanyRegulator
}

func (c *corporationOwner) AddCompany(company cC.CompanyRegulator) {
	c.companies = append(c.companies, company)
}

func (c *corporationOwner) VisitCompanies(i iI.Inspector) string{
	var result string
	for _, places := range c.companies {
		result += places.AcceptInspector(i) + " "
	}
	return result
}

func NewCorporation() Corporation {
	return &corporationOwner{}
}
