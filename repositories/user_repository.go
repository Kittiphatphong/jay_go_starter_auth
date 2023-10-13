package repositories

import (
	"errors"
	"go_starter/models"
	"go_starter/requests"
	"go_starter/trails"
	"gorm.io/gorm"
)

type UserRepository interface {
	//Insert your function interface
	CreateUser(model models.User) error
	CreateRole(model models.Role) error
	UserAddRole(user models.User) error
	RoleAddPermissions(roleId int, permissionsId []int) error

	FindByCredentials(request requests.LoginRequest) (*models.User, error)
	GetUserInfo(id uint) (*models.User, error)
	GetPermission() ([]models.Permission, error)
	GetRole() ([]models.Role, error)
	GetUser() ([]models.User, error)
}

type userRepository struct{ db *gorm.DB }

func (u userRepository) RoleAddPermissions(roleId int, permissionsId []int) error {
	begin := u.db.Begin()
	deleteError := begin.Unscoped().Delete(&models.RolePermission{}, "role_id = ?", roleId).Error
	if deleteError != nil {
		begin.Rollback()
		return deleteError
	}
	for i := 0; i < len(permissionsId); i++ {
		rolePermission := models.RolePermission{
			RoleID:       roleId,
			PermissionID: permissionsId[i],
		}
		createError := begin.Create(&rolePermission).Error
		if createError != nil {
			begin.Rollback()
			return createError
		}
	}
	begin.Commit()
	return nil
}

func (u userRepository) CreateUser(model models.User) error {
	hashPassword, err := trails.HashPassword(model.Password)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     model.Name,
		Username: model.Username,
		Password: hashPassword,
	}

	result := u.db.Create(&user)

	if result.Error != nil {
		if result.Error.Error() == `ERROR: duplicate key value violates unique constraint "users_username_key" (SQLSTATE 23505)` {
			return errors.New("Username is exist")
		}

		return result.Error
	}

	return nil
}

func (u userRepository) CreateRole(model models.Role) error {
	err := u.db.Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) UserAddRole(user models.User) error {
	err := u.db.Model(&user).Where("id = ?", user.ID).Updates(models.User{RoleID: user.RoleID}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) FindByCredentials(request requests.LoginRequest) (*models.User, error) {
	var model models.User

	result := u.db.Where("username", request.Username).First(&model)
	if result.RowsAffected <= 0 {
		return nil, errors.New("User not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	if !trails.CheckPasswordHash(request.Password, model.Password) {
		return nil, errors.New("Wrong password")
	}

	if request.Username == model.Username && trails.CheckPasswordHash(request.Password, model.Password) {
		return &models.User{
			ID:       model.ID,
			Name:     model.Name,
			Username: model.Username,
		}, nil
	}
	return nil, errors.New("user not found")
}

func (u userRepository) GetUserInfo(id uint) (*models.User, error) {
	var model models.User
	err := u.db.Preload("Role").Preload("Role.Permissions").First(&model, id).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (u userRepository) GetPermission() ([]models.Permission, error) {
	var model []models.Permission
	err := u.db.Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (u userRepository) GetRole() ([]models.Role, error) {
	var model []models.Role
	err := u.db.Preload("Permissions").Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (u userRepository) GetUser() ([]models.User, error) {
	var model []models.User
	err := u.db.Preload("Role").Preload("Role.Permissions").Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	// db.Migrator().DropTable(models.User{})
	//db.AutoMigrate(models.User{})
	//db.AutoMigrate(models.Role{})
	//db.AutoMigrate(models.Permission{})
	return &userRepository{db: db}
}
