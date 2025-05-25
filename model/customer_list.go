package model

import "encoding/json"

type CustomerList struct {
	HasMore    bool       `json:"hasMore"`
	TotalCount int        `json:"totalCount"`
	Limit      int        `json:"limit"`
	Offset     int        `json:"offset"`
	Data       []Customer `json:"data"`
}

func NewCustomerList() *CustomerList {
	return &CustomerList{
		HasMore:    false,
		TotalCount: 0,
		Limit:      10,
		Offset:     0,
		Data:       []Customer{},
	}
}

func (cl *CustomerList) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, cl); err != nil {
		return err
	}
	return nil
}
