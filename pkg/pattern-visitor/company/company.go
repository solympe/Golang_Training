package company

type inspector interface {
	VisitCompany(company CompanyRegulator) string
}

// CompanyRegulator represents companies interface
type CompanyRegulator interface {
	MakeReport(report string)
	SendReport() string
	AddInspector(inspector inspector)
	AcceptInspector() string
}

type company struct {
	inspector inspector
	report    string
}

// MakeReport ...
func (c *company) MakeReport(report string) {
	c.report = report
}

// SendReport ...
func (c *company) SendReport() string {
	return c.report
}

// AcceptInspector ...
func (c *company) AcceptInspector() string {
	return c.inspector.VisitCompany(c)
}

// AddInspector ...
func (c *company) AddInspector(inspector inspector) {
	c.inspector = inspector
}

// NewCompanyRegulator ...
func NewCompanyRegulator() CompanyRegulator {
	return &company{}
}
