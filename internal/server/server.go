package server

import (
	"Human_Resources_Managament_System/config"
	"Human_Resources_Managament_System/pkg/mongodb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	cfg     *config.Config
	fiber   *fiber.App
	mongoDB mongodb.MongoDBManager
}

func NewServer(cfg *config.Config, mongoDB mongodb.MongoDBManager) *Server {
	return &Server{
		cfg:     cfg,
		fiber:   fiber.New(),
		mongoDB: mongoDB,
	}
}

func (s *Server) Run() error {
	err := s.MapHandlers()
	if err != nil {
		return err
	}

	go func() {
		s.fiber.Get("/health_check", func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})
		log.Info("Server is started ", s.cfg.Server.Host)

		err := s.fiber.Listen(s.cfg.Server.Host)
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit

	err = s.fiber.Shutdown()
	if err != nil {
		log.Error(err)
	} else {
		log.Infof("Fiber closed properly")
	}

	return nil
}
