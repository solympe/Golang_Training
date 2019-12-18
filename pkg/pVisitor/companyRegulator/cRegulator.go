package companyRegulator

import (
	iI "github.com/solympe/Golang_Training/pkg/pVisitor/inspectorInterface"
)

// CompanyRegulator represents companies interface
type CompanyRegulator interface {
	MakeReport(report string)
	SendReport() string
	AddInspector(inspector iI.Inspector)
	AcceptInspector() string
}
