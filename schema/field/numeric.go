package field

import "entgo.io/ent/schema/field"

type int8Builder struct {
	desc *field.Descriptor
}

func (t *int8Builder) Descriptor() *field.Descriptor {
	return t.desc
}

type int16Builder struct {
	desc *field.Descriptor
}

func (t *int16Builder) Descriptor() *field.Descriptor {
	return t.desc
}

type int32Builder struct {
	desc *field.Descriptor
}

func (t *int32Builder) Descriptor() *field.Descriptor {
	return t.desc
}

type int64Builder struct {
	desc *field.Descriptor
}

func (t *int64Builder) Descriptor() *field.Descriptor {
	return t.desc
}

type floatBuilder struct {
	desc *field.Descriptor
}

func (t *floatBuilder) Descriptor() *field.Descriptor {
	return t.desc
}

type doubleBuilder struct {
	desc *field.Descriptor
}

func (t *doubleBuilder) Descriptor() *field.Descriptor {
	return t.desc
}

type halfFloatBuilder struct {
	desc *field.Descriptor
}

func (t *halfFloatBuilder) Descriptor() *field.Descriptor {
	return t.desc
}

type scaledFloatBuilder struct {
	desc *field.Descriptor
}

func (t *scaledFloatBuilder) Descriptor() *field.Descriptor {
	return t.desc
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
