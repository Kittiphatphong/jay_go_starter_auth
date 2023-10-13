package requests

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type NameRequest struct {
	Name string `json:"name" validate:"required"`
}

type UserAddRoleRequest struct {
	RoleId uint `json:"role_id" validate:"required"`
	UserId uint `json:"user_id" validate:"required"`
}

type RoleAddPermissionRequest struct {
	RoleId        int   `json:"role_id" validate:"required"`
	PermissionsId []int `json:"permissions_id" validate:"required,min=1"`
}
