package owner

import "github.com/solympe/Golang_Training/pkg/pattern-visitor/company"

// CompanyOwnerManager ...
type Owner interface {
	VisitCompanies() string
}

type owner struct {
	companies [] company.Company
}

// VisitCompanies ...
func (o *owner) VisitCompanies() string {
	var result string
	for _, places := range o.companies {
		result += places.AcceptInspector() + " "
	}
	return result
}

// NewOwner ...
func NewOwner(companies ...company.Company) Owner {
	var allCompanies []company.Company
	for _, eachCompany := range companies {
		allCompanies = append(allCompanies, eachCompany)
	}
	return &owner{
		companies: allCompanies,
	}
}
