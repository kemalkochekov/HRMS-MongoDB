package user

import (
	"Human_Resources_Managament_System/models/employee"
	"context"
)

type Repository interface {
	GetAllEmployeesData(ctx context.Context) ([]employee.EmployeeData, error)
	CreateEmployee(ctx context.Context, info employee.Employee) (employee.EmployeeData, error)
	UpdateEmployeeByID(ctx context.Context, id string, info employee.Employee) error
	DeleteEmployeeByID(ctx context.Context, id string) error
}
