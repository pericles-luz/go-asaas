package rest_asaas_test

import (
	"os"
	"testing"
	"time"

	"github.com/pericles-luz/go-asaas/factory/factory_client_asaas"
	"github.com/pericles-luz/go-asaas/model"
	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestRestShouldCreateCustomer(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	restEntity, err := factory_client_asaas.NewClient(utils.GetBaseDirectory("config") + "/sandbox.json")
	require.NoError(t, err, "Failed to create rest entity")
	customer := model.NewCustomer().
		SetName("John Doe").
		SetCpfCnpj("00000000191").
		SetMobilePhone("31999999999") // replace with a valid mobile phone number
	created, err := restEntity.CreateCustomer(customer)
	require.NoError(t, err, "Failed to create customer")
	require.Equal(t, customer.Name, created.Name, "Customer name should match")
	require.Equal(t, customer.CpfCnpj, created.CpfCnpj, "Customer CPF/CNPJ should match")
	require.Equal(t, customer.MobilePhone, created.MobilePhone, "Customer mobile phone should match")
	require.NotEmpty(t, created.ID, "Customer ID should not be empty")
}

func TestRestShouldRetrieveCustomer(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	restEntity, err := factory_client_asaas.NewClient(utils.GetBaseDirectory("config") + "/sandbox.json")
	require.NoError(t, err, "Failed to create rest entity")
	customerID := "cus_000006724433" // Replace with a valid customer ID
	customer, err := restEntity.GetCustomer(customerID)
	require.NoError(t, err, "Failed to retrieve customer")
	require.Equal(t, customerID, customer.ID, "Customer ID should match")
	require.NotEmpty(t, customer.Name, "Customer name should not be empty")
}

func TestRestShouldListCustomers(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	restEntity, err := factory_client_asaas.NewClient(utils.GetBaseDirectory("config") + "/sandbox.json")
	require.NoError(t, err, "Failed to create rest entity")
	customers, err := restEntity.ListCustomers(map[string]interface{}{
		"name": "John Doe", // Adjust the filter as needed
	})
	require.NoError(t, err, "Failed to list customers")
	require.NotEmpty(t, customers, "Customers list should not be empty")
}

func TestRestShouldSubscribe(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	restEntity, err := factory_client_asaas.NewClient(utils.GetBaseDirectory("config") + "/sandbox.json")
	require.NoError(t, err, "Failed to create rest entity")
	subscription := model.NewSubscription().
		SetCustomerID("cus_000006724433"). // Replace with a valid customer ID
		SetBillingType(model.BILLING_TYPE_BOLETO).
		SetNextDueDate(time.Now().AddDate(0, 1, 0).Format("2006-01-02")).
		SetValue(100.00).
		SetCycle(model.CYCLE_MONTHLY).
		SetDescription("Monthly Subscription for John Doe")
	createdSubscription, err := restEntity.Subscribe(subscription)
	require.NoError(t, err, "Failed to create subscription")
	require.Equal(t, subscription.CustomerID, createdSubscription.CustomerID, "Subscription customer ID should match")
	require.Equal(t, subscription.BillingType, createdSubscription.BillingType, "Subscription billing type should match")
	require.Equal(t, subscription.Value, createdSubscription.Value, "Subscription value should match")
	require.Equal(t, subscription.Cycle, createdSubscription.Cycle, "Subscription cycle should match")
	require.NotEmpty(t, createdSubscription.ID, "Subscription ID should not be empty")
	require.NotEmpty(t, createdSubscription.NextDueDate, "Subscription next due date should not be empty")
	require.Equal(t, subscription.Description, createdSubscription.Description, "Subscription description should match")
}

func TestRestShouldUnsubscribe(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "yes" {
		t.Skip("Skip test in GitHub Actions")
	}
	restEntity, err := factory_client_asaas.NewClient(utils.GetBaseDirectory("config") + "/sandbox.json")
	require.NoError(t, err, "Failed to create rest entity")
	subscriptionID := "sub_1ifrhps9m8mwficw" // Replace with a valid subscription ID
	err = restEntity.Unsubscribe(subscriptionID)
	require.NoError(t, err, "Failed to unsubscribe")
}
