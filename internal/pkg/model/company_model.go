package model

type CreateCompanyRequest struct {
	Name string
}

type UpdateCompanyRequest struct {
	ID   uint
	Name string
}
