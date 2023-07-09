package surah

import (
	"quranapi/pkg/validator"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) GetSurah(c *fiber.Ctx) error {
	req := new(SurahRequest)

	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	errors := validator.ValidateRequest(*req)
	if !reflect.ValueOf(errors).IsNil() {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	surah, _ := GetSurah(*req)
	return c.Status(fiber.StatusOK).JSON(surah)
}
