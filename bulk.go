package elastic

import "github.com/basebytes/elastic-go/client/api"

const (
	ActionTypeCreate ActionType = "create"
	ActionTypeDelete ActionType = "delete"
	ActionTypeIndex  ActionType = "index"
	ActionTypeUpdate ActionType = "update"
)

type ActionType = string

type DocType = string

func NewMeta(docType DocType) *Meta {
	return &Meta{docType: docType, action: &api.BulkAction{}}
}

type Meta struct {
	docType DocType
	action  *api.BulkAction
	item    *api.BulkItem
	act     ActionType
}

func (m *Meta) SetAction(action ActionType) *Meta {
	if m.act == "" { //初始化为默认值create
		if action == m.act {
			action = ActionTypeCreate
		}
		m.action = &api.BulkAction{}
		m.item = &api.BulkItem{}
		item := m.getItem(action)
		*item = m.item
		m.act = action
	} else if item := m.getItem(action); item != nil {
		*item = m.item
		*m.getItem(m.act) = nil
		m.act = action
	}
	return m
}

func (m *Meta) Doctype() DocType {
	return m.docType
}

func (m *Meta) ActionType() ActionType {
	return m.act
}

func (m *Meta) GetId() string {
	return m.item.Id
}

func (m *Meta) SetId(id string) {
	m.item.Id = id
}

func (m *Meta) SetIndex(index string) *Meta {
	(*m.getItem(m.act)).Index = index
	return m
}

func (m *Meta) SetDoctype(docType DocType) *Meta {
	m.docType = docType
	return m
}

func (m *Meta) GetAction() *api.BulkAction {
	return m.action
}

func (m *Meta) RetryOnConflict(retry int) {
	(*m.getItem(m.act)).RetryOnConflict = retry
}

func (m *Meta) getItem(action ActionType) **api.BulkItem {
	var item **api.BulkItem
	switch action {
	case ActionTypeCreate:
		item = &m.action.Create
	case ActionTypeDelete:
		item = &m.action.Delete
	case ActionTypeIndex:
		item = &m.action.Index
	case ActionTypeUpdate:
		item = &m.action.Update
	}
	return item
}

func NewRetry() *Retry {
	return &Retry{}
}

type Retry struct {
	retry int
}

func (r *Retry) Retied() int {
	r.retry++
	return r.retry - 1
}
