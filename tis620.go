package tis620

func Convert(b []byte) []byte {
	ret := []byte{}
	for _, v := range b {
		if v > 0xa0 {
			d := v - 0xa0
			ret = append(ret, 0xe0, 0xb8|(d>>6), 0x80|(d&0x3f))
		} else {
			ret = append(ret, v)
		}
	}
	return ret
}
