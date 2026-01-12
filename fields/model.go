package fields

type Extend interface {
	IsNestedField(Field) (parent string, nested bool)
	Group(f Field, next map[string]any) map[string]any
	Statistics(Field) map[string]any
}

type Fields interface {
	Appends(...Field) Fields
	Fields() []Field
	Len() int
	Set(int, Field) Fields
	Get(int) Field
}

type Field interface {
	Name() string
	DateInterval() string
	DataInterval() int
	Missing() any
	MinDocCount() int
	FixedDateInterval() bool
	ScriptId() string
}

type FieldEditor interface {
	Field
	WithDateInterval(dateInterval string) FieldEditor
	WithDataInterval(dataInterval int) FieldEditor
	WithDefault(missing any) FieldEditor
	WithMinDocCount(min int) FieldEditor
	WithScript(script string) FieldEditor
	AddParam(key string, value any) FieldEditor
}
