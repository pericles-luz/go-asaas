package model

import (
	"encoding/json"
	"errors"
)

var (
	ErrAccessTokenIsRequired = errors.New("access token is required")
	ErrLinkIsRequired        = errors.New("link is required")
)

type Credential struct {
	AccessToken string `json:"access_token"`
	Link        string `json:"link"`
}

func NewCredential() *Credential {
	return &Credential{
		AccessToken: "",
		Link:        "",
	}
}

func (a *Credential) SetAccessToken(token string) *Credential {
	a.AccessToken = token
	return a
}

func (a *Credential) SetLink(link string) *Credential {
	a.Link = link
	return a
}

func (a *Credential) Validate() error {
	if a.AccessToken == "" {
		return ErrAccessTokenIsRequired
	}
	if a.Link == "" {
		return ErrLinkIsRequired
	}
	return nil
}

func (a *Credential) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, a); err != nil {
		return err
	}
	return a.Validate()
}
