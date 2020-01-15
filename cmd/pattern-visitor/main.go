package main

import (
	"fmt"

	"github.com/solympe/Golang_Training/pkg/pattern-visitor/company"
	"github.com/solympe/Golang_Training/pkg/pattern-visitor/inspector"
	"github.com/solympe/Golang_Training/pkg/pattern-visitor/owner"
)

func main() {
	companyInspector := inspector.NewInspector()
	companyA := company.NewCompany(companyInspector)
	companyB := company.NewCompany(companyInspector)
	companyC := company.NewCompany(companyInspector)

	companyOwner := owner.NewOwner(companyA, companyB, companyC)

	//companyA.MakeReport("report A")
	companyB.MakeReport("report B")
	companyC.MakeReport("report C")

	fmt.Println(companyOwner.VisitCompanies())
}
