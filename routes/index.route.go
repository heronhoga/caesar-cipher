package routes

import "github.com/gofiber/fiber/v2"

func RouteInit(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Name": "Hoga Cavan Afrinata",
			"Kelas": "Kriptografi C",
			"NIM": "21120121130040",
			"ImagePath": "/static/images/image.jpg",
		})
	})

	r.Get("/encryption", func (c *fiber.Ctx) error  {
		return c.Render("encrypt", fiber.Map{
			"Type": "Caesar Cipher",
			"Char": "26",
		})
	})
}