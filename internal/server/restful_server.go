package server

import (
	"example/microabstraction/internal/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RestfulServer struct {
	sumService *service.SumService
}

func (s *RestfulServer) Run() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		sum := s.sumService.GetSum(1, 2)
		return c.SendString(fmt.Sprintf("Hello, World 👋! %f", sum))
	})

	app.Listen(":3000")
}

func NewRestfulServer(service *service.SumService) *RestfulServer {
	return &RestfulServer{sumService: service}
}
