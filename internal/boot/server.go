package boot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nexryai/backdance/internal/router"
	"github.com/nexryai/backdance/internal/service/news"
)

func StartServer() {
	app := fiber.New()

	feedProxyService := news.NewCommonFeedProxyService()
	router.FeedProxyRouter(app, feedProxyService)

	app.Listen(":3000")
}
