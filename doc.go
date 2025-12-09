package elastic

import "github.com/basebytes/elastic-go/client/api"

type IndexDoc interface {
	api.Doc
	RetryOnConflict(retry int)
	Retied() int
}

func NewUpdateOnlyDocWrapper(doc any) *WrappedDoc {
	return &WrappedDoc{docWrapper: &docWrapper{Doc: doc}}
}

func NewUpsertDocWrapper(doc any) *WrappedDoc {
	return &WrappedDoc{docWrapper: &docWrapper{Doc: doc, DocAsUpsert: true}}
}

func NewMustRunScriptWrapper(id string, params any) *WrappedDoc {
	return &WrappedDoc{scriptWrapper: &scriptWrapper{Script: &script{ID: id, Params: params}, ScriptedUpsert: true, Upsert: struct{}{}}}
}

func NewUpsertScriptWrapper(id string, params any, inserted any, mustRunScript bool) *WrappedDoc {
	return &WrappedDoc{scriptWrapper: &scriptWrapper{Script: &script{ID: id, Params: params}, ScriptedUpsert: mustRunScript, Upsert: inserted}}
}

func NewUpdateExistsScriptWrapper(id string, params any) *WrappedDoc {
	return &WrappedDoc{scriptWrapper: &scriptWrapper{Script: &script{ID: id, Params: params}}}
}

type WrappedDoc struct {
	*docWrapper
	*scriptWrapper
}

type docWrapper struct {
	Doc         any  `json:"doc,omitempty"`
	DocAsUpsert bool `json:"doc_as_upsert,omitempty"`
}

type scriptWrapper struct {
	Script         *script `json:"script,omitempty"`
	ScriptedUpsert bool    `json:"scripted_upsert,omitempty"`
	Upsert         any     `json:"upsert,omitempty"`
}

type script struct {
	ID     string `json:"id"`
	Params any    `json:"params,omitempty"`
}
