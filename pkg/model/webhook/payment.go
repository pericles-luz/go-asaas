package webhook

import "encoding/json"

type WebhookPayment struct {
	EventID     string `json:"id"`
	Event       string `json:"event"`
	DateCreated string `json:"dateCreated"`
	Payment     struct {
		Object                 string  `json:"object"`
		ID                     string  `json:"id"`
		DateCreated            string  `json:"dateCreated"`
		Customer               string  `json:"customer"`
		Subscription           string  `json:"subscription,omitempty"` // only when part of a subscription
		Installment            string  `json:"installment,omitempty"`  // only when part of an installment
		PaymentLink            string  `json:"paymentLink,omitempty"`  // identifier of the payment link
		DueDate                string  `json:"dueDate"`
		OriginalDueDate        string  `json:"originalDueDate"`
		Value                  float64 `json:"value"`
		NetValue               float64 `json:"netValue"`
		OriginalValue          float64 `json:"originalValue,omitempty"` // when the paid value differs from the charge value
		InterestValue          float64 `json:"interestValue,omitempty"`
		NossoNumero            string  `json:"nossoNumero,omitempty"`
		Description            string  `json:"description"`
		ExternalReference      string  `json:"externalReference"`
		BillingType            string  `json:"billingType"`
		Status                 string  `json:"status"`
		PixTransaction         string  `json:"pixTransaction,omitempty"`
		ConfirmedDate          string  `json:"confirmedDate"`
		PaymentDate            string  `json:"paymentDate"`
		ClientPaymentDate      string  `json:"clientPaymentDate"`
		InstallmentNumber      int     `json:"installmentNumber,omitempty"`
		CreditDate             string  `json:"creditDate"`
		Custody                string  `json:"custody,omitempty"`
		EstimatedCreditDate    string  `json:"estimatedCreditDate"`
		InvoiceURL             string  `json:"invoiceUrl"`
		BankSlipURL            string  `json:"bankSlipUrl,omitempty"`
		TransactionReceiptURL  string  `json:"transactionReceiptUrl"`
		InvoiceNumber          string  `json:"invoiceNumber"`
		Deleted                bool    `json:"deleted"`
		Anticipated            bool    `json:"anticipated"`
		Anticipable            bool    `json:"anticipable"`
		LastInvoiceViewedDate  string  `json:"lastInvoiceViewedDate"`
		LastBankSlipViewedDate string  `json:"lastBankSlipViewedDate,omitempty"`
		PostalService          bool    `json:"postalService"`
		CreditCard             struct {
			CreditCardNumber string `json:"creditCardNumber"`
			CreditCardBrand  string `json:"creditCardBrand"`
			CreditCardToken  string `json:"creditCardToken"`
		} `json:"creditCard"`
		Discount struct {
			Value            float64 `json:"value"`
			DueDateLimitDays int     `json:"dueDateLimitDays"`
			LimitedDate      string  `json:"limitedDate,omitempty"`
			Type             string  `json:"type"`
		} `json:"discount"`
		Fine struct {
			Value float64 `json:"value"`
			Type  string  `json:"type"`
		} `json:"fine"`
		Interest struct {
			Value float64 `json:"value"`
			Type  string  `json:"type"`
		} `json:"interest"`
		Split []struct {
			ID                string  `json:"id"`
			WalletID          string  `json:"walletId"`
			FixedValue        float64 `json:"fixedValue,omitempty"`
			PercentualValue   float64 `json:"percentualValue,omitempty"`
			Status            string  `json:"status"`
			RefusalReason     string  `json:"refusalReason,omitempty"`
			ExternalReference string  `json:"externalReference,omitempty"`
			Description       string  `json:"description,omitempty"`
		} `json:"split"`
		Chargeback struct {
			Status string `json:"status"`
			Reason string `json:"reason"`
		} `json:"chargeback,omitempty"`
		Refunds []struct {
			ID          string  `json:"id"`
			Value       float64 `json:"value"`
			Description string  `json:"description"`
			Status      string  `json:"status"`
			DateCreated string  `json:"dateCreated"`
		} `json:"refunds,omitempty"`
	} `json:"payment"`
}

func NewWebhookPayment() *WebhookPayment {
	return &WebhookPayment{}
}

func (w *WebhookPayment) Unmarshal(data []byte) error {
	return json.Unmarshal(data, w)
}

func (w *WebhookPayment) IsPaid() bool {
	return w.Event == "PAYMENT_RECEIVED"
}

func (w *WebhookPayment) IsCancelled() bool {
	return w.Event == "PAYMENT_DELETED"
}

func (w *WebhookPayment) IsOpen() bool {
	return w.Event == "PAYMENT_CREATED"
}

func (w *WebhookPayment) IsOverdue() bool {
	return w.Event == "PAYMENT_OVERDUE"
}

func (w *WebhookPayment) ValueAsInt() int {
	return int(w.Payment.Value * 100)
}

func (w *WebhookPayment) ID() string {
	return w.Payment.ID
}

func (w *WebhookPayment) Amount() int {
	return int(w.Payment.Value * 100)
}

func (w *WebhookPayment) PaymentDate() string {
	if w.Payment.PaymentDate != "" {
		return w.Payment.PaymentDate
	}
	return w.Payment.ClientPaymentDate
}
