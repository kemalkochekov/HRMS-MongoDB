package user

import (
	"Human_Resources_Managament_System/models/employee"
	"context"
)

type UseCase interface {
	GetAllEmployees(ctx context.Context) ([]employee.Employee, error)
	CreateEmployee(ctx context.Context, info employee.Employee) (employee.Employee, error)
	UpdateEmployee(ctx context.Context, id string, info employee.Employee) error
	DeleteEmployee(ctx context.Context, id string) error
}
