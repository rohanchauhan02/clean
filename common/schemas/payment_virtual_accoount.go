package schemas

import "time"

type (
	VirtualAccountRequestDTO struct {
		PayerID        string  `json:"payer_id"`
		PayerName      string  `json:"payer_name"`
		PayerEmail     string  `json:"payer_email"`
		PayerBankCode  string  `json:"payer_bank_code"`
		Type           string  `json:"payer_type"`
		Flag           string  `json:"flag,omitempty"`
		IsClosed       bool    `json:"is_closed,omitempty"`
		ExpectedAmount float64 `json:"expected_amount,omitempty"`
		IsSingleUse    bool    `json:"is_single_use,omitempty"`
		Channel        string  `json:"channel,omitempty"`
		Country        string  `json:"country,omitempty"`
		Duration       int64   `json:"duration,omitempty"` // time in Second
	}

	// VirtualAccount ...
	VirtualAccount struct {
		ID               int64      `db:"id" gorm:"primary_key" json:"id,omitempty"`
		VirtualAccountID string     `db:"virtual_account_id" gorm:"primary_key" json:"virtual_account_id,omitempty"`
		Channel          string     `db:"channel" gorm:"channel" json:"channel,omitempty"`
		BankCode         string     `db:"bank_code" gorm:"bank_code" json:"bank_code,omitempty"`
		HolderName       string     `db:"holder_name" gorm:"holder_name" json:"holder_name,omitempty"`
		FixedAccount     string     `db:"fixed_account" gorm:"fixed_account" json:"fixed_account,omitempty"`
		Status           string     `db:"status" gorm:"status" json:"status,omitempty"`
		ExpirationDate   string     `db:"expiration_date" gorm:"expiration_date" json:"expiration_date,omitempty"`
		Type             string     `db:"type" gorm:"type" json:"type,omitempty"`
		PayerID          string     `db:"payer_id" gorm:"payer_id" json:"payer_id,omitempty"`
		HolderEmail      string     `db:"holder_email" gorm:"holder_email" json:"holder_email,omitempty"`
		IsActive         int        `db:"is_active" gorm:"is_active" json:"is_active,omitempty"`
		IsClosed         bool       `db:"is_closed" gorm:"is_closed" json:"is_closed,omitempty"`
		IsSingleUse      bool       `db:"is_single_use" gorm:"is_single_use" json:"is_single_use,omitempty"`
		ExpectedAmount   float64    `db:"expected_amount" gorm:"expected_amount" json:"expected_amount,omitempty"`
		CreatedAt        time.Time  `db:"created_at" gorm:"created_at" json:"created_at,omitempty"`
		UpdatedAt        time.Time  `db:"updated_at" gorm:"updated_at" json:"updated_at,omitempty"`
		DeletedAt        *time.Time `db:"deleted_at" gorm:"deleted_at" json:"deleted_at,omitempty"`
	}
)
