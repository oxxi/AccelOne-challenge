package model

import "github.com/oxxi/accel-one/pkg/dto"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

func (c Contact) ToDto() dto.ContactDto {
	return dto.ContactDto{
		ID:    c.ID,
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
	}
}
