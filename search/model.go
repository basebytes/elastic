package search

import (
	"github.com/basebytes/elastic-go/client/entity"
	"github.com/basebytes/elastic/index"
)

type Query interface {
	Name() string
	InvalidParam() bool
	Build() []byte
	ParseQueryResult(result *entity.EsQueryResult) (QueryResult, int64, byte)
}

type QueryResult []map[string]any

type Sort interface {
	Sort(index index.Index) []map[string]any
}

type ReturnFiled interface {
	Source(index index.Index) map[string]any
}

type Page interface {
	From() int
	Size() int
}
