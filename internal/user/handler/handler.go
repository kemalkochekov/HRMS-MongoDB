package handler

import (
	"Human_Resources_Managament_System/config"
	"Human_Resources_Managament_System/internal/user"
	"Human_Resources_Managament_System/models/employee"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	cfg    *config.Config
	userUC user.UseCase
}

func NewUserHandler(cfg *config.Config, clientUC user.UseCase) *UserHandlers {
	return &UserHandlers{userUC: clientUC, cfg: cfg}
}

func (u *UserHandlers) GetAllEmployees() fiber.Handler {
	return func(c *fiber.Ctx) error {
		employees, err := u.userUC.GetAllEmployees(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(employees)
	}
}

func (u *UserHandlers) CreateEmployee() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var info employee.Employee

		if err := c.BodyParser(&info); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		employeeInfo, err := u.userUC.CreateEmployee(c.Context(), info)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(employeeInfo)
	}
}
func (u *UserHandlers) UpdateEmployee() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var info employee.Employee

		if err := c.BodyParser(&info); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		err := u.userUC.UpdateEmployee(c.Context(), id, info)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func (u *UserHandlers) DeleteEmployee() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		err := u.userUC.DeleteEmployee(c.Context(), id)
		if err != nil {
			if errors.Is(err, errors.New("Not Found")) {
				return c.Status(fiber.StatusNotFound).SendString("Employee ID Not Found")
			}
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
