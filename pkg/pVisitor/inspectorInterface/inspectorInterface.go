package inspectorInterface

import cR "github.com/solympe/Golang_Training/pkg/pVisitor/companyRegulator"

type Inspector interface {
	VisitCompany(company cR.CompanyRegulator) string
}
