package altMap

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	kvm := KVMap{
		"index":      23,
		"percentage": 23.45,
		"b":          'a',
		"value":      "Strings",
	}

	if v := kvm.Get("index", -1, reflect.Int); v != 23 {
		t.Errorf("for `index` expected 23, got %v", v)
	}

	if vs := kvm.Get("value", 23.5, reflect.Int); vs != 23.5 {
		t.Errorf("for `value` expected 23.5, got %v", vs)
	}

	if vp := kvm.Get("value", 23.5, reflect.String); vp != "Strings" {
		t.Errorf("for `value` expected `Strings`, got %v", vp)
	}

}
