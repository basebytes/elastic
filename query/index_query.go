package query

import (
	"github.com/basebytes/elastic-go/client/entity"
	"github.com/basebytes/elastic/fields"
	"github.com/basebytes/elastic/filter"
	"github.com/basebytes/elastic/index"
	"github.com/basebytes/elastic/search"
	"github.com/basebytes/tools"
)

func NewIndexQuery(index index.Index, filters, not filter.Filter, extend fields.Extend, group, stats fields.Fields,
	queryBuilder filter.QueryBuilder, aggBuilder filter.AggBuilder, fieldsNotReturn ...string) *IndexQuery {
	return &IndexQuery{
		Filters:    filters,
		Not:        not,
		Group:      group,
		Statistics: stats,
		internal: &internal{
			idx:                   index,
			extend:                extend,
			queryBuilder:          queryBuilder,
			aggBuilder:            aggBuilder,
			defaultQueryNotReturn: fieldsNotReturn,
		},
	}
}

type IndexQuery struct {
	Page
	*Sort
	*ReturnFiled
	Filters    filter.Filter `json:"filters,omitempty"`
	Not        filter.Filter `json:"not,omitempty"`
	Group      fields.Fields `json:"fields,omitempty"`
	Statistics fields.Fields `json:"statistics,omitempty"`
	*internal
}

func (q *IndexQuery) Name() string {
	return q.idx.Name()
}

func (q *IndexQuery) InvalidParam() bool {
	if q.Filters != nil && q.Filters.CheckNumberRange() {
		return true
	}
	if q.Not != nil && q.Not.CheckNumberRange() {
		return true
	}
	if len(q.Statistics.Fields()) == 0 && q.Page.From() < 0 {
		return true
	}

	if len(q.defaultQueryNotReturn) > 0 {
		if q.ReturnFiled == nil {
			q.ReturnFiled = &ReturnFiled{Excludes: q.defaultQueryNotReturn}
		} else {
			q.ReturnFiled.Excludes = append(q.ReturnFiled.Excludes, q.defaultQueryNotReturn...)
		}
	}
	return false
}

func (q *IndexQuery) Build() []byte {
	_query := map[string]any{}
	if queryFilter := q.queryBuilder(q.Filters, q.Not); queryFilter != nil {
		_query["query"] = queryFilter
	}
	aggs, lastField := q.aggBuilder(q.Group, q.Statistics, q.internal.extend)
	q.lastField = lastField
	if q.stat = len(aggs) > 0; q.stat {
		_query["aggs"] = aggs
		_query["size"] = 0
	} else {
		if q.ReturnFiled != nil {
			_query["_source"] = q.ReturnFiled.Source(q.index())
		}
		if q.Sort != nil && len(q.Sort.Sorts) > 0 {
			_query["sort"] = q.Sort.Sort(q.index())
		}
		if from := q.Page.From(); from > 0 {
			_query["from"] = from
		}
		if size := q.Page.Size(); size >= 0 {
			_query["size"] = size
		}
	}
	return tools.EncodeBytes(_query)
}

func (q *IndexQuery) ParseQueryResult(result *entity.EsQueryResult) (queryResult search.QueryResult, count int64, imprecise byte) {
	if q.stat {
		queryResult, imprecise = q.index().TransAggs(&result.Aggs, q.lastField)
	} else {
		for _, hit := range *result.Hits.Hits {
			queryResult = append(queryResult, hit.Source)
		}
		if count = result.Hits.Total.Value; count > maxCount {
			count = maxCount
		}
	}
	return
}

type internal struct {
	idx                   index.Index
	extend                fields.Extend
	stat                  bool
	lastField             string
	queryBuilder          filter.QueryBuilder
	aggBuilder            filter.AggBuilder
	defaultQueryNotReturn []string
}

func (q *internal) index() index.Index {
	return q.idx
}
