package company_owner

import "github.com/solympe/Golang_Training/pkg/pattern-visitor/company"

// CompanyOwnerManager ...
type CompanyOwnerManager interface {
	AddCompany(company company.CompanyRegulator)
	VisitCompanies() string
}

type companyOwner struct {
	companies [] company.CompanyRegulator
}

// AddCompany ...
func (c *companyOwner) AddCompany(company company.CompanyRegulator) {
	c.companies = append(c.companies, company)
}

// VisitCompanies ...
func (c *companyOwner) VisitCompanies() string {
	var result string
	for _, places := range c.companies {
		result += places.AcceptInspector() + " "
	}
	return result
}

// NewCompanyOwnerManager ...
func NewCompanyOwnerManager() CompanyOwnerManager {
	return &companyOwner{}
}
