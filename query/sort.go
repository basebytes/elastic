package query

import (
	"strings"

	"github.com/basebytes/elastic/index"
)

type Sort struct {
	Sorts []string `json:"sorts,omitempty"`
}

func (s *Sort) Sort(index index.Index) (sorts []map[string]any) {
	for _, _sort := range s.Sorts {
		sorts = append(sorts, generateSort(index, _sort))
	}
	return
}

func generateSort(index index.Index, sort string) map[string]any {
	result := make(map[string]any, 1)
	orderBy := strings.Split(sort, " ")
	field, _order := index.QueryField(orderBy[0]), defaultSort
	if len(orderBy) > 1 && strings.ToLower(orderBy[1]) == sortAsc {
		_order = sortAsc
	}
	if p, nested := index.IsNestedField(orderBy[0]); nested {
		params := map[string]any{
			"nested": map[string]any{
				"path": p,
			},
		}
		params["order"] = _order
		result[field] = params
	} else {
		result[field] = _order
	}
	return result
}

const (
	defaultSort = sortDesc
	sortAsc     = "asc"
	sortDesc    = "desc"
)
