package verse

import "github.com/gofiber/fiber/v2"

type Rest struct{}

func NewRoute() *Rest {
	return &Rest{}
}

func (c *Rest) Routes(r fiber.Router) {
	routes := r.Group("/verse")
	// routes.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusOK).JSON(map[string]string{
	// 		"message": "OKE!",
	// 	})
	// })
	routes.Get(":surahId", c.GetVerseBySurah)
}
