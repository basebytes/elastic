package index

import "github.com/basebytes/elastic-go/client/entity"

// Index index define
type Index interface {
	Name() string
	IsNestedField(field string) (parent string, nested bool)
	Skip(field string) bool
	QueryField(field string) string
	FieldTermSize(field string) int
	TransAggs(aggs *entity.Aggregations, lastField string) ([]map[string]any, byte)
}

// Field index field define
type Field interface {
	Name() string
	DateInterval() string
	DataInterval() int
	Missing() any
	MinDocCount() int
	IsNestedField() (parent string, nested bool)
	Group(next map[string]any) map[string]any
	Statistics() map[string]any
}
