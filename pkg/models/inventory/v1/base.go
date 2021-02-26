package v1

// Common interface that defines a type model
// Must implement `TableName()` to be treated as a Model object in QueryBuilder
type BaseModel interface {
	TableName() string
}
