package filter

import "github.com/basebytes/elastic/fields"

type Filter interface {
	Filters() []map[string]any
	CheckNumberRange() bool
}

type NumberRanges struct {
	Start int64 `json:"start,omitempty"`
	End   int64 `json:"end,omitempty"`
}

type QueryBuilder func(filters, not Filter) map[string]any

type AggBuilder func(group, fields fields.Fields, extend fields.Extend) (map[string]any, string)
