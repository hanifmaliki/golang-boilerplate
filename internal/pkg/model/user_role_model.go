package model

type CreateUserRoleRequest struct {
	UserID uint
	RoleID uint
}

type UpdateUserRoleRequest struct {
	ID     uint
	UserID uint
	RoleID uint
}
