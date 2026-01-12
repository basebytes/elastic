package query

import "github.com/basebytes/elastic/index"

type ReturnFiled struct {
	Excludes []string `json:"excludes,omitempty"`
	Includes []string `json:"includes,omitempty"`
}

func (f *ReturnFiled) Source(index index.Index) map[string]any {
	for i, field := range f.Excludes {
		f.Excludes[i] = index.QueryField(field)
	}
	for i, field := range f.Includes {
		f.Includes[i] = index.QueryField(field)
	}
	_source := make(map[string]any)
	if len(f.Excludes) > 0 {
		_source["excludes"] = f.Excludes
	}
	if len(f.Includes) > 0 {
		_source["includes"] = f.Includes
	}
	return _source
}
