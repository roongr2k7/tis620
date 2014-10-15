package tis620

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	expectUTF8 := []byte("hello สวัสดี")
	b := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0xca, 0xc7, 0xd1, 0xca, 0xb4, 0xd5}
	actual := Convert(b)
	if !reflect.DeepEqual(expectUTF8, actual) {
		t.Errorf("expect %v but %v", expectUTF8, actual)
		t.Errorf("expect %q but %q", expectUTF8, actual)
	}
}
