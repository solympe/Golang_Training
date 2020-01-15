package main

import (
	"fmt"

	cp "github.com/solympe/Golang_Training/pkg/pattern-visitor/company"
	cm "github.com/solympe/Golang_Training/pkg/pattern-visitor/company_owner"
	ci "github.com/solympe/Golang_Training/pkg/pattern-visitor/inspector"
)

func main() {

	companyA := cp.NewCompanyRegulator()
	companyB := cp.NewCompanyRegulator()
	companyC := cp.NewCompanyRegulator()
	companyManager := cm.NewCompanyOwnerManager()
	inspector := ci.NewInspectorVisitor()

	companyManager.AddCompany(companyA)
	companyManager.AddCompany(companyB)
	companyManager.AddCompany(companyC)

	companyA.MakeReport("report A")
	companyB.MakeReport("report B")
	companyC.MakeReport("report C")

	companyA.AddInspector(inspector)
	companyB.AddInspector(inspector)
	companyC.AddInspector(inspector)

	fmt.Println(companyManager.VisitCompanies())
}
