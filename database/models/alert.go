package models

type Alert struct {
	Type        string `json:"type" validate:"required,min=1,max=40"`
	Description string `json:"description" validate:"required,min=1,max=80"`
	Created     string `json:"created" validate:"required"`
	Country     string `json:"country" validate:"required,min=1,max=80"`
}
