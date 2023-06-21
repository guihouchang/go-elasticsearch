package field

import (
	"entgo.io/ent/schema/field"
	"reflect"
	"strings"
)

type Type uint8

// List of field types.
const (
	TypeInvalid Type = iota
	TypeBool
	TypeTime
	TypeJSON
	TypeUUID
	TypeBytes
	TypeEnum
	TypeString
	TypeOther
	TypeInt8
	TypeInt16
	TypeInt32
	TypeInt
	TypeInt64
	TypeUint8
	TypeUint16
	TypeUint32
	TypeUint
	TypeUint64
	TypeFloat32
	TypeFloat64
	TypeNothing
	TypeKeyword
	TypeHalfFloat
	TypeScaledFloat
)

var (
	typeNames = [...]string{
		TypeInvalid:     "invalid",
		TypeBool:        "boolean",
		TypeTime:        "date",
		TypeJSON:        "keyword",
		TypeUUID:        "[16]byte",
		TypeBytes:       "binary",
		TypeEnum:        "string",
		TypeString:      "text",
		TypeOther:       "other",
		TypeInt:         "integer",
		TypeInt8:        "byte",
		TypeInt16:       "short",
		TypeInt32:       "integer",
		TypeInt64:       "long",
		TypeUint:        "integer",
		TypeUint8:       "short",
		TypeUint16:      "integer",
		TypeUint32:      "integer",
		TypeUint64:      "long",
		TypeFloat32:     "float",
		TypeFloat64:     "double",
		TypeNothing:     "nothing",
		TypeKeyword:     "keyword",
		TypeHalfFloat:   "half_float",
		TypeScaledFloat: "scaled_float",
	}
)

func (t Type) String() string {
	return typeNames[t]
}

// pkgName returns the package name from a Go
// identifier with a package qualifier.
func pkgName(ident string) string {
	i := strings.LastIndexByte(ident, '.')
	if i == -1 {
		return ""
	}
	s := ident[:i]
	if i := strings.LastIndexAny(s, "]*"); i != -1 {
		s = s[i+1:]
	}
	return s
}

func methods(t reflect.Type, rtype *field.RType) {
	// For type T, add methods with
	// pointer receiver as well (*T).
	if t.Kind() != reflect.Ptr {
		t = reflect.PtrTo(t)
	}
	n := t.NumMethod()
	for i := 0; i < n; i++ {
		m := t.Method(i)
		in := make([]*field.RType, m.Type.NumIn()-1)
		for j := range in {
			arg := m.Type.In(j + 1)
			in[j] = &field.RType{Name: arg.Name(), Ident: arg.String(), Kind: arg.Kind(), PkgPath: arg.PkgPath()}
		}
		out := make([]*field.RType, m.Type.NumOut())
		for j := range out {
			ret := m.Type.Out(j)
			out[j] = &field.RType{Name: ret.Name(), Ident: ret.String(), Kind: ret.Kind(), PkgPath: ret.PkgPath()}
		}
		rtype.Methods[m.Name] = struct{ In, Out []*field.RType }{in, out}
	}
}

// indirect returns the type at the end of indirection.
func indirect(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func pkgPath(t reflect.Type) string {
	pkg := t.PkgPath()
	if pkg != "" {
		return pkg
	}
	switch t.Kind() {
	case reflect.Slice, reflect.Array, reflect.Ptr, reflect.Map:
		return pkgPath(t.Elem())
	}
	return pkg
}
