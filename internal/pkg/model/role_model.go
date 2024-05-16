package model

type CreateRoleRequest struct {
	Name string
}

type UpdateRoleRequest struct {
	ID   uint
	Name string
}
