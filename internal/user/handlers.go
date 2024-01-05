package user

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	GetAllEmployees() fiber.Handler
	CreateEmployee() fiber.Handler
	UpdateEmployee() fiber.Handler
	DeleteEmployee() fiber.Handler
}
