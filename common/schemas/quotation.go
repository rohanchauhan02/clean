package schemas

import (
	"encoding/json"
	"time"
)

type CreateQuotationRequest struct {
	ProductCategory          string               `json:"product_category"` // HEALTH
	QoalaPolicyNumber        *string              `json:"qoala_policy_number"`
	Status                   *string              `json:"status,omitempty"`
	PartnerTransactionNumber *string              `json:"partner_transaction_number" validate:"required"`
	InsuranceNumber          *string              `json:"insurance_number"`
	ProductCode              string               `json:"product_code"`
	StartProtectionAt        time.Time            `json:"start_protection_at"` // format: YYYY-MM-DD
	EndProtectionAt          time.Time            `json:"end_protection_at"`   // format: YYYY-MM-DD
	PurchasedAt              *time.Time           `json:"purchased_at"`        // format: YYYY-MM-DD
	Premium                  *float64             `json:"premium"`
	PoolAmount               *float64             `json:"pool_amount"`
	Discount                 *float64             `json:"discount"`
	Addons                   []AddonRequest       `json:"addons"`
	Booker                   *PolicyHolderRequest `json:"booker"`
	PolicyHolder             *PolicyHolderRequest `json:"policy_holder"`
	Insureds                 []InsuredRequest     `json:"insureds"`
}

type InsuredRequest struct {
	Type      string            `json:"type"`
	Details   interface{}       `json:"details"`
	Documents []DocumentRequest `json:"documents"`
}

type UserInsuredDetail struct {
	IdentityType   string          `json:"identity_type"`
	IdentityNumber string          `json:"identity_number"`
	FullName       string          `json:"full_name"`
	Email          string          `json:"email"`
	PhoneNumber    string          `json:"phone_number"`
	BirthDate      string          `json:"birth_date"`
	Gender         string          `json:"gender"`
	Relationship   string          `json:"relationship"`
	Address        string          `json:"address"`
	Occupation     string          `json:"occupation"`
	AdditionalInfo json.RawMessage `json:"additional_info,omitempty"`
}

type VehicleInsuredDetail struct {
	PlateCode          string  `json:"plate_code"`
	PlateNumber        string  `json:"plate_number"`
	Category           string  `json:"category"`
	Manufacturer       string  `json:"manufacturer"`
	Brand              string  `json:"brand"`
	Series             string  `json:"series"`
	Year               int     `json:"year"`
	Price              float64 `json:"price"`
	AccessoriesPrice   float64 `json:"accessories_price"`
	AccessoriesNote    string  `json:"accessories_note"`
	UsageType          string  `json:"usage_type"`
	Condition          string  `json:"condition"`
	Protection         string  `json:"protection"`
	ChassisNumber      string  `json:"chassis_number"`
	EngineNumber       string  `json:"engine_number"`
	RegistrationNumber string  `json:"registration_number"`
}

// We'll remove this later, and will use TravelInsuredDetail
type FlightInsuredDetail struct {
	IdentityType        string  `json:"identity_type"`
	IdentityNumber      string  `json:"identity_number"`
	BookingTime         string  `json:"booking_time"`
	Name                string  `json:"name"`
	Number              string  `json:"number"`
	Departure           string  `json:"departure"`
	DepartureTime       string  `json:"departure_time"`
	Arrival             string  `json:"arrival"`
	ArrivalTime         string  `json:"arrival_time"`
	Class               string  `json:"class"`
	SubClass            string  `json:"sub_class"`
	CityOrigin          string  `json:"city_origin"`
	CityDestination     string  `json:"city_destination"`
	CountryOrigin       string  `json:"country_origin"`
	CountryDestination  string  `json:"country_destination"`
	TimezoneOrigin      string  `json:"timezone_origin"`
	TimezoneDestination string  `json:"timezone_destination"`
	TotalPriceTicket    float64 `json:"total_price_ticket,omitempty"`
	SubTotalPriceTicket float64 `json:"sub_total_price_ticket,omitempty"`
}

type TravelInsuredDetail struct {
	IdentityType        string          `json:"identity_type"`
	IdentityNumber      string          `json:"identity_number"`
	BookingTime         string          `json:"booking_time"`
	Name                string          `json:"name"`
	Number              string          `json:"number"`
	Departure           string          `json:"departure"`
	DepartureTime       string          `json:"departure_time"`
	Arrival             string          `json:"arrival"`
	ArrivalTime         string          `json:"arrival_time"`
	Class               string          `json:"class"`
	SubClass            string          `json:"sub_class"`
	CityOrigin          string          `json:"city_origin"`
	CityDestination     string          `json:"city_destination"`
	CountryOrigin       string          `json:"country_origin"`
	CountryDestination  string          `json:"country_destination"`
	TimezoneOrigin      string          `json:"timezone_origin"`
	TimezoneDestination string          `json:"timezone_destination"`
	TotalPriceTicket    float64         `json:"total_price_ticket,omitempty"`
	SubTotalPriceTicket float64         `json:"sub_total_price_ticket,omitempty"`
	TripType            TravelTripType  `json:"trip_type,omitempty"`
	TripOrder           int             `json:"trip_order,omitempty"`
	AdditionalInfo      json.RawMessage `json:"additional_info,omitempty"`
	NumberOfRoom        int             `json:"number_of_room,omitempty"`
}

