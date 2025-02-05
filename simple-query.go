// Package smpsearch provides simple search query functionality
package smpsearch

import "time"

// SimpleQuery represents a search query structure with date range, filters,
// aggregations, and pagination parameters
type SimpleQuery struct {
	// DateRange specifies the time period for the search
	DateRange DataRange `json:"date_range" validate:"required,dive,required"`

	// Filters contains conditions to narrow down search results
	Filters []Filter `json:"filters" validate:"dive"`

	// Aggregations defines grouping and calculation operations on the data
	Aggregations []Aggregation `json:"aggregations" validate:"dive"`

	// Limit specifies maximum number of results to return
	Limit int `json:"limit" validate:"required"`

	// Offset specifies number of results to skip for pagination
	Offset int `json:"offset" validate:"required"`
}

// DataRange defines a time period with start and end times
type DataRange struct {
	// From is the start time of the range
	From time.Time `json:"from" validate:"required"`

	// To is the end time of the range
	To time.Time `json:"to" validate:"required"`
}

// Aggregation defines how to group or calculate statistics on the search results
type Aggregation struct {
	// Field specifies which field to aggregate on
	Field string `json:"field" validate:"required"`

	// Type specifies the kind of aggregation to perform
	Type AggregationType `json:"type" validate:"required"`

	// Size specifies the maximum number of aggregation buckets to return
	Size int `json:"size"`
}

// AggregationType represents the type of aggregation operation
type AggregationType string

// Supported aggregation types
const (
	// TermsAggregation groups results by field values
	TermsAggregation AggregationType = "terms"

	// SumAggregation calculates the sum of numeric values
	SumAggregation AggregationType = "sum"

	// AvgAggregation calculates the average of numeric values
	AvgAggregation AggregationType = "avg"
)

// Filter defines a condition to filter search results
type Filter struct {
	// Field specifies which field to apply the filter on
	Field string `json:"field" validate:"required"`

	// Operator specifies the comparison operation
	Operator Operator `json:"operator" validate:"required"`

	// Value specifies the comparison value
	Value string `json:"value" validate:"required"`
}

// Operator represents comparison operations for filters
type Operator string

// Supported filter operators
const (
	// EqualOperator represents equality comparison (=)
	EqualOperator Operator = "="

	// NotEqualOperator represents inequality comparison (!=)
	NotEqualOperator Operator = "!="

	// GreaterThanOperator represents greater than comparison (>)
	GreaterThanOperator Operator = ">"

	// LessThanOperator represents less than comparison (<)
	LessThanOperator Operator = "<"

	// GreaterOrEqual represents greater than or equal comparison (>=)
	GreaterOrEqual Operator = ">="

	// LessOrEqual represents less than or equal comparison (<=)
	LessOrEqual Operator = "<="
)
