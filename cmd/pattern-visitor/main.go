package main

import (
	"fmt"

	"github.com/solympe/Golang_Training/pkg/pattern-visitor/company"
	ci "github.com/solympe/Golang_Training/pkg/pattern-visitor/company-inspector"
	corp "github.com/solympe/Golang_Training/pkg/pattern-visitor/cooperative"
)

func main() {

	companyA := company.NewCompanyA()
	companyB := company.NewCompanyB()
	companyC := company.NewCompanyC()
	companyA.MakeReport("report A")
	companyB.MakeReport("report B")
	companyC.MakeReport("report C")

	corporation := corp.NewCorporation()
	corporation.AddCompany(companyA)
	corporation.AddCompany(companyB)
	corporation.AddCompany(companyC)

	inspector := ci.NewInspector()
	fmt.Println(corporation.VisitCompanies(inspector))



}
