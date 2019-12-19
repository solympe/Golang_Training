package company_regulator

import (
	iI "github.com/solympe/Golang_Training/pkg/pattern-visitor/inspector-interface"
)

// CompanyRegulator represents companies interface
type CompanyRegulator interface {
	MakeReport(report string)
	SendReport() string
	AddInspector(inspector iI.Inspector)
	AcceptInspector() string
}
