package models

type Alert struct {
	Type        string `json:"email" validate:"required,min=1,max=40"`
	Description string `json:"city" validate:"required,min=1,max=80"`
	Created     string `json:"date" validate:"required"`
	Country     string `json:"country" validate:"required,min=1,max=80"`
}
