package model

type CreateCreditCardRequest struct {
	UserID uint
	Number string
}

type UpdateCreditCardRequest struct {
	ID     uint
	UserID uint
	Number string
}
