package field

import (
	"entgo.io/ent/schema/field"
)

type textBuilder struct {
	desc *field.Descriptor
}

// Analyzer 分词器
func (t *textBuilder) Analyzer(analyzer string) *textBuilder {
	if t.desc.SchemaType == nil {
		t.desc.SchemaType = map[string]string{}
	}
	t.desc.SchemaType["analyzer"] = analyzer
	return t
}

// SearchAnalyzer 搜索分词器
// Deprecated: 该方法已经弃用,在es7.x版本中不在设置search_analyzer
func (t *textBuilder) SearchAnalyzer(sa string) *textBuilder {
	if t.desc.SchemaType == nil {
		t.desc.SchemaType = map[string]string{}
	}
	//t.desc.SchemaType["search_analyzer"] = sa
	return t
}

func (t *textBuilder) Descriptor() *field.Descriptor {
	return t.desc
}

type keywordBuilder struct {
	desc *field.Descriptor
}

type dateBuilder struct {
	desc *field.Descriptor
}

func (d *dateBuilder) Descriptor() *field.Descriptor {
	return d.desc
}

func (d *dateBuilder) Format(format string) *dateBuilder {
	if d.desc.SchemaType == nil {
		d.desc.SchemaType = map[string]string{}
	}
	d.desc.SchemaType["format"] = format
	return d
}

type boolBuilder struct {
	desc *field.Descriptor
}

func (b *boolBuilder) Descriptor() *field.Descriptor {
	return b.desc
}

func Text(name string) *textBuilder {
	builder := &textBuilder{desc: &field.Descriptor{
		Info: &field.TypeInfo{Type: field.TypeString},
		Name: name,
	}}
	return builder
}

func (k *keywordBuilder) Descriptor() *field.Descriptor {
	return k.desc
}

func Keyword(name string) *keywordBuilder {
	return &keywordBuilder{desc: &field.Descriptor{
		Info: &field.TypeInfo{Type: field.Type(TypeKeyword)},
		Name: name,
	}}
}

func Date(name string) *dateBuilder {
	return &dateBuilder{desc: &field.Descriptor{
		Info: &field.TypeInfo{Type: field.TypeTime},
		Name: name,
	}}
}

func Bool(name string) *boolBuilder {
	return &boolBuilder{desc: &field.Descriptor{
		Info: &field.TypeInfo{Type: field.TypeBool},
		Name: name,
	}}
}

type jsonBuilder struct {
	desc *field.Descriptor
}

func (j *jsonBuilder) Descriptor() *field.Descriptor {
	return j.desc
}

func Strings(name string) *jsonBuilder {
	return &jsonBuilder{desc: field.JSON(name, []string{}).Descriptor()}
}

func Ints(name string) *jsonBuilder {
	return &jsonBuilder{desc: field.JSON(name, []int{}).Descriptor()}
}

func Int64s(name string) *jsonBuilder {
	return &jsonBuilder{desc: field.JSON(name, []int64{}).Descriptor()}
}

func Floats(name string) *jsonBuilder {
	return &jsonBuilder{desc: field.JSON(name, []float32{}).Descriptor()}
}

func JSON(name string, typ any) *jsonBuilder {
	return &jsonBuilder{desc: field.JSON(name, typ).Descriptor()}
}
