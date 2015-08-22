package altMap

import "reflect"

type KVMap map[string]interface{}

func (kvm KVMap) Get(key string, alt interface{}, wanted reflect.Kind) interface{} {
	retr, ok := kvm[key]
	if !ok {
		return alt
	}

	var coercedKind reflect.Kind

	switch retr.(type) {
	case string:
		coercedKind = reflect.String
	case bool:
		coercedKind = reflect.Bool
	case int32:
		coercedKind = reflect.Int32
	case int:
		coercedKind = reflect.Int
	case uint:
		coercedKind = reflect.Uint
	case uint16:
		coercedKind = reflect.Uint16
	case uint32:
		coercedKind = reflect.Uint32
	case uint64:
		coercedKind = reflect.Uint64
	case int8:
		coercedKind = reflect.Int8
	case uint8:
		coercedKind = reflect.Uint8
	case float32:
		coercedKind = reflect.Float32
	case float64:
		coercedKind = reflect.Float64
	case uintptr:
		coercedKind = reflect.Uintptr
	}

	if coercedKind == wanted {
		return retr
	}

	return alt
}
