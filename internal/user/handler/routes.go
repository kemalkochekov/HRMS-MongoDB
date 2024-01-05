package handler

import (
	"Human_Resources_Managament_System/internal/user"
	"github.com/gofiber/fiber/v2"
)

func MapUserRoutes(group fiber.Router, h user.Handlers) {
	group.Get("/employee", h.GetAllEmployees())
	group.Post("/employee", h.CreateEmployee())
	group.Put("/employee/:id", h.UpdateEmployee())
	group.Delete("/employee/:id", h.DeleteEmployee())
}
