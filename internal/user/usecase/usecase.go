package usecase

import (
	"Human_Resources_Managament_System/config"
	"Human_Resources_Managament_System/internal/user"
	"Human_Resources_Managament_System/models/employee"
	"Human_Resources_Managament_System/pkg/utils"
	"context"
)

type UserUC struct {
	cfg       *config.Config
	mongoRepo user.Repository
}

func NewUserUC(cfg *config.Config, mongoDB user.Repository) *UserUC {
	return &UserUC{
		cfg:       cfg,
		mongoRepo: mongoDB,
	}
}

func (u *UserUC) GetAllEmployees(ctx context.Context) ([]employee.Employee, error) {
	employeeData, err := u.mongoRepo.GetAllEmployeesData(ctx)
	if err != nil {
		return []employee.Employee{}, err
	}

	employees := utils.Map(
		employeeData,
		func(item employee.EmployeeData) employee.Employee {
			return item.ToServer()
		},
	)

	return employees, nil
}

func (u *UserUC) CreateEmployee(ctx context.Context, info employee.Employee) (employee.Employee, error) {
	employeeData, err := u.mongoRepo.CreateEmployee(ctx, info)
	if err != nil {
		return employee.Employee{}, err
	}

	return employeeData.ToServer(), nil
}

func (u *UserUC) UpdateEmployee(ctx context.Context, id string, info employee.Employee) error {
	err := u.mongoRepo.UpdateEmployeeByID(ctx, id, info)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUC) DeleteEmployee(ctx context.Context, id string) error {
	err := u.mongoRepo.DeleteEmployeeByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
