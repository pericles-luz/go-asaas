package model

import (
	"encoding/json"
	"errors"
)

var (
	ErrNameIsRequired     = errors.New("name is required")
	ErrDocumentIsRequired = errors.New("document is required")
	ErrNoContactInfo      = errors.New("at least one contact info (mobile phone or email) is required")
)

type Customer struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	CpfCnpj           string `json:"cpfCnpj"`
	MobilePhone       string `json:"mobilePhone"`
	Email             string `json:"email"`
	PersonType        string `json:"personType"` // "FISICA" or "JURIDICA"
	ExternalReference string `json:"externalReference"`
}

func NewCustomer() *Customer {
	return &Customer{}
}

func (c *Customer) SetName(name string) *Customer {
	c.Name = name
	return c
}

func (c *Customer) SetCpfCnpj(cpfCnpj string) *Customer {
	c.CpfCnpj = cpfCnpj
	return c
}

func (c *Customer) SetMobilePhone(mobilePhone string) *Customer {
	c.MobilePhone = mobilePhone
	return c
}

func (c *Customer) SetEmail(email string) *Customer {
	c.Email = email
	return c
}

func (c *Customer) SetPersonType(personType string) *Customer {
	c.PersonType = personType
	return c
}

func (c *Customer) SetExternalReference(externalReference string) *Customer {
	c.ExternalReference = externalReference
	return c
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return ErrNameIsRequired
	}
	if c.CpfCnpj == "" {
		return ErrDocumentIsRequired
	}
	if c.MobilePhone == "" && c.Email == "" {
		return ErrNoContactInfo
	}
	return nil
}

func (c *Customer) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"name":    c.Name,
		"cpfCnpj": c.CpfCnpj,
	}
	if c.MobilePhone != "" {
		result["mobilePhone"] = c.MobilePhone
	}
	if c.Email != "" {
		result["email"] = c.Email
	}
	if c.PersonType != "" {
		result["personType"] = c.PersonType
	}
	if c.ExternalReference != "" {
		result["externalReference"] = c.ExternalReference
	}
	return result
}

func (c *Customer) Unmarshal(raw []byte) error {
	if err := json.Unmarshal(raw, c); err != nil {
		return err
	}
	return c.Validate()
}
