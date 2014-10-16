package tis620

type TIS620 struct {
	buf []byte
}

func NewTIS620(b []byte) *TIS620 {
	return &TIS620{buf: b}
}

func (t *TIS620) UTF8() []byte {
	ret := []byte{}
	for _, v := range t.buf {
		if v > 0xa0 {
			d := v - 0xa0
			ret = append(ret, 0xe0, 0xb8|(d>>6), 0x80|(d&0x3f))
		} else {
			ret = append(ret, v)
		}
	}
	return ret
}
