package employee

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Salary float64 `json:"salary,omitempty"`
	Age    int64   `json:"age,omitempty"`
}

type EmployeeData struct {
	ID     string  `db:"id" bson:"_id,omitempty"`
	Name   string  `db:"name"`
	Salary float64 `db:"salary"`
	Age    int64   `db:"age"`
}

func (e *EmployeeData) ToServer() Employee {
	return Employee{
		ID:     e.ID,
		Name:   e.Name,
		Salary: e.Salary,
		Age:    e.Age,
	}
}

func (e *Employee) ToStorage() EmployeeData {
	return EmployeeData{
		ID:     e.ID,
		Name:   e.Name,
		Salary: e.Salary,
		Age:    e.Age,
	}
}
