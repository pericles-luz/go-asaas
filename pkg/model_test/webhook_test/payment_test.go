package webhooktest

import (
	"testing"

	"github.com/pericles-luz/go-asaas/pkg/model/webhook"
	"github.com/stretchr/testify/require"
)

func TestPaymentShouldUnmarshal(t *testing.T) {
	data := []byte(`{
   "id": "evt_05b708f961d739ea7eba7e4db318f621&368604920",
   "event":"PAYMENT_RECEIVED",
   "dateCreated": "2024-06-12 16:45:03",
   "payment":{
      "object":"payment",
      "id":"pay_080225913252",
      "dateCreated":"2021-01-01",
      "customer":"cus_G7Dvo4iphUNk",
      "subscription":"sub_VXJBYgP2u0eO",  
      "installment":"2765d086-c7c5-5cca-898a-4262d212587c",
      "paymentLink":"123517639363",
      "dueDate":"2021-01-01",
      "originalDueDate":"2021-01-01",
      "value":100.01,
      "netValue":94.51,
      "originalValue":null,
      "interestValue":null,
      "nossoNumero": null,
      "description":"Pedido 056984",
      "externalReference":"056984",
      "billingType":"CREDIT_CARD",
      "status":"RECEIVED",
      "pixTransaction":null,
      "confirmedDate":"2021-01-01",
      "paymentDate":"2021-01-01",
      "clientPaymentDate":"2021-01-01",
      "installmentNumber": null,
      "creditDate":"2021-02-01",
      "custody": null,
      "estimatedCreditDate":"2021-02-01",
      "invoiceUrl":"https://www.asaas.com/i/080225913252",
      "bankSlipUrl":null,
      "transactionReceiptUrl":"https://www.asaas.com/comprovantes/4937311816045162",
      "invoiceNumber":"00005101",
      "deleted":false,
      "anticipated":false,
      "anticipable":false,
      "lastInvoiceViewedDate":"2021-01-01 12:54:56",
      "lastBankSlipViewedDate":null,
      "postalService":false,
      "creditCard":{
         "creditCardNumber":"8829",
         "creditCardBrand":"MASTERCARD",
         "creditCardToken":"a75a1d98-c52d-4a6b-a413-71e00b193c99"
      },
      "discount":{
         "value":0.00,
         "dueDateLimitDays":0,
         "limitedDate": null,
         "type":"FIXED"
      },
      "fine":{
         "value":0.00,
         "type":"FIXED"
      },
      "interest":{
         "value":0.00,
         "type":"PERCENTAGE"
      },
      "split":[
         {
            "id": "c788f2e1-0a5b-41b9-b0be-ff3641fb0cbe",
            "walletId":"48548710-9baa-4ec1-a11f-9010193527c6",
            "fixedValue":20,
            "status":"PENDING",
            "refusalReason": null,
            "externalReference": null,
            "description": null
         },
         {
            "id": "e754f2e1-09mn-88pj-l552-df38j1fbll1c",
            "walletId":"0b763922-aa88-4cbe-a567-e3fe8511fa06",
            "percentualValue":10,
            "status":"PENDING",
            "refusalReason": null,
            "externalReference": null,
            "description": null
         }
      ],
      "chargeback": {
          "status": "REQUESTED",
          "reason": "PROCESS_ERROR"
      },
      "refunds": null
   }
}`)
	entity := webhook.NewWebhookPayment()
	require.NoError(t, entity.Unmarshal(data), "should unmarshal webhook data")
	require.Equal(t, "evt_05b708f961d739ea7eba7e4db318f621&368604920", entity.EventID)
	require.Equal(t, "PAYMENT_RECEIVED", entity.Event)
	require.True(t, entity.IsPaid(), "should be a paid event")
	require.Equal(t, "2024-06-12 16:45:03", entity.DateCreated)
	require.Equal(t, 10001, entity.ValueAsInt())
	require.Equal(t, 10001, entity.Amount())
	require.Equal(t, "pay_080225913252", entity.ID())
	require.Equal(t, "2021-01-01", entity.PaymentDate())
	require.Equal(t, "sub_VXJBYgP2u0eO", entity.SubscriptionID())
}
