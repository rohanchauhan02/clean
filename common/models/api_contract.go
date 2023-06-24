package models

type QueryRequest struct {
	Filter map[string]interface{}         `json:"filter"`
	Order  map[string]QueryOrderDirection `json:"order"`
	Select []string                       `json:"select"`
}

type QueryFilterOperator string

const (
	QUERY_FILTER_EQ    QueryFilterOperator = "eq"
	QUERY_FILTER_NEQ   QueryFilterOperator = "neq"
	QUERY_FILTER_IN    QueryFilterOperator = "in"
	QUERY_FILTER_NIN   QueryFilterOperator = "nin"
	QUERY_FILTER_LT    QueryFilterOperator = "lt"
	QUERY_FILTER_LTE   QueryFilterOperator = "lte"
	QUERY_FILTER_GT    QueryFilterOperator = "gt"
	QUERY_FILTER_GTE   QueryFilterOperator = "gte"
	QUERY_FILTER_BW    QueryFilterOperator = "bw"
	QUERY_FILTER_INON  QueryFilterOperator = "inon"
	QUERY_FILTER_NINON QueryFilterOperator = "ninon"
)

var (
	ARITHMETIC_OPERATIONS []QueryFilterOperator = []QueryFilterOperator{
		QUERY_FILTER_LTE, QUERY_FILTER_LT, QUERY_FILTER_GTE, QUERY_FILTER_GT,
		QUERY_FILTER_BW,
	}
	ALL_OPERATIONS []QueryFilterOperator = []QueryFilterOperator{
		QUERY_FILTER_LTE, QUERY_FILTER_EQ, QUERY_FILTER_NEQ, QUERY_FILTER_LT,
		QUERY_FILTER_GTE, QUERY_FILTER_GT, QUERY_FILTER_BW, QUERY_FILTER_IN,
		QUERY_FILTER_NIN, QUERY_FILTER_INON, QUERY_FILTER_NINON,
	}
)

type QueryOrderDirection string

const (
	QUERY_ORDER_DESC QueryOrderDirection = "DESC"
	QUERY_ORDER_ASC  QueryOrderDirection = "ASC"
)

type QueryFilter map[QueryFilterOperator]interface{}
