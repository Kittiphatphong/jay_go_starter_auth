package routes

import (
	"github.com/gofiber/fiber/v2"
	"go_starter/controllers"
	"go_starter/middlewares"
)

type webRoutes struct {
	controller                controllers.Controller
	userController            controllers.UserController
	checkPermissionMiddleware middlewares.CheckPermissionMiddleware
	partnerController         controllers.PartnerController
}

func (w webRoutes) Install(app *fiber.App) {
	route := app.Group("web/", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	route.Post("login", w.userController.LoginController)
	route.Post("hello", w.controller.StartController)

	//with authentication
	routeAuth := app.Group("web/", middlewares.NewAuthWebMiddleware, func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	routeAuth.Post("user-info", w.userController.UserInfoController)

	//with authentication and user permission
	routeUser := app.Group("web/user/", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	routeUser.Use(w.checkPermissionMiddleware.CheckPermission(1))
	routeUser.Post("register", w.userController.RegisterUserController)
	routeUser.Post("create-role", w.userController.CreateRoleController)
	routeUser.Post("get-permissions", w.userController.GetPermissionController)
	routeUser.Post("get-roles", w.userController.GetRoleController)
	routeUser.Post("get-users", w.userController.GetUserController)
	routeUser.Post("user-add-role", w.userController.UserAddRoleController)
	routeUser.Post("role-add-permission", w.userController.RoleAddPermissionController)

	//with authentication and partner permission
	routePartner := app.Group("web/partner/")
	routePartner.Use(w.checkPermissionMiddleware.CheckPermission(2))
	routePartner.Post("create-partner", w.partnerController.CreatePartnerController)
	routePartner.Post("list-partner", w.partnerController.ListPartnerController)
}

func NewWebRoutes(
	controller controllers.Controller,
	userController controllers.UserController,
	checkPermissionMiddleware middlewares.CheckPermissionMiddleware,
	partnerController controllers.PartnerController,
	// controller
) Routes {
	return &webRoutes{
		controller:                controller,
		userController:            userController,
		checkPermissionMiddleware: checkPermissionMiddleware,
		partnerController:         partnerController,
		//controller
	}
}
