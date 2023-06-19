package field

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
		TypeJSON:        "json.RawMessage",
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