type LoanInsuredDetail struct {
	IdentityType             string      `json:"identity_type"`
	IdentityNumber           string      `json:"identity_number"`
	LoanAmount               float64     `json:"loan_amount"`
	TotalPrincipal           float64     `json:"loan_total_principal,omitempty"`
	TotalPrincipalPercentage float64     `json:"loan_total_principal_percentage,omitempty"`
	TenureValue              float64     `json:"tenure_value"`
	TenurePeriod             string      `json:"tenure_period"`
	InterestRate             float64     `json:"interest_rate,omitempty"`
	InterestPeriod           string      `json:"interest_period,omitempty"`
	AdditionalInfo           interface{} `json:"additional_info"`
}

type GadgetInsuredDetail struct {
	IdentityType       string               `json:"identity_type"`
	IdentityNumber     string               `json:"identity_number"`
	SumInsured         float64              `json:"sum_insured"`
	Manufacturer       string               `json:"manufacturer"`
	Brand              string               `json:"brand"`
	Series             string               `json:"series"`
	Model              string               `json:"model"`
	DeductibleInfo     GadgetDeductibleInfo `json:"deductible_info"`
	ActivationNumber   string               `json:"activation_number,omitempty"`
	ExternalURL        string               `json:"external_url,omitempty"`
	UploadActivationAt *time.Time           `json:"upload_activation_at,omitempty"`
	DeviceCategory     string               `json:"device_category"`
	DevicePrice        float64              `json:"device_price,omitempty"`
}

type GadgetDeductibleInfo struct {
	UnitFee       DeductibleType `json:"unit_fee"`
	UnitFeeAmount float64        `json:"unit_fee_amount"`
}

type DocumentRequest struct {
	FileName string `json:"file_name"`
	Alias    string `json:"alias"`
	Type     string `json:"type"`
	URL      string `json:"url"`
}

type PolicyHolderRequest struct {
	IdentityType   *string           `json:"identity_type"`
	IdentityNumber *string           `json:"identity_number"`
	FullName       string            `json:"full_name"`
	Email          *string           `json:"email"`
	PhoneNumber    *string           `json:"phone_number"`
	BirthDate      *string           `json:"birth_date"`   // format YYYY-MM-DD
	Gender         *string           `json:"gender"`       // MALE / FEMALE
	Relationship   *string           `json:"relationship"` // SELF
	Address        *string           `json:"address"`
	Documents      []DocumentRequest `json:"documents"`
}

type AddonRequest struct {
	Code       string  `json:"code"`
	Unit       int     `json:"unit"`
	SumInsured float64 `json:"sum_insured"`
}

type CreateQuotationResponse struct {
	PartnerTransactionNumber *string                         `json:"partner_transaction_number,omitempty"`
	TransactionNumber        string                          `json:"transaction_number"`
	ProductCode              string                          `json:"product_code"`
	Policies                 []CreateQuotationPolicyResponse `json:"policies,omitempty"`
}

type CreateQuotationPolicyResponse struct {
	Number            string             `json:"number"`
	Status            string             `json:"status"`
	StartProtectionAt time.Time          `json:"start_protection_at"`
	EndProtectionAt   time.Time          `json:"end_protection_at"`
	Documents         []DocumentResponse `json:"documents"`
}

type DocumentResponse struct {
	Type     string `json:"type"`
	FileName string `json:"file_name"`
	Alias    string `json:"alias"`
	URL      string `json:"url"`
}

type MotorInsuredDetail struct {
	IdentifierType     string  `json:"identifier_type"`
	IdentifierNumber   string  `json:"identifier_number"`
	Manufacturer       string  `json:"manufacturer"`
	Brand              string  `json:"brand"`
	Series             string  `json:"series"`
	Condition          string  `json:"condition"`
	ChasisNumber       string  `json:"chasis_number"`
	EngineNumber       string  `json:"engine_number"`
	RegistrationNumber string  `json:"registration_number"`
	Year               int     `json:"year"`
	Price              float64 `json:"price"`
}

type PostProductDetailDTO struct {
	PartnerTransactionId      string                            `json:"partner_transaction_id"`
	InsuredAmount             float64                           `json:"insured_amount"`
	CoveredUsers              []PostProductDetailCoveredUserDTO `json:"covered_users"`
	StartProtectionAt         time.Time                         `json:"start_protection_at"`
	EndProtectionAt           time.Time                         `json:"end_protection_at,omitempty"`
	PremiumUnitDuration       string                            `json:"premium_unit_duration"`
	PremiumUnitDurationAmount float64                           `json:"premium_unit_duration_amount"`
	Idempotent                bool                              `json:"idempotent,omitempty"`
	Premium                   float64                           `json:"premium"`
}

type PostProductDetailCoveredUserDTO struct {
	Type           string  `json:"type"`
	IdentityType   *string `json:"identity_type"`
	IdentityNumber *string `json:"identity_number"`
	FullName       string  `json:"full_name"`
	Email          *string `json:"email"`
	PhoneNumber    *string `json:"phone_number"`
	DateOfBirth    string  `json:"date_of_birth"`
	Gender         *string `json:"gender"`
	Relationship   *string `json:"relationship"`
	Address        *string `json:"address"`
}

type FinanceAdditionalInfoProductDetail struct {
	LoanDetails     LoanDetails   `json:"loan_details,omitempty"`
	BorrowerDetails interface{}   `json:"borrower_details,omitempty"`
	LenderDetails   LenderDetails `json:"lender_details,omitempty"`
}

type LenderDetails struct {
	BeneficiaryName string `json:"beneficiary_name"`
}

type LoanDetails struct {
	PremiRate              float64 `json:"premium_rate,omitempty"`
	FallbackTenureNextDate float64 `json:"fallback_tenure_next_date,omitempty"`
}
