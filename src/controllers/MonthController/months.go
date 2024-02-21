package monthcontroller

import (
	"encoding/json"
	models "fetchAPI_gofiber/src/models/MonthModel"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetAllMonthsPaginated(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodGet {
		var page, limit int

		pageStr := c.Query("page")
		limitStr := c.Query("limit")

		if pageStr != "" {
			page, _ = strconv.Atoi(pageStr)
		}

		if limitStr != "" {
			limit, _ = strconv.Atoi(limitStr)
		}

		offset := (page - 1) * limit

		sort := c.Query("sort")
		if sort == "" {
			sort = "ASC"
		}
		sortBy := c.Query("sortBy")
		if sortBy == "" {
			sortBy = "name"
		}
		sort = sortBy + " " + strings.ToLower(sort)
		response := models.FindCond(sort, limit, offset)
		totalData := models.CountData()
		totalPage := math.Ceil(float64(totalData) / float64(limit))

		result := map[string]interface{}{
			"status":      "Success",
			"data":        response,
			"currentPage": page,
			"limit":       limit,
			"totalData":   totalData,
			"totalPage":   totalPage,
		}

		return c.JSON(result)
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak Diizinkan")
	}
}

func GetAllMonths(c *fiber.Ctx) error {
	months := models.SelectAllMonth()
	res, err := json.Marshal(months)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}

func GetMonthById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectMonthById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostMonth(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var month models.Month
		if err := c.BodyParser(&month); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Month{
			Name: month.Name,
			Day:  month.Day,
		}
		models.PostMonth(&item)

		return c.JSON(fiber.Map{
			"Message": "Month Posted",
		})

	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func UpdateMonth(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam) 
		var month models.Month
		if err := c.BodyParser(&month); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		newMonth := models.Month{
			Name: month.Name,
			Day:  month.Day,
		}
		models.UpdateMonth(id , &newMonth)

		return c.JSON(fiber.Map{
			"Message": "Month Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteMonth(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam) 
	models.DeleteMonth(id)

	return c.JSON(fiber.Map{
		"Message": "Month Deleted",
	})

}
