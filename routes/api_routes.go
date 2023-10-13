package routes

import (
	"github.com/gofiber/fiber/v2"
	"go_starter/controllers"
	"go_starter/controllers/api"
)

type apiRoutes struct {
	controllerApi     api.ControllerApi
	partnerController controllers.PartnerController
}

func (a apiRoutes) Install(app *fiber.App) {
	route := app.Group("api/", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	route.Post("hello", a.controllerApi.StartController)

	route.Post("login", a.partnerController.LoginController)
}

func NewApiRoutes(
	controllerApi api.ControllerApi,
	partnerController controllers.PartnerController,
// controller
) Routes {
	return &apiRoutes{
		controllerApi:     controllerApi,
		partnerController: partnerController,
		//controller
	}
}
