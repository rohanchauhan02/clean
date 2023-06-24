package schemas

type PaymentDetail struct {
	Channel                 string                    `json:"channel,omitempty"`
	PartnerName             string                    `json:"partner_name"`
	InsuranceName           string                    `json:"insurance_name"`
	Country                 string                    `json:"country,omitempty"`
	Currency                string                    `json:"currency,omitempty"`
	PaymentFlag             string                    `json:"payment_flag,omitempty"`
	SubAccountName          string                    `json:"sub_account_name,omitempty"`
	PublicSubAccountName    string                    `json:"public_sub_account_name,omitempty"`
	SubAccountId            string                    `json:"sub_account_id,omitempty"`
	SubTransferTopic        string                    `json:"sub_transfer_topic,omitempty"`
	MasterAccountName       string                    `json:"master_account_name,omitempty"`
	PublicMasterAccountName string                    `json:"public_master_account_name,omitempty"`
	MasterAccountId         string                    `json:"master_account_id,omitempty"`
	BalanceAccountType      string                    `json:"balance_account_type,omitempty"`
	PaymentMethod           []string                  `json:"payment_methods,omitempty"`
	RoundingType            string                    `json:"rounding_type,omitempty"`
	Timezone                string                    `json:"timezone,omitempty"`
	ProductType             string                    `json:"product_type,omitempty"`
	VirtualAccountConfig    VirtualAccountConfig      `json:"virtual_account_config,omitempty"`
	InvoiceConfig           InvoiceConfig             `json:"invoice_config,omitempty"`
	DisbursementConfig      DisbursementConfig        `json:"disbursement_config,omitempty"`
	Notifications           map[string][]Notification `json:"notifications"`
	FinancePaymentConfig    FinancePaymentConfig      `json:"finance_payment_config,omitempty"`
	InsuranceLogo           string                    `json:"insurance_logo,omitempty"`
	PartnerLogo             string                    `json:"partner_logo,omitempty"`
}

type VirtualAccountConfig struct {
	Enabled     bool   `json:"enabled,omitempty"`
	Name        string `json:"name,omitempty"`
	BankCode    string `json:"bank_code,omitempty"`
	IsClosed    bool   `json:"is_closed,omitempty"`
	IsSingleUse bool   `json:"is_single_use,omitempty"`
	Duration    int64  `json:"duration,omitempty"`
}
type InvoiceConfig struct {
	Enabled           bool   `json:"enabled,omitempty"`
	ShouldSendEmail   bool   `json:"should_send_email,omitempty"`
	PayerEmail        string `json:"payer_email,omitempty"`
	Duration          int64  `json:"duration,omitempty"`
	ProductType       string `json:"product_type,omitempty"`
	BankCode          string `json:"bank_code,omitempty"`
	BankAccountName   string `json:"bank_account_name,omitempty"`
	BankAccountNumber string `json:"bank_account_number,omitempty"`
	Recipient         string `json:"recipient,omitempty"`
	TermOfPayment     string `json:"term_of_payment,omitempty"`
}
type DisbursementConfig struct {
	Enabled                   bool     `json:"enabled,omitempty"`
	BankCode                  string   `json:"bank_code,omitempty"`
	BankAccountName           string   `json:"bank_account_name,omitempty"`
	BankAccountNumber         string   `json:"bank_account_number,omitempty"`
	AvailableStatusToDisburse []string `json:"available_status_to_disburse"`
}

type SubAccountRequestDTO struct {
	TransactionID        string  `json:"transaction_id"`
	Amount               float64 `json:"amount"`
	SubAccountName       string  `json:"sub_account_name"` //destination
	SourceSubAccountName string  `json:"source_sub_account_name"`
	Flag                 string  `json:"flag"`
	Topic                string  `json:"topic"`
}

// SubAccountTransferResponseDTO ...
type SubAccountTransferResponseDTO struct {
	CreatedAt         string  `json:"created"`
	TransferID        string  `json:"transfer_id"`
	Reference         string  `json:"reference"`
	SourceUserID      string  `json:"source_user_id"`
	DestinationUserID string  `json:"destination_user_id"`
	Status            string  `json:"status"`
	Amount            float64 `json:"amount"`
}

// GetBalanceResponseDTO ...
type GetBalanceResponseDTO struct {
	Balance float64 `json:"balance"`
}

// SubAccountTransferCallbackSQS ...
type SubAccountTransferCallbackSQS struct {
	TransactionID          string  `json:"transaction_id"`
	ExternalTransactionID  string  `json:"external_transaction_id"`
	SourceAccountName      string  `json:"source_account_name"`
	DestinationAccountName string  `json:"destination_account_name"`
	Status                 string  `json:"status"`
	Amount                 float64 `json:"amount"`
}

type SubAccountTransferResponseSQS struct {
	TransactionID    string `json:"transaction_id"`
	Status           string `json:"status"`
	RequestTimestamp string `json:"request_timestamp"`
}

type FinancePaymentConfig struct {
	StoplossConfig map[string]interface{} `json:"stoploss,omitempty"`
}
