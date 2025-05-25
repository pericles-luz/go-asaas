package model

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrCustomerIDIsRequired  = errors.New("customer ID is required")
	ErrBillingTypeIsRequired = errors.New("billing type is required")
	ErrNextDueDateIsRequired = errors.New("next due date is required")
	ErrValueMustBePositive   = errors.New("value must be positive")
	ErrDescriptionIsRequired = errors.New("description is required")
	ErrCycleIsRequired       = errors.New("cycle is required")
	ErrOnlyBoletoAllowed     = errors.New("only boleto billing type is allowed for subscriptions")
)

const (
	BILLING_TYPE_BOLETO = "BOLETO"
	CYCLE_MONTHLY       = "MONTHLY"
)

type Subscription struct {
	ID          string    `json:"id"`
	CustomerID  string    `json:"customer"`
	BillingType string    `json:"billingType"`
	NextDue     string    `json:"nextDueDate"`
	NextDueDate time.Time `json:"-"`
	Value       float64   `json:"value"`
	Cycle       string    `json:"cycle"`
	Description string    `json:"description"`
}

func NewSubscription() *Subscription {
	return &Subscription{}
}

func (s *Subscription) SetID(id string) *Subscription {
	s.ID = id
	return s
}

func (s *Subscription) SetCustomerID(customerID string) *Subscription {
	s.CustomerID = customerID
	return s
}

func (s *Subscription) SetBillingType(billingType string) *Subscription {
	s.BillingType = billingType
	return s
}

func (s *Subscription) SetNextDueDate(nextDueDate string) *Subscription {
	parsedDate, err := time.Parse("2006-01-02", nextDueDate)
	if err == nil {
		s.NextDueDate = parsedDate
	}
	return s
}

func (s *Subscription) SetValue(value float64) *Subscription {
	s.Value = value
	return s
}

func (s *Subscription) SetCycle(cycle string) *Subscription {
	s.Cycle = cycle
	return s
}

func (s *Subscription) SetDescription(description string) *Subscription {
	s.Description = description
	return s
}

func (s *Subscription) Validate() error {
	if s.CustomerID == "" {
		return ErrCustomerIDIsRequired
	}
	if s.BillingType == "" {
		return ErrBillingTypeIsRequired
	}
	if s.NextDueDate.IsZero() {
		return ErrNextDueDateIsRequired
	}
	if s.Value <= 0 {
		return ErrValueMustBePositive
	}
	if s.Cycle == "" {
		return ErrCycleIsRequired
	}
	if !s.IsBoleto() {
		return ErrOnlyBoletoAllowed
	}
	return nil
}

func (s *Subscription) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"id":          s.ID,
		"customer":    s.CustomerID,
		"billingType": s.BillingType,
		"nextDueDate": s.NextDueDate.Format("2006-01-02"),
		"value":       s.Value,
		"cycle":       s.Cycle,
		"description": s.Description,
	}
	return result
}

func (s *Subscription) Unmarshal(raw []byte) error {
	if err := json.Unmarshal(raw, s); err != nil {
		return err
	}
	s.SetNextDueDate(s.NextDue)
	if s.NextDueDate.IsZero() {
		return ErrNextDueDateIsRequired
	}
	return s.Validate()
}

func (s *Subscription) IsBoleto() bool {
	return s.BillingType == "BOLETO"
}
