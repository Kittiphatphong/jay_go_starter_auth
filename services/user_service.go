package services

import (
	"go_starter/models"
	"go_starter/repositories"
	"go_starter/requests"
)

type UserService interface {
	//Insert your function interface
	CreateUserService(model models.User) error
	CreateRoleService(model models.Role) error
	UserAddRoleService(model models.User) error
	RoleAddPermission(roleId int, permissionId []int) error

	FindByCredentialsService(request requests.LoginRequest) (*models.User, error)
	UserInfoService(id uint) (*models.User, error)
	GetPermissionService() ([]models.Permission, error)
	GetRoleService() ([]models.Role, error)
	GetUserService() ([]models.User, error)
}

type userService struct {
	repositoryUser repositories.UserRepository
}

func (u userService) RoleAddPermission(roleId int, permissionId []int) error {
	err := u.repositoryUser.RoleAddPermissions(roleId, permissionId)
	if err != nil {
		return err
	}
	return nil
}

func (u userService) GetUserService() ([]models.User, error) {
	responses, err := u.repositoryUser.GetUser()
	if err != nil {
		return nil, err
	}
	return responses, nil
}

func (u userService) GetRoleService() ([]models.Role, error) {
	responses, err := u.repositoryUser.GetRole()
	if err != nil {
		return nil, err
	}
	return responses, nil
}

func (u userService) UserAddRoleService(model models.User) error {
	err := u.repositoryUser.UserAddRole(model)
	if err != nil {
		return err
	}
	return nil
}

func (u userService) GetPermissionService() ([]models.Permission, error) {
	responses, err := u.repositoryUser.GetPermission()
	if err != nil {
		return nil, err
	}
	return responses, nil
}

func (u userService) CreateRoleService(model models.Role) error {
	err := u.repositoryUser.CreateRole(model)
	if err != nil {
		return err
	}

	return nil
}

func (u userService) UserInfoService(id uint) (*models.User, error) {

	response, err := u.repositoryUser.GetUserInfo(id)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u userService) FindByCredentialsService(request requests.LoginRequest) (*models.User, error) {
	response, err := u.repositoryUser.FindByCredentials(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u userService) CreateUserService(model models.User) error {
	err := u.repositoryUser.CreateUser(model)
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(
	repositoryUser repositories.UserRepository,
	// repo
) UserService {
	return &userService{
		repositoryUser: repositoryUser,
		//repo
	}
}
