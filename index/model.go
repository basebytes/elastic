package index

import "github.com/basebytes/elastic-go/client/entity"

type Index interface {
	Name() string
	IsNestedField(field string) (parent string, nested bool)
	Skip(field string) bool
	QueryField(field string) string
	FieldTermSize(field string) int
	TransAggs(aggs *entity.Aggregations, lastField string) ([]map[string]any, byte)
}

const Item = "item"
