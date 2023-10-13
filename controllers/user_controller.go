package controllers

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"go_starter/models"
	"go_starter/requests"
	"go_starter/responses"
	"go_starter/services"
	"go_starter/validation"
	"net/http"
	"strconv"
)

type UserController interface {
	//Insert your function interface
	RegisterUserController(ctx *fiber.Ctx) error
	CreateRoleController(ctx *fiber.Ctx) error
	LoginController(ctx *fiber.Ctx) error
	UserInfoController(ctx *fiber.Ctx) error
	GetPermissionController(ctx *fiber.Ctx) error
	UserAddRoleController(ctx *fiber.Ctx) error
	GetRoleController(ctx *fiber.Ctx) error
	GetUserController(ctx *fiber.Ctx) error
	RoleAddPermissionController(ctx *fiber.Ctx) error
}

type userController struct {
	serviceUser services.UserService
}

func (u userController) RoleAddPermissionController(ctx *fiber.Ctx) error {
	request := requests.RoleAddPermissionRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}

	errValidate := validation.Validate(request)

	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}

	err = u.serviceUser.RoleAddPermission(request.RoleId, request.PermissionsId)
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}

	return NewSuccessMsg(ctx, "success")
}

func (u userController) GetUserController(ctx *fiber.Ctx) error {
	responsesData, err := u.serviceUser.GetUserService()
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}

	return NewSuccessResponse(ctx, responsesData)
}

func (u userController) GetRoleController(ctx *fiber.Ctx) error {
	responsesData, err := u.serviceUser.GetRoleService()
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}

	return NewSuccessResponse(ctx, responsesData)
}

func (u userController) UserAddRoleController(ctx *fiber.Ctx) error {
	request := requests.UserAddRoleRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}

	errValidate := validation.Validate(request)

	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}

	err = u.serviceUser.UserAddRoleService(models.User{
		ID:     request.UserId,
		RoleID: &request.RoleId,
	})
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}

	return NewSuccessMsg(ctx, "success")
}

func (u userController) GetPermissionController(ctx *fiber.Ctx) error {
	responsesData, err := u.serviceUser.GetPermissionService()
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}

	return NewSuccessResponse(ctx, responsesData)
}

func (u userController) UserInfoController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	id := claims["id"].(string)
	i, err := strconv.Atoi(id)
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}
	response, err := u.serviceUser.UserInfoService(uint(i))
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}
	return NewSuccessResponse(ctx, response)
}

func (u userController) LoginController(ctx *fiber.Ctx) error {

	loginRequest := new(requests.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return NewErrorValidate(ctx, err.Error())
	}
	// Find the user by credentials
	user, err := u.serviceUser.FindByCredentialsService(*loginRequest)
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}

	token, err := GenerateTokenWeb(*user)
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": true,
		"user": responses.UserResponse{
			Name:     user.Name,
			Username: user.Username,
		},
		"token": token,
	})
}

func (u userController) RegisterUserController(ctx *fiber.Ctx) error {
	request := requests.RegisterRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}

	err = u.serviceUser.CreateUserService(models.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}

	return NewSuccessMsg(ctx, "success")

}

func (u userController) CreateRoleController(ctx *fiber.Ctx) error {
	request := requests.NameRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}

	err = u.serviceUser.CreateRoleService(models.Role{
		Name: request.Name,
	})

	return NewSuccessMsg(ctx, "success")
}

func NewUserController(
	serviceUser services.UserService,
	// services
) UserController {
	return &userController{
		serviceUser: serviceUser,
		//services
	}
}
