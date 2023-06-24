package schemas

import (
	"bytes"
	"errors"
	"strings"
)

type BalanceLevelType int64

const (
	BenefitLevel BalanceLevelType = iota
	PolicyLevel
	ProductLevel
)

type UnitDurationType int64

const (
	Days UnitDurationType = iota
	Months
	Years
)

type PolicyNumberingType int64

const (
	RandomAlphaNumeric PolicyNumberingType = iota
	RunningNumber
	Random
)

type FallbackDateType int64

const (
	FallbackToPreviousDate FallbackDateType = iota
	FallbackToNextDate
)

type PremiumType int64

const (
	Static PremiumType = iota
	DistinctValues
	Range
	Percentage
	StaticPartner
	Dynamic
)

type CertificateActionStrategy int64

const (
	Split CertificateActionStrategy = iota
	Merge
)

type CommissionType int64

const (
	CommissionTypePrice CommissionType = iota
	CommissionTypePercentage
)

type CommissionRecipient int64

const (
	PartnerRecipient CommissionRecipient = iota
	QoalaRecipient
	InsurerRecipient
	ServiceCenterRecipient
	QoalaAdditionalRecipient
	SalesRecipient
	GwpRecipient
)

type ProductCategoryCode int64

// travel product category
const (
	FlightCode ProductCategoryCode = iota
	TrainCode
	HotelCode
	BusCode
	ExprCode
)

type RoundingType int64

const (
	Nearest RoundingType = iota
	Up
	Down
	ToEven
	None
)

type DeductibleType int64

const (
	DeductiblePercentage DeductibleType = iota
	DeductibleStatic
)

type TravelTripType int64

const (
	SingleTrip TravelTripType = iota + 1
	RoundTripDepart
	RoundTripReturn
)

func (t BalanceLevelType) GetBalanceLevelType() (string, error) {
	switch t {
	case BenefitLevel:
		return "BENEFIT", nil
	case PolicyLevel:
		return "POLICY", nil
	case ProductLevel:
		return "PRODUCT", nil
	}

	return "", errors.New("invalid balance level type")
}

func (t BalanceLevelType) ToEnumType(s string) BalanceLevelType {
	switch s {
	case "BENEFIT":
		return BenefitLevel
	case "POLICY":
		return PolicyLevel
	case "PRODUCT":
		return ProductLevel
	}
	// default value for BalanceLevelType
	return BenefitLevel
}

func (t *BalanceLevelType) UnmarshalJSON(b []byte) error {
	*t = t.ToEnumType(strings.ReplaceAll(string(b), `"`, ""))
	return nil
}

func (t BalanceLevelType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	s, err := t.GetBalanceLevelType()
	if err != nil {
		return nil, err
	}
	buffer.WriteString(s)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t UnitDurationType) GetUnitDuration() (string, error) {
	switch t {
	case Days:
		return "DAYS", nil
	case Months:
		return "MONTHS", nil
	case Years:
		return "YEARS", nil
	}
	return "", errors.New("invalid unit duration type")
}

func (t UnitDurationType) ToEnumType(s string) UnitDurationType {
	switch s {
	case "DAYS":
		return Days
	case "MONTHS":
		return Months
	case "YEARS":
		return Years
	}
	// default value for UnitDurationType
	return Days
}

func (t *UnitDurationType) UnmarshalJSON(b []byte) error {
	*t = t.ToEnumType(strings.ReplaceAll(string(b), `"`, ""))
	return nil
}

func (t UnitDurationType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	s, err := t.GetUnitDuration()
	if err != nil {
		return nil, err
	}
	buffer.WriteString(s)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t PolicyNumberingType) GetPolicyNumberingAlgorithm() (string, error) {
	switch t {
	case Random:
		return "RANDOM", nil
	case RunningNumber:
		return "RUNNING_NUMBER", nil
	case RandomAlphaNumeric:
		return "RANDOM_ALPHANUMERIC", nil
	}
	return "", errors.New("invalid policy prefix algorithm")
}

func (t PolicyNumberingType) ToEnumType(s string) PolicyNumberingType {
	switch s {
	case "RANDOM":
		return Random
	case "RUNNING_NUMBER":
		return RunningNumber
	case "RANDOM_ALPHANUMERIC":
		return RandomAlphaNumeric
	}
	// default value for PolicyNumberingType
	return RandomAlphaNumeric
}

func (t *PolicyNumberingType) UnmarshalJSON(b []byte) error {
	*t = t.ToEnumType(strings.ReplaceAll(string(b), `"`, ""))
	return nil
}

