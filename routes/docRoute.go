package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func DocRoute(app *fiber.App) {
	// app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("/docs/*", swagger.New(swagger.Config{ // custom
		DeepLinking:  false,
		DocExpansion: "none",
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		Title:             "Documentación",
		ValidatorUrl:      "nil",
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))
}
