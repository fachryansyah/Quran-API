package main

import (
	"quranapi/internal/surah"
	"quranapi/internal/verse"
	_ "quranapi/pkg/daemon"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Quran",
		AppName:       "Quran API v1.0",
	})

	apiEndpoint := app.Group("/api")

	surah := surah.NewRoute()
	surah.Routes(apiEndpoint)

	verse := verse.NewRoute()
	verse.Routes(apiEndpoint)

	app.Listen(":3000")
}
