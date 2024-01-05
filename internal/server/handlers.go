package server

import (
	"Human_Resources_Managament_System/internal/user/handler"
	"Human_Resources_Managament_System/internal/user/repository/mongo"
	"Human_Resources_Managament_System/internal/user/usecase"
)

func (s *Server) MapHandlers() error {
	userMongoRepo := mongo.NewUserRepo(s.mongoDB)
	userUC := usecase.NewUserUC(s.cfg, userMongoRepo)
	userHandler := handler.NewUserHandler(s.cfg, userUC)

	userGroup := s.fiber.Group("app")
	handler.MapUserRoutes(userGroup, userHandler)

	return nil
}