func (t PolicyNumberingType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	s, err := t.GetPolicyNumberingAlgorithm()
	if err != nil {
		return nil, err
	}
	buffer.WriteString(s)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t FallbackDateType) GetFallbackDateType() (string, error) {
	switch t {
	case FallbackToPreviousDate:
		return "FALLBACK_TO_PREVIOUS_DATE", nil
	case FallbackToNextDate:
		return "FALLBACK_TO_NEXT_DATE", nil
	}
	return "", errors.New("invalid fallback date type")
}

func (t FallbackDateType) ToEnumType(s string) FallbackDateType {
	switch s {
	case "FALLBACK_TO_PREVIOUS_DATE":
		return FallbackToPreviousDate
	case "FALLBACK_TO_NEXT_DATE":
		return FallbackToNextDate
	}
	// default value for FallbackDateType
	return FallbackToPreviousDate
}

func (t *FallbackDateType) UnmarshalJSON(b []byte) error {
	*t = t.ToEnumType(strings.ReplaceAll(string(b), `"`, ""))
	return nil
}

func (t FallbackDateType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	s, err := t.GetFallbackDateType()
	if err != nil {
		return nil, err
	}
	buffer.WriteString(s)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t PremiumType) GetPremiumType() (string, error) {
	switch t {
	case Static:
		return "STATIC", nil
	case DistinctValues:
		return "DISTINCT_VALUES", nil
	case Percentage:
		return "PERCENTAGE", nil
	case Range:
		return "RANGE", nil
	case StaticPartner:
		return "STATIC_PARTNER", nil
	case Dynamic:
		return "DYNAMIC", nil
	}
	return "", errors.New("invalid premium type")
}

func (t PremiumType) ToEnumType(s string) PremiumType {
	switch s {
	case "STATIC":
		return Static
	case "DISTINCT_VALUES":
		return DistinctValues
	case "PERCENTAGE":
		return Percentage
	case "RANGE":
		return Range
	case "STATIC_PARTNER":
		return StaticPartner
	case "DYNAMIC":
		return Dynamic
	}
	// default value for PremiumType
	return Static
}

func (t TravelTripType) GetTravelTripType() (string, error) {
	switch t {
	case SingleTrip:
		return "SINGLETRIP", nil
	case RoundTripDepart:
		return "ROUNDTRIP-1", nil
	case RoundTripReturn:
		return "ROUNDTRIP-2", nil
	}
	return "", errors.New("invalid travel trip type")
}

func (t TravelTripType) ToEnumType(s string) TravelTripType {
	switch s {
	case "SINGLETRIP":
		return SingleTrip
	case "ROUNDTRIP-1":
		return RoundTripDepart
	case "ROUNDTRIP-2":
		return RoundTripReturn
	}
	// default value for Travel Trip Type
	return SingleTrip
}

func (t *TravelTripType) UnmarshalJSON(b []byte) error {
	*t = t.ToEnumType(strings.ReplaceAll(string(b), `"`, ""))
	return nil
}

func (t TravelTripType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	s, err := t.GetTravelTripType()
	if err != nil {
		return nil, err
	}
	buffer.WriteString(s)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *PremiumType) UnmarshalJSON(b []byte) error {
	*t = t.ToEnumType(strings.ReplaceAll(string(b), `"`, ""))
	return nil
}

func (t PremiumType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	s, err := t.GetPremiumType()
	if err != nil {
		return nil, err
	}
	buffer.WriteString(s)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t CertificateActionStrategy) GetCertificateActionStrategy() (string, error) {
	switch t {
	case Split:
		return "SPLIT", nil
	case Merge:
		return "MERGE", nil
	}
	return "", errors.New("invalid certificate action strategy")
}

func (t CertificateActionStrategy) ToEnumType(s string) CertificateActionStrategy {
	switch s {
	case "SPLIT":
		return Split
	case "MERGE":
		return Merge
	}
	// default value for CertificateActionStrategy
	return Split
}

func (t *CertificateActionStrategy) UnmarshalJSON(b []byte) error {
	*t = t.ToEnumType(strings.ReplaceAll(string(b), `"`, ""))
	return nil
}

func (t CertificateActionStrategy) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	s, err := t.GetCertificateActionStrategy()
	if err != nil {
		return nil, err
	}
	buffer.WriteString(s)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t CommissionType) GetCommissionType() (string, error) {
	switch t {
	case CommissionTypePrice:
		return "PRICE", nil
	case CommissionTypePercentage:
		return "PERCENTAGE", nil
	}
	return "", errors.New("invalid comission type")
}

func (t CommissionType) ToEnumType(s string) CommissionType {
	switch s {
	case "PRICE":
		return CommissionTypePrice
	case "PERCENTAGE":
		return CommissionTypePercentage
	}
	// default value for CommissionType
	return CommissionTypePrice
}

