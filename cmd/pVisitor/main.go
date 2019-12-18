package main

import (
	"fmt"

	cA "github.com/solympe/Golang_Training/pkg/pVisitor/companyA"
	cB "github.com/solympe/Golang_Training/pkg/pVisitor/companyB"
	cC "github.com/solympe/Golang_Training/pkg/pVisitor/companyC"
	corp "github.com/solympe/Golang_Training/pkg/pVisitor/companyOwner"
	ci "github.com/solympe/Golang_Training/pkg/pVisitor/companyInspector"
)

func main() {

	companyA := cA.NewCompanyA()
	companyB := cB.NewCompanyB()
	companyC := cC.NewCompanyC()
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
