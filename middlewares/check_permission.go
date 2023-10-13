package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"go_starter/repositories"
	"go_starter/trails"
	"net/http"
	"strconv"
)

type CheckPermissionMiddleware interface {
	GetPermission(userId uint) ([]uint, error)
	CheckPermission(id uint) fiber.Handler
}

type checkPermissionMiddleware struct {
	userRepository repositories.UserRepository
}

func (c checkPermissionMiddleware) CheckPermission(id uint) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*jtoken.Token)
		claims := user.Claims.(jtoken.MapClaims)
		idUser := claims["id"].(string)
		i, err := strconv.Atoi(idUser)
		permissions, err := c.GetPermission(uint(i))
		if err != nil {
			return err
		}
		if trails.UintContains(permissions, id) {
			return ctx.Next()
		}
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
			"status": http.StatusForbidden,
			"error":  "NO PERMISSION",
		})
	}

}

func (c checkPermissionMiddleware) GetPermission(userId uint) ([]uint, error) {

	user, err := c.userRepository.GetUserInfo(userId)
	if err != nil {
		return nil, err
	}
	var permissions []uint
	for _, permission := range user.Role.Permissions {
		permissions = append(permissions, permission.ID)
	}

	return permissions, nil
}

func NewCheckPermissionMiddleware(userRepository repositories.UserRepository) CheckPermissionMiddleware {
	return &checkPermissionMiddleware{userRepository: userRepository}
}
