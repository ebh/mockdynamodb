# github.com/ebh/mockdynamodb

In-memory implementation of DynamoDB, ideal for unit testing

![Continuous Integration](https://github.com/ebh/mockdynamodb/workflows/Continuous%20Integration/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/ebh/mockdynamodb.svg)](https://pkg.go.dev/github.com/ebh/mockdynamodb)

## Credit

This module was inspired by and basically copies [github.com/elliotchance/mocksqs](https://github.com/elliotchance/mocksqs) by [Elliot Chance](https://github.com/elliotchance).

## Creating the Service

The simplest way to create a new DynamoDB service is with `mockdynamodb.New()`. However,
if you need tables you can use `mockdynamodb.NewWithTables()`:

```go
tableName := "foobar"
client := mocksqs.NewWithTables({tableName})

result, err := client.PutItem(&dynamodb.PutItemInput{
    TableName: aws.String(tableName),
})
```

## Supported Functionality

Only some DynamoDB methods are implemented, at moment. Methods not implemented
will panic. If you want a method implemented please [raise an issue](https://github.com/ebh/mockdynamodb/issues/new).

You can view the specific implementation details in the
[godoc documentation](https://pkg.go.dev/github.com/ebh/mockdynamodb).
