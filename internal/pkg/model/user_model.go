package model

type GetUserRequest struct {
	Search    string
	CompanyID []uint
	RoleID    []uint
}

type CreateUserRequest struct {
	Name        string
	Email       string
	PhoneNumber string
	CompanyID   uint

	Address     *CreateAddressRequest
	CreditCards []*CreateCreditCardRequest
	UserRoles   []*CreateUserRoleRequest
}

type UpdateUserRequest struct {
	ID          uint
	Name        string
	Email       string
	PhoneNumber string
	CompanyID   uint

	Address     *UpdateAddressRequest
	CreditCards []*UpdateCreditCardRequest
	UserRoles   []*UpdateUserRoleRequest
}
