package fields

const (
	DateIntervalHour  = "hour"
	DateIntervalDay   = "day"
	DateIntervalMonth = "month"
	DateIntervalYear  = "year"
)

func NewFields() Fields {
	return &StandardFields{}
}

type StandardFields []*StandardField

func (f *StandardFields) Appends(fields ...Field) Fields {
	if f == nil {
		*f = []*StandardField{}
	}
	for _, field := range fields {
		if t, OK := field.(*StandardField); OK {
			*f = append(*f, t)
		}
	}
	return f
}

func (f *StandardFields) Fields() []Field {
	_fields := make([]Field, 0, len(*f))
	for _, g := range *f {
		_fields = append(_fields, g)
	}
	return _fields
}

func (f *StandardFields) Len() int {
	return len(*f)
}

func (f *StandardFields) Set(i int, field Field) Fields {
	if i >= 0 && i < f.Len() {
		if t, OK := field.(*StandardField); OK {
			(*f)[i] = t
		}
	}
	return f
}

func (f *StandardFields) Get(i int) Field {
	if i >= 0 && i < f.Len() {
		return (*f)[i]
	}
	return nil
}

func NewField() *StandardField {
	return &StandardField{}
}

type StandardField struct {
	FieldName         string         `json:"name"`
	DateFieldInterval string         `json:"dateInterval,omitempty"`
	DataFieldInterval int            `json:"interval,omitempty"`
	Default           any            `json:"missing,omitempty"`
	MinCount          int            `json:"minCount,omitempty"`
	Params            map[string]any `json:"params,omitempty"`
	Script            string         `json:"-"`
}

func (f *StandardField) Name() string {
	return f.FieldName
}

func (f *StandardField) DateInterval() string {
	return f.DateFieldInterval
}

func (f *StandardField) DataInterval() int {
	return f.DataFieldInterval
}

func (f *StandardField) Missing() any {
	return f.Default
}

func (f *StandardField) MinDocCount() int {
	return f.MinCount
}

func (f *StandardField) FixedDateInterval() bool {
	return f.Params != nil && f.Params["type"] == "fixed"
}

func (f *StandardField) ScriptId() string {
	return f.Script
}

func (f *StandardField) WithDateInterval(dateInterval string) FieldEditor {
	f.DateFieldInterval = dateInterval
	return f
}

func (f *StandardField) WithDataInterval(dataInterval int) FieldEditor {
	f.DataFieldInterval = dataInterval
	return f
}

func (f *StandardField) WithDefault(missing any) FieldEditor {
	f.Default = missing
	return f
}

func (f *StandardField) WithMinDocCount(min int) FieldEditor {
	f.MinCount = min
	return f
}

func (f *StandardField) WithScript(script string) FieldEditor {
	f.Script = script
	return f
}

func (f *StandardField) AddParam(key string, value any) FieldEditor {
	if f.Params == nil {
		f.Params = make(map[string]any)
	}
	f.Params[key] = value
	return f
}
