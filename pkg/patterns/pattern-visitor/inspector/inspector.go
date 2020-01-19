package inspector

import "github.com/solympe/Golang_Training/pkg/patterns/pattern-visitor/company"

// Inspector ...
type Inspector interface {
	VisitCompany(company company.Company) string
}

type inspector struct {
}

// VisitCompany ...
func (ins *inspector) VisitCompany(company company.Company) string {
	return company.SendReport()
}

// NewInspector ...
func NewInspector() Inspector {
	return &inspector{
	}
}
