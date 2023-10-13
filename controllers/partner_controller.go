package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go_starter/models"
	"go_starter/requests"
	"go_starter/responses"
	"go_starter/services"
	"go_starter/validation"
	"net/http"
)

type PartnerController interface {
	//Insert your function interface
	CreatePartnerController(ctx *fiber.Ctx) error
	ListPartnerController(ctx *fiber.Ctx) error

	//api
	LoginController(ctx *fiber.Ctx) error
}

type partnerController struct {
	servicePartner services.PartnerService
}

func (p partnerController) ListPartnerController(ctx *fiber.Ctx) error {
	response, err := p.servicePartner.ListPartnerService()
	if err != nil {
		return NewErrorValidate(ctx, err)
	}
	return NewSuccessResponse(ctx, response)

}

func (p partnerController) CreatePartnerController(ctx *fiber.Ctx) error {
	request := requests.RegisterRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}

	err = p.servicePartner.CreateSellerService(models.Partner{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}

	return NewSuccessMsg(ctx, "success")
}

func (p partnerController) LoginController(ctx *fiber.Ctx) error {
	loginRequest := new(requests.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return NewErrorValidate(ctx, err.Error())
	}
	// Find the user by credentials
	partner, err := p.servicePartner.FindByCredentialsService(*loginRequest)
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}

	token, err := GenerateTokenApi(*partner)
	if err != nil {
		return NewErrorValidate(ctx, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": true,
		"data": responses.UserResponse{
			Name:     partner.Name,
			Username: partner.Username,
		},
		"token": token,
	})
}

func NewPartnerController(
	servicePartner services.PartnerService,
	// services
) PartnerController {
	return &partnerController{
		servicePartner: servicePartner,
		//services
	}
}
