package model

type CreateUserRoleRequest struct {
	UserID uint
	RoleID uint

	Role *CreateRoleRequest
}

type UpdateUserRoleRequest struct {
	ID     uint
	UserID uint
	RoleID uint

	Role *UpdateRoleRequest
}
