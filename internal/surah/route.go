package surah

import (
	"github.com/gofiber/fiber/v2"
)

type Rest struct{}

func NewRoute() *Rest {
	return &Rest{}
}

func (c *Rest) Routes(r fiber.Router) {
	routes := r.Group("/surah")
	routes.Get("/", c.GetSurah) // => BASE_URL/api/surah/
	routes.Get("", c.GetSurah)  // => BASE_URL/api/surah
}
