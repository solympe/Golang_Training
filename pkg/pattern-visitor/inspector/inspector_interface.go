package inspector

import "github.com/solympe/Golang_Training/pkg/pattern-visitor/company"

// InspectorVisitor ...
type InspectorVisitor interface {
	VisitCompany(company company.CompanyRegulator) string
}
