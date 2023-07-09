package verse

import (
	"quranapi/pkg/validator"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) GetVerseBySurah(c *fiber.Ctx) error {
	req := new(VerseRequest)

	if err := c.ParamsParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	errors := validator.ValidateRequest(*req)
	if !reflect.ValueOf(errors).IsNil() {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	verse, err := GetVerseBySurah(*req)
	if err != nil {
		return c.Status(verse.StatusCode).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(verse)
}
