# Simple Search

Simple Search is a Go package that provides a simplified interface for converting structured queries into Elasticsearch query DSL. It allows you to build complex Elasticsearch queries using a more intuitive and maintainable format.

## Features

- Date range filtering
- Multiple filter types (=, !=, >, <, >=, <=)
- Aggregations support (Terms, Sum, Average)
- Pagination
- Type-safe query construction

## Installation

```bash
go get github.com/solrac97gr/smpsearch
```

## Usage

### Basic Query

```go
package main

import (
    "fmt"
    "time"
    "github.com/solrac97gr/smpsearch"
)

func main() {
    query := smpsearch.SimpleQuery{
        DateRange: smpsearch.DataRange{
            From: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
            To:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
        },
        Limit:  10,
        Offset: 0,
    }

    converter := &smpsearch.ConverterImpl{}
    elasticQuery := converter.ToElastic(query)
    fmt.Println(elasticQuery)
}
```

### Query with Filters

```go
query := smpsearch.SimpleQuery{
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
}
```

### Query with Aggregations

```go
query := smpsearch.SimpleQuery{
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
}
```

## Available Operators

- `=` (EqualOperator)
- `!=` (NotEqualOperator)
- `>` (GreaterThanOperator)
- `<` (LessThanOperator)
- `>=` (GreaterOrEqual)
- `<=` (LessOrEqual)

## Aggregation Types

- `terms` (TermsAggregation)
- `sum` (SumAggregation)
- `avg` (AvgAggregation)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Running Tests

```bash
make test
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Authors

- Carlos Garcia ([@solrac97gr](https://github.com/solrac97gr))
```

This README.md provides:
1. A clear description of the project
2. Installation instructions
3. Usage examples with different query types
4. Available operators and aggregation types
5. Contributing guidelines
6. Testing instructions
7. License information
8. Author information

You may want to add more sections like:
- Detailed API documentation
- More complex examples
- Performance considerations
- Known limitations
- Changelog
- Dependencies
