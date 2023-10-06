package web

type CompanyCreateRequest struct {
	CompanyName     string  `validate:"required,min=3,max=50" json:"company_name"`
	TelephoneNumber *string `validate:"omitempty,min=8,max=16" json:"telephone_number"`
	Address         *string `validate:"omitempty,min=10,max=50" json:"address"`
}

type CompanyUpdateRequest struct {
	ID              string  `validate:"required" json:"id"`
	CompanyName     *string `validate:"omitempty,min=3,max=50" json:"company_name"`
	TelephoneNumber *string `validate:"omitempty,min=8,max=16" json:"telephone_number"`
	IsActive        *bool   `validate:"omitempty" json:"is_active"`
	Address         *string `validate:"omitempty,min=10,max=50" json:"address"`
}
