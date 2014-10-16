package tis620

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	expectUTF8 := []byte("hello สวัสดี")
	b := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0xca, 0xc7, 0xd1, 0xca, 0xb4, 0xd5}
	tis620buf := NewTIS620(b)
	actual := tis620buf.UTF8()
	if !reflect.DeepEqual(expectUTF8, actual) {
		t.Errorf("expect %v but %v", expectUTF8, actual)
		t.Errorf("expect %q but %q", expectUTF8, actual)
	}
}

func TestIsTIS620(t *testing.T) {
	dt := []DataTableStringBool{
		DataTableStringBool{input: "TIS-620", expected: true},
		DataTableStringBool{input: "tis-620", expected: true},
		DataTableStringBool{input: "tis620", expected: false},
		DataTableStringBool{input: "utf-8", expected: false},
		DataTableStringBool{input: "", expected: false},
	}

	for _, test := range dt {
		actual := IsTIS620(test.input)
		if test.expected != actual {
			t.Errorf("expect %v but %v", test.expected, actual)
		}
	}
}

type DataTableStringBool struct {
	input    string
	expected bool
}
