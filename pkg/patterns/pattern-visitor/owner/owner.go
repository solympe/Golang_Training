package owner

import "github.com/solympe/Golang_Training/pkg/patterns/pattern-visitor/company"

// Owner ...
type Owner interface {
	VisitCompanies() string
}

type owner struct {
	companies [] company.Company
}

// VisitCompanies ...
func (o *owner) VisitCompanies() string {
	var result string
	if len(o.companies) == 0 {
		return "empty response"
	}
	for _, places := range o.companies {
		result += places.AcceptInspector() + " "
	}
	return result
}

// NewOwner ... TODO: Rewrite constructor on simple
func NewOwner(companies ...company.Company) Owner {
	var allCompanies []company.Company
	for _, eachCompany := range companies {
		allCompanies = append(allCompanies, eachCompany)
	}
	return &owner{
		companies: allCompanies,
	}
}
