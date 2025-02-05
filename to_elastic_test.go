package smpsearch_test

import (
	"encoding/json"
	"testing"
	"time"

	smpsearch "github.com/solrac97gr/smpsearch"
	"github.com/stretchr/testify/assert"
)

func TestConverterImpl_ToElastic(t *testing.T) {
	converter := &smpsearch.ConverterImpl{}

	tests := []struct {
		name        string
		simpleQuery smpsearch.SimpleQuery
		want        map[string]interface{}
	}{
		{
			name: "Basic query with date range only",
			simpleQuery: smpsearch.SimpleQuery{
				DateRange: smpsearch.DataRange{
					From: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					To:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				Limit:  10,
				Offset: 0,
			},
			want: map[string]interface{}{
				"query": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": []interface{}{
							map[string]interface{}{
								"range": map[string]interface{}{
									"@timestamp": map[string]interface{}{
										"gte": "2023-01-01T00:00:00Z",
										"lte": "2023-01-02T00:00:00Z",
									},
								},
							},
						},
					},
				},
				"from": float64(0),
				"size": float64(10),
			},
		},
		{
			name: "Query with filters",
			simpleQuery: smpsearch.SimpleQuery{
				DateRange: smpsearch.DataRange{
					From: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					To:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				Filters: []smpsearch.Filter{
					{
						Field:    "status",
						Operator: smpsearch.EqualOperator,
						Value:    "active",
					},
					{
						Field:    "age",
						Operator: smpsearch.GreaterThanOperator,
						Value:    "25",
					},
				},
				Limit:  10,
				Offset: 0,
			},
			want: map[string]interface{}{
				"query": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": []interface{}{
							map[string]interface{}{
								"range": map[string]interface{}{
									"@timestamp": map[string]interface{}{
										"gte": "2023-01-01T00:00:00Z",
										"lte": "2023-01-02T00:00:00Z",
									},
								},
							},
							map[string]interface{}{
								"match": map[string]interface{}{
									"status": "active",
								},
							},
							map[string]interface{}{
								"range": map[string]interface{}{
									"age": map[string]interface{}{
										"gt": "25",
									},
								},
							},
						},
					},
				},
				"from": float64(0),
				"size": float64(10),
			},
		},
		{
			name: "Query with aggregations",
			simpleQuery: smpsearch.SimpleQuery{
				DateRange: smpsearch.DataRange{
					From: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					To:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				Aggregations: []smpsearch.Aggregation{
					{
						Field: "status",
						Type:  smpsearch.TermsAggregation,
						Size:  5,
					},
					{
						Field: "amount",
						Type:  smpsearch.SumAggregation,
					},
				},
				Limit:  10,
				Offset: 0,
			},
			want: map[string]interface{}{
				"query": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": []interface{}{
							map[string]interface{}{
								"range": map[string]interface{}{
									"@timestamp": map[string]interface{}{
										"gte": "2023-01-01T00:00:00Z",
										"lte": "2023-01-02T00:00:00Z",
									},
								},
							},
						},
					},
				},
				"aggs": map[string]interface{}{
					"status": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "status",
							"size":  float64(5),
						},
					},
					"amount": map[string]interface{}{
						"sum": map[string]interface{}{
							"field": "amount",
						},
					},
				},
				"from": float64(0),
				"size": float64(10),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := converter.ToElastic(tt.simpleQuery)

			// Parse the JSON string back to map for comparison
			var gotMap map[string]interface{}
			err := json.Unmarshal([]byte(got), &gotMap)
			assert.NoError(t, err)

			// Compare the resulting maps
			assert.Equal(t, tt.want, gotMap)
		})
	}
}
