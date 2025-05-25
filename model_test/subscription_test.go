package model_test

import (
	"testing"

	"github.com/pericles-luz/go-asaas/model"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionShouldValidate(t *testing.T) {
	subscription := model.NewSubscription().
		SetCustomerID("12345").
		SetBillingType("BOLETO").
		SetNextDueDate("2023-10-01").
		SetValue(100.0).
		SetCycle("MONTHLY").
		SetDescription("Test Subscription")
	require.NoError(t, subscription.Validate(), "Subscription should be valid")
}

func TestSubscriptionShouldNotValidateWithoutCustomerID(t *testing.T) {
	subscription := model.NewSubscription().
		SetBillingType("BOLETO").
		SetNextDueDate("2023-10-01").
		SetValue(100.0).
		SetCycle("MONTHLY").
		SetDescription("Test Subscription")
	err := subscription.Validate()
	require.Error(t, err, "Subscription should not be valid without CustomerID")
	require.ErrorIs(t, model.ErrCustomerIDIsRequired, err, "Error should be ErrCustomerIDIsRequired")
}

func TestSubscriptionShouldNotValidateWithoutBillingType(t *testing.T) {
	subscription := model.NewSubscription().
		SetCustomerID("12345").
		SetNextDueDate("2023-10-01").
		SetValue(100.0).
		SetCycle("MONTHLY").
		SetDescription("Test Subscription")
	err := subscription.Validate()
	require.Error(t, err, "Subscription should not be valid without BillingType")
	require.ErrorIs(t, model.ErrBillingTypeIsRequired, err, "Error should be ErrBillingTypeIsRequired")
}

func TestSubscriptionShouldNotValidateWithoutNextDueDate(t *testing.T) {
	subscription := model.NewSubscription().
		SetCustomerID("12345").
		SetBillingType("BOLETO").
		SetValue(100.0).
		SetCycle("MONTHLY").
		SetDescription("Test Subscription")
	err := subscription.Validate()
	require.Error(t, err, "Subscription should not be valid without NextDueDate")
	require.ErrorIs(t, model.ErrNextDueDateIsRequired, err, "Error should be ErrNextDueDateIsRequired")
}

func TestSubscriptionShouldNotValidateWithZeroValue(t *testing.T) {
	subscription := model.NewSubscription().
		SetCustomerID("12345").
		SetBillingType("BOLETO").
		SetNextDueDate("2023-10-01").
		SetValue(0.0).
		SetCycle("MONTHLY").
		SetDescription("Test Subscription")
	err := subscription.Validate()
	require.Error(t, err, "Subscription should not be valid with zero Value")
	require.ErrorIs(t, model.ErrValueMustBePositive, err, "Error should be ErrValueMustBePositive")
}

func TestSubscriptionShouldNotValidateWithNegativeValue(t *testing.T) {
	subscription := model.NewSubscription().
		SetCustomerID("12345").
		SetBillingType("BOLETO").
		SetNextDueDate("2023-10-01").
		SetValue(-50.0).
		SetCycle("MONTHLY").
		SetDescription("Test Subscription")
	err := subscription.Validate()
	require.Error(t, err, "Subscription should not be valid with negative Value")
	require.ErrorIs(t, model.ErrValueMustBePositive, err, "Error should be ErrValueMustBePositive")
}

func TestSubscriptionShouldNotValidateWithoutCycle(t *testing.T) {
	subscription := model.NewSubscription().
		SetCustomerID("12345").
		SetBillingType("BOLETO").
		SetNextDueDate("2023-10-01").
		SetValue(100.0).
		SetDescription("Test Subscription")
	err := subscription.Validate()
	require.Error(t, err, "Subscription should not be valid without Cycle")
	require.ErrorIs(t, model.ErrCycleIsRequired, err, "Error should be ErrCycleIsRequired")
}

func TestSubscriptionShouldNotValidateWithInvalidNextDueDate(t *testing.T) {
	subscription := model.NewSubscription().
		SetCustomerID("12345").
		SetBillingType("BOLETO").
		SetNextDueDate("invalid-date").
		SetValue(100.0).
		SetCycle("MONTHLY").
		SetDescription("Test Subscription")
	err := subscription.Validate()
	require.Error(t, err, "Subscription should not be valid with invalid NextDueDate")
	require.ErrorIs(t, model.ErrNextDueDateIsRequired, err, "Error should be ErrNextDueDateIsRequired")
}

func TestSubscriptionShouldNotValidateIfBillingTypeIsNotBoleto(t *testing.T) {
	subscription := model.NewSubscription().
		SetCustomerID("12345").
		SetBillingType("INVALID_TYPE").
		SetNextDueDate("2023-10-01").
		SetValue(100.0).
		SetCycle("MONTHLY").
		SetDescription("Test Subscription")
	err := subscription.Validate()
	require.Error(t, err, "Subscription should not be valid with invalid BillingType")
	require.ErrorIs(t, model.ErrOnlyBoletoAllowed, err, "Error should be ErrOnlyBoletoAllowed")
}

func TestSubscriptionShouldUnmarshal(t *testing.T) {
	data := []byte(`{"object":"subscription","id":"sub_1ifrhps9m8mwficw","dateCreated":"2025-05-24","customer":"cus_000006724433","paymentLink":null,"value":100.00,"nextDueDate":"2025-07-24","cycle":"MONTHLY","description":"Monthly Subscription for John Doe","billingType":"BOLETO","deleted":false,"status":"ACTIVE","externalReference":null,"checkoutSession":null,"sendPaymentByPostalService":false,"fine":{"value":0,"type":"FIXED"},"interest":{"value":0,"type":"PERCENTAGE"},"split":null}`)
	subscription := model.NewSubscription()
	require.NoError(t, subscription.Unmarshal(data), "Subscription should unmarshal successfully")
	require.Equal(t, "sub_1ifrhps9m8mwficw", subscription.ID, "Subscription ID should match")
	require.Equal(t, "cus_000006724433", subscription.CustomerID, "Subscription CustomerID should match")
	require.Equal(t, "BOLETO", subscription.BillingType, "Subscription BillingType should match")
	require.Equal(t, 100.00, subscription.Value, "Subscription Value should match")
	require.Equal(t, "2025-07-24", subscription.NextDueDate.Format("2006-01-02"), "Subscription NextDueDate should match")
	require.Equal(t, "MONTHLY", subscription.Cycle, "Subscription Cycle should match")
	require.Equal(t, "Monthly Subscription for John Doe", subscription.Description, "Subscription Description should match")
}
