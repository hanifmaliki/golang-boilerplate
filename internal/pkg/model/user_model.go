package model

type CreateUserRequest struct {
	Name      string
	Email     string
	CompanyID uint

	Address     *CreateAddressRequest
	CreditCards []*CreateCreditCardRequest
	UserRoles   []*CreateUserRoleRequest
}

type UpdateUserRequest struct {
	ID        uint
	Name      string
	Email     string
	CompanyID uint

	Address     *UpdateAddressRequest
	CreditCards []*UpdateCreditCardRequest
	UserRoles   []*UpdateUserRoleRequest
}
