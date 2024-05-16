package model

type CreateAddressRequest struct {
	UserID  uint
	Street  string
	City    string
	Country string
}

type UpdateAddressRequest struct {
	ID      uint
	UserID  uint
	Street  string
	City    string
	Country string
}
