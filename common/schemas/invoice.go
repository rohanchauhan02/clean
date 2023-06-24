package schemas

import "time"

type (
	// InvoiceRequestDTO ...
	InvoiceRequestDTO struct {
		TransactionID            string   `json:"transaction_id" validate:"required"`
		PayerID                  string   `json:"payer_id"`
		PartnerID                string   `json:"partner_id"`
		SubAccountName           string   `json:"sub_account_name"`
		PayerEmail               string   `json:"payer_email" validate:"required"`
		Description              string   `json:"description" validate:"required"`
		Amount                   float64  `json:"amount" validate:"required"`
		CallbackVirtualAccountID string   `json:"callback_virtual_account_id,omitempty"`
		InvoiceDate              string   `json:"invoice_date" validate:"required"`
		InvoiceDuration          int64    `json:"invoice_duration,omitempty"` // time in Second
		ProductType              string   `json:"product_type" validate:"required"`
		ShouldSendEmail          bool     `json:"should_send_email,omitempty"`
		SuccessRedirectURL       string   `json:"success_redirect_url,omitempty"`
		PaymentMethods           []string `json:"payment_methods,omitempty"`
		PayerType                string   `json:"payer_type"`
		Flag                     string   `json:"flag" validate:"required"`
		Channel                  string   `json:"channel"`
		Country                  string   `json:"country" validate:"required"`
		Currency                 string   `json:"currency"`
	}

	// CancelInvoiceRequestDTO ...
	CancelInvoiceRequestDTO struct {
		Flag      string `json:"flag" validate:"required"`
		Channel   string `json:"channel,omitempty"`
		InvoiceID string `json:"invoice_id" validate:"required"`
	}

	// InvoiceResponseDTO ...
	InvoiceResponseDTO struct {
		InvoiceID      string          `json:"invoice_id,omitempty"`
		InvoiceURL     string          `json:"invoice_url,omitempty"`
		ExpiryDate     string          `json:"expiry_date,omitempty"`
		Token          string          `json:"token,omitempty"`
		Status         string          `json:"status"`
		ErrorMessage   string          `json:"error_message,omitempty"`
		AvailableBanks []AvailableBank `json:"available_banks,omitempty"`
	}

	// InvoiceCallbackStatus ...
	InvoiceCallbackStatus struct {
		ID                    int64      `json:"id"`
		TransactionID         string     `json:"transaction_id"`
		PartnerTransactionID  string     `json:"partner_transaction_id"`
		Amount                float64    `json:"amount"`
		Currency              string     `json:"currency"`
		AdjustedReceiveAmount float64    `json:"adjusted_receive_amount"`
		FeesPaidAmount        float64    `json:"fees_paid_amount"`
		Status                string     `json:"status"`
		PayerIdentity         string     `json:"payer_identity"`
		PayerEmail            string     `json:"payer_email"`
		PayerType             string     `json:"payer_type"`
		InvoiceDate           string     `json:"invoice_date"`
		Description           string     `json:"description"`
		PaymentMethod         string     `json:"payment_method"`
		InvoiceURL            string     `json:"invoice_url"`
		PaymentBank           string     `json:"payment_bank"`
		ProductType           string     `json:"product_type"`
		Channel               string     `json:"channel"`
		Country               string     `json:"country"`
		IsActive              int        `json:"is_active"`
		RawInvoice            string     `json:"raw_invoice"`
		RawCallback           string     `json:"raw_callback"`
		RecurringID           int64      `json:"recurring_id"`
		SourceUpdate          string     `json:"source_update"`
		CreatedAt             time.Time  `json:"created_at"`
		UpdatedAt             time.Time  `json:"updated_at"`
		DeletedAt             *time.Time `json:"deleted_at"`
		PaidAt                *time.Time `json:"paid_at"`
		DueAt                 *time.Time `json:"due_at"`
		NextScheduleAt        *time.Time `json:"next_schedule_at"`
	}

	// AvailableBank ...
	AvailableBank struct {
		BankName          string `json:"bank_name"`
		BankAccountNumber string `json:"account_number"`
	}
)
