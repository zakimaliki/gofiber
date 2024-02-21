package routes

import (
	monthcontroller "fetchAPI_gofiber/src/controllers/MonthController"

	"github.com/gofiber/fiber/v2"
)

func Router(c *fiber.App) {

	v1 := c.Group("/api/v1")

	month := v1.Group("/month")
	{
		month.Get("/data", monthcontroller.GetAllMonths)
		month.Get("/:id", monthcontroller.GetMonthById)
		month.Get("/paginated-data", monthcontroller.GetAllMonthsPaginated)
		month.Post("/create", monthcontroller.PostMonth)
		month.Put("/update/:id", monthcontroller.UpdateMonth)
		month.Delete("/delete/:id", monthcontroller.DeleteMonth)
	}

}
