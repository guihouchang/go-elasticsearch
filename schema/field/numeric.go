package field

import "entgo.io/ent/schema/field"

type int8Builder struct {
	desc *field.Descriptor
}

func (t *int8Builder) Descriptor() *field.Descriptor {
	return t.desc
}

func (t *int8Builder) Comment(c string) *int8Builder {
	t.desc.Comment = c
	return t
}

type int16Builder struct {
	desc *field.Descriptor
}

func (t *int16Builder) Descriptor() *field.Descriptor {
	return t.desc
}

func (t *int16Builder) Comment(c string) *int16Builder {
	t.desc.Comment = c
	return t
}

type int32Builder struct {
	desc *field.Descriptor
}

func (t *int32Builder) Descriptor() *field.Descriptor {
	return t.desc
}

func (t *int32Builder) Comment(c string) *int32Builder {
	t.desc.Comment = c
	return t
}

type int64Builder struct {
	desc *field.Descriptor
}

func (t *int64Builder) Descriptor() *field.Descriptor {
	return t.desc
}

func (t *int64Builder) Comment(c string) *int64Builder {
	t.desc.Comment = c
	return t
}

type floatBuilder struct {
	desc *field.Descriptor
}

func (t *floatBuilder) Descriptor() *field.Descriptor {
	return t.desc
}

func (t *floatBuilder) Comment(c string) *floatBuilder {
	t.desc.Comment = c
	return t
}

type doubleBuilder struct {
	desc *field.Descriptor
}

func (t *doubleBuilder) Descriptor() *field.Descriptor {
	return t.desc
}

func (t *doubleBuilder) Comment(c string) *doubleBuilder {
	t.desc.Comment = c
	return t
}

type halfFloatBuilder struct {
	desc *field.Descriptor
}

func (t *halfFloatBuilder) Descriptor() *field.Descriptor {
	return t.desc
}

func (t *halfFloatBuilder) Comment(c string) *halfFloatBuilder {
	t.desc.Comment = c
	return t
}

type scaledFloatBuilder struct {
	desc *field.Descriptor
}

func (t *scaledFloatBuilder) Descriptor() *field.Descriptor {
	return t.desc
}

func (t *scaledFloatBuilder) Comment(c string) *scaledFloatBuilder {
	t.desc.Comment = c
	return t
}

func Byte(name string) *int8Builder {
	return &int8Builder{desc: &field.Descriptor{
		Info: &field.TypeInfo{Type: field.Type(TypeInt8)},
		Name: name,
	}}
}

func Short(name string) *int16Builder {
	return &int16Builder{desc: &field.Descriptor{
		Info: &field.TypeInfo{Type: field.Type(TypeInt16)},
		Name: name,
	}}
}

func Int(name string) *int32Builder {
	return &int32Builder{desc: &field.Descriptor{
		Info: &field.TypeInfo{
			Type: field.Type(TypeInt32),
		},
		Name: name,
	}}
}

func Long(name string) *int64Builder {
	return &int64Builder{desc: &field.Descriptor{
		Info: &field.TypeInfo{Type: field.Type(TypeInt64)},
		Name: name,
	}}
}

func Float(name string) *floatBuilder {
	return &floatBuilder{desc: &field.Descriptor{
		Info: &field.TypeInfo{Type: field.Type(TypeFloat32)},
		Name: name,
	}}
}

func Double(name string) *doubleBuilder {
	return &doubleBuilder{desc: &field.Descriptor{
		Info: &field.TypeInfo{Type: field.Type(TypeFloat64)},
		Name: name,
	}}
}

// HalfFloat 缩放类型的的浮点数, 比如price字段只需精确到分, 57.34缩放因子为100, 存储结果为5734
func HalfFloat(name string) *halfFloatBuilder {
	return &halfFloatBuilder{desc: &field.Descriptor{
		Info: &field.TypeInfo{Type: field.Type(TypeHalfFloat)},
		Name: name,
	}}
}

func ScaledFloat(name string) *scaledFloatBuilder {
	return &scaledFloatBuilder{desc: &field.Descriptor{
		Info: &field.TypeInfo{Type: field.Type(TypeScaledFloat)},
		Name: name,
	}}
}
