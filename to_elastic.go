// Package smpsearch provides functionality to convert simple search queries to Elasticsearch format
package smpsearch

import (
	"encoding/json"
	"time"
)

// Converter defines interface for converting SimpleQuery to Elasticsearch query string
type Converter interface {
	ToElastic(search SimpleQuery) string
}

// ConverterImpl implements the Converter interface
type ConverterImpl struct {
}

var _ Converter = &ConverterImpl{}

// ToElastic converts a SimpleQuery into an Elasticsearch query string
// It handles date ranges, filters, aggregations and pagination
func (c *ConverterImpl) ToElastic(search SimpleQuery) string {
	// Create the base query structure
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{
					// Add date range
					map[string]interface{}{
						"range": map[string]interface{}{
							"@timestamp": map[string]interface{}{
								"gte": search.DateRange.From.Format(time.RFC3339),
								"lte": search.DateRange.To.Format(time.RFC3339),
							},
						},
					},
				},
			},
		},
		// Add pagination
		"from": search.Offset,
		"size": search.Limit,
	}

	// Process filters
	if len(search.Filters) > 0 {
		filters := make([]interface{}, 0, len(search.Filters))
		for _, filter := range search.Filters {
			filterQuery := c.convertFilter(filter)
			if filterQuery != nil {
				filters = append(filters, filterQuery)
			}
		}

		if len(filters) > 0 {
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
				query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}),
				filters...,
			)
		}
	}

	// Process aggregations
	if len(search.Aggregations) > 0 {
		aggs := make(map[string]interface{})
		for _, agg := range search.Aggregations {
			aggQuery := c.convertAggregation(agg)
			if aggQuery != nil {
				// Use field name as the aggregation name
				aggs[agg.Field] = aggQuery
			}
		}

		if len(aggs) > 0 {
			query["aggs"] = aggs
		}
	}

	// Convert to JSON string
	jsonBytes, err := json.Marshal(query)
	if err != nil {
		return "{}"
	}
	return string(jsonBytes)
}

// convertAggregation converts an Aggregation into its Elasticsearch representation
func (c *ConverterImpl) convertAggregation(agg Aggregation) map[string]interface{} {
	switch agg.Type {
	case TermsAggregation:
		return map[string]interface{}{
			"terms": map[string]interface{}{
				"field": agg.Field,
				"size":  agg.Size,
			},
		}
	case SumAggregation:
		return map[string]interface{}{
			"sum": map[string]interface{}{
				"field": agg.Field,
			},
		}
	case AvgAggregation:
		return map[string]interface{}{
			"avg": map[string]interface{}{
				"field": agg.Field,
			},
		}
	default:
		return nil
	}
}

// convertFilter converts a Filter into its Elasticsearch representation
func (c *ConverterImpl) convertFilter(filter Filter) map[string]interface{} {
	switch filter.Operator {
	case EqualOperator:
		return map[string]interface{}{
			"match": map[string]interface{}{
				filter.Field: filter.Value,
			},
		}
	case NotEqualOperator:
		return map[string]interface{}{
			"bool": map[string]interface{}{
				"must_not": map[string]interface{}{
					"match": map[string]interface{}{
						filter.Field: filter.Value,
					},
				},
			},
		}
	case GreaterThanOperator:
		return map[string]interface{}{
			"range": map[string]interface{}{
				filter.Field: map[string]interface{}{
					"gt": filter.Value,
				},
			},
		}
	case LessThanOperator:
		return map[string]interface{}{
			"range": map[string]interface{}{
				filter.Field: map[string]interface{}{
					"lt": filter.Value,
				},
			},
		}
	case GreaterOrEqual:
		return map[string]interface{}{
			"range": map[string]interface{}{
				filter.Field: map[string]interface{}{
					"gte": filter.Value,
				},
			},
		}
	case LessOrEqual:
		return map[string]interface{}{
			"range": map[string]interface{}{
				filter.Field: map[string]interface{}{
					"lte": filter.Value,
				},
			},
		}
	default:
		return nil
	}
}
