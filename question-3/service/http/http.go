package http

import (
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"seven-hunter-assignment/question-3/logic"
	config2 "seven-hunter-assignment/question-3/util/config"
)

func HandleGetBeef(c *fiber.Ctx) error {
	beefLogic := logic.GetBeefLogic()
	beefs, err := beefLogic.GetBeefSummary()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := fiber.Map{}
	for _, beef := range beefs {
		response[beef.Name] = beef.Count
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func StartHttpServer() error {
	appConfig := config2.GetAppConfig()
	server := fiber.New()
	server.Use(recover2.New())

	server.Get("/beef/summary", HandleGetBeef)

	log.Fatal(server.Listen(":" + appConfig.CommonConfig.HttpPort))
	return nil
}