func (t *CommissionType) UnmarshalJSON(b []byte) error {
	*t = t.ToEnumType(strings.ReplaceAll(string(b), `"`, ""))
	return nil
}

func (t CommissionType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	s, err := t.GetCommissionType()
	if err != nil {
		return nil, err
	}
	buffer.WriteString(s)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t ProductCategoryCode) GetProductCategoryCode() (string, error) {
	switch t {
	case FlightCode:
		return "FLIGHT", nil
	case TrainCode:
		return "TRIN", nil
	case HotelCode:
		return "HTEL", nil
	case BusCode:
		return "TBUS", nil
	case ExprCode:
		return "EXPR", nil
	}
	return "", errors.New("invalid product category code")
}

func (t ProductCategoryCode) ToEnumType(s string) ProductCategoryCode {
	switch s {
	case "FLIGHT":
		return FlightCode
	case "TRIN":
		return TrainCode
	case "HTEL":
		return HotelCode
	case "TBUS":
		return BusCode
	case "EXPR":
		return ExprCode
	}
	return FlightCode
}

func (t *ProductCategoryCode) UnmarshalJSON(b []byte) error {
	*t = t.ToEnumType(strings.ReplaceAll(string(b), `"`, ""))
	return nil
}

func (t ProductCategoryCode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	s, err := t.GetProductCategoryCode()
	if err != nil {
		return nil, err
	}
	buffer.WriteString(s)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t CommissionRecipient) GetCommissionTypeRecipient() (string, error) {
	switch t {
	case PartnerRecipient:
		return "PARTNER", nil
	case QoalaRecipient:
		return "QOALA", nil
	case InsurerRecipient:
		return "INSURER", nil
	case ServiceCenterRecipient:
		return "SERVICE_CENTER", nil
	case QoalaAdditionalRecipient:
		return "QOALA_ADDITIONAL", nil
	case SalesRecipient:
		return "SALES", nil
	case GwpRecipient:
		return "GWP", nil
	}
	return "", errors.New("invalid commission type recipient")
}

func (t CommissionRecipient) ToEnumType(s string) CommissionRecipient {
	switch s {
	case "PARTNER":
		return PartnerRecipient
	case "QOALA":
		return QoalaRecipient
	case "INSURER":
		return InsurerRecipient
	case "SERVICE_CENTER":
		return ServiceCenterRecipient
	case "QOALA_ADDITIONAL":
		return QoalaAdditionalRecipient
	case "SALES":
		return SalesRecipient
	case "GWP":
		return GwpRecipient
	}
	// default value for CommissionRecipient
	return PartnerRecipient
}

func (t *CommissionRecipient) UnmarshalJSON(b []byte) error {
	*t = t.ToEnumType(strings.ReplaceAll(string(b), `"`, ""))
	return nil
}

func (t CommissionRecipient) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	s, err := t.GetCommissionTypeRecipient()
	if err != nil {
		return nil, err
	}
	buffer.WriteString(s)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t RoundingType) GetRoundingType() (string, error) {
	switch t {
	case Nearest:
		return "NEAREST", nil
	case Up:
		return "UP", nil
	case Down:
		return "DOWN", nil
	case ToEven:
		return "TO_EVEN", nil
	case None:
		return "NONE", nil
	}
	return "", errors.New("invalid rounding type")
}

func (t RoundingType) ToEnumType(s string) RoundingType {
	switch s {
	case "NEAREST":
		return Nearest
	case "UP":
		return Up
	case "DOWN":
		return Down
	case "TO_EVEN":
		return ToEven
	case "NONE":
		return None
	}
	// default value for RoundingType
	return None
}

func (t *RoundingType) UnmarshalJSON(b []byte) error {
	*t = t.ToEnumType(strings.ReplaceAll(string(b), `"`, ""))
	return nil
}

func (t RoundingType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	s, err := t.GetRoundingType()
	if err != nil {
		return nil, err
	}
	buffer.WriteString(s)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t DeductibleType) GetDeductibleType() (string, error) {
	switch t {
	case DeductibleStatic:
		return "STATIC", nil
	case DeductiblePercentage:
		return "PERCENTAGE", nil
	}
	return "", errors.New("invalid deductible type")
}

func (t DeductibleType) ToEnumType(s string) DeductibleType {
	switch s {
	case "STATIC":
		return DeductibleStatic
	case "PERCENTAGE":
		return DeductiblePercentage
	}
	// default value for DeductibleType
	return DeductibleStatic
}

func (t *DeductibleType) UnmarshalJSON(b []byte) error {
	*t = t.ToEnumType(strings.ReplaceAll(string(b), `"`, ""))
	return nil
}

func (t DeductibleType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	s, err := t.GetDeductibleType()
	if err != nil {
		return nil, err
	}
	buffer.WriteString(s)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
