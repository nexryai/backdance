package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nexryai/backdance/internal/service/news"
	"strings"
)

func FeedProxyRouter(app *fiber.App, service news.FeedProxyService) {
	app.Get("/feed", func(c *fiber.Ctx) error {
		url := c.Query("url")
		if !strings.HasPrefix(url, "https://") {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		res, err := service.Fetch(url)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(res)
	})
}
