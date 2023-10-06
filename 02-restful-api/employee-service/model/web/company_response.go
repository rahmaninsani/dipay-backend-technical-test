package web

type CompanyCreateResponse struct {
	ID string `json:"id"`
}

type CompanyResponse struct {
	ID              string `json:"id"`
	CompanyName     string `json:"company_name"`
	TelephoneNumber string `json:"telephone_number"`
	IsActive        bool   `json:"is_active"`
	Address         string `json:"address"`
}
