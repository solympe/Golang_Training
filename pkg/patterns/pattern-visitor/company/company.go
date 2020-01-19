package company

type inspector interface {
	VisitCompany(company Company) string
}

// Company ...
type Company interface {
	MakeReport(report string)
	SendReport() string
	AcceptInspector() string
}

type company struct {
	inspector inspector
	report    string
	reported  bool
}

// MakeReport ...
func (c *company) MakeReport(report string) {
	c.report = report
	c.reported = true
}

// SendReport ...
func (c *company) SendReport() string {
	if c.reported == false {
		return "EMPTY"
	}
	return c.report
}

// AcceptInspector ...
func (c *company) AcceptInspector() string {
	return c.inspector.VisitCompany(c)
}

// NewCompany ...
func NewCompany(ins inspector) Company {
	return &company{
		inspector: ins,
	}
}
