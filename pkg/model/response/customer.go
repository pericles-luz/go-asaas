package response

import (
	"encoding/json"

	"github.com/pericles-luz/go-asaas/pkg/model"
)

type Customer struct {
	Object model.Customer `json:"object"`
}

func (c *Customer) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}

func NewCustomer(data []byte) (*Customer, error) {
	customer := &Customer{}
	if err := customer.Unmarshal(data); err != nil {
		return nil, err
	}
	return customer, nil
}
