package model_test

import (
	"testing"

	"github.com/pericles-luz/go-asaas/pkg/model"
	"github.com/stretchr/testify/require"
)

func TestCustomerShouldValidateWithMobilePhone(t *testing.T) {
	customer := model.NewCustomer()
	customer.SetName("John Doe").
		SetCpfCnpj("00000000191").
		SetMobilePhone("11999999999")
	require.NoError(t, customer.Validate(), "Customer should be valid with mobile phone")
}

func TestCustomerShouldValidateWithEmail(t *testing.T) {
	customer := model.NewCustomer()
	customer.SetName("Jane Doe").
		SetCpfCnpj("00000000192").
		SetEmail("teste@testando.com")
	require.NoError(t, customer.Validate(), "Customer should be valid with email")
}

func TestCustomerShouldNotValidateWithoutContactInfo(t *testing.T) {
	customer := model.NewCustomer()
	customer.SetName("No Contact").
		SetCpfCnpj("00000000193")
	err := customer.Validate()
	require.Error(t, err, "Customer should not be valid without contact info")
	require.ErrorIs(t, model.ErrNoContactInfo, err, "Error should be ErrNoContactInfo")
}

func TestCustomerShouldNotValidateWithoutName(t *testing.T) {
	customer := model.NewCustomer()
	customer.SetCpfCnpj("00000000194").
		SetMobilePhone("11999999999")
	err := customer.Validate()
	require.Error(t, err, "Customer should not be valid without name")
	require.ErrorIs(t, model.ErrNameIsRequired, err, "Error should be ErrNameIsRequired")
}

func TestCustomerShouldNotValidateWithoutDocument(t *testing.T) {
	customer := model.NewCustomer()
	customer.SetName("No Document").
		SetMobilePhone("11999999999")
	err := customer.Validate()
	require.Error(t, err, "Customer should not be valid without document")
	require.ErrorIs(t, model.ErrDocumentIsRequired, err, "Error should be ErrDocumentIsRequired")
}

func TestCustomerShouldUnmarshal(t *testing.T) {
	customer := model.NewCustomer()
	data := []byte(`{"object":"customer","id":"cus_000006724433","dateCreated":"2025-05-24","name":"John Doe","email":null,"company":null,"phone":null,"mobilePhone":"31986058910","address":null,"addressNumber":null,"complement":null,"province":null,"postalCode":null,"cpfCnpj":"00000000191","personType":"FISICA","deleted":false,"additionalEmails":null,"externalReference":null,"notificationDisabled":false,"observations":null,"municipalInscription":null,"stateInscription":null,"canDelete":true,"cannotBeDeletedReason":null,"canEdit":true,"cannotEditReason":null,"city":null,"cityName":null,"state":null,"country":"Brasil"}`)
	require.NoError(t, customer.Unmarshal(data), "Customer should unmarshal successfully")
	require.Equal(t, "John Doe", customer.Name, "Customer name should match")
	require.Equal(t, "00000000191", customer.CpfCnpj, "Customer CPF/CNPJ should match")
	require.Equal(t, "31986058910", customer.MobilePhone, "Customer mobile phone should match")
	require.Equal(t, "FISICA", customer.PersonType, "Customer person type should match")
}
