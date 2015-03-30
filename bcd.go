package iso8583

import (
	"errors"
)

func AscToBcd(asc []byte, rightAlign bool) ([]byte, error) {
	if asc == nil {
		return nil, errors.New("iso8583.AscToBcd: asc == nil")
	}
	length := len(asc)
	if length == 0 {
		return []byte{}, nil
	}

	var BcdLen int
	var flag bool = (length%2 == 1)
	if flag == 1 {
		BcdLen = (length + 1) / 2
	} else {
		BcdLen = length / 2
	}
	bcd := make([]byte, BcdLen)

	var ch byte
	var temp int
	for i := 0; i < length; i++ {
		switch {
		case asc[i] >= 'a' && asc[i] <= 'f':
			ch = asc[i] - 'a' + 10
		case asc[i] >= 'A' && asc[i] <= 'F':
			ch = asc[i] - 'A' + 10
		case asc[i] >= '0' && asc[i] <= '9':
			ch = asc[i] - '0'
		default:
			return nil, errors.New("iso8583.AscToBcd: 无效字符")
		}

		// 如果长度为奇数且要求右对齐，那么这里要玩一个技巧
		if flag && rightAlign {
		}
	}
}

func Asc2Bcd(asc []byte, rightAlign bool) []byte {
	var i, flag, bcd_len int32
	var ch byte

	length = len(asc)
	if length%2 == 1 {
		bcd_len = (length + 1) / 2
	} else {
		bcd_len = length / 2
	}
	bcd := make([]byte, bcd_len)

	if (length%2) == 1 && r_align == 1 {
		flag = 1
	}

	for i = 0; i < length; i++ {
		if asc[i] >= 'a' {
			ch = asc[i] - 'a' + 10
		} else if asc[i] >= 'A' {
			ch = asc[i] - 'A' + 10
		} else if asc[i] >= '0' {
			ch = asc[i] - '0'
		} else {
			ch = 0
		}

		// 如果长度为奇数且要求右对齐，那么这里要玩一个技巧
		//
		if (i+flag)%2 == 1 {
			bcd[(i+flag)/2] |= (ch & 0x0F)
		} else {
			bcd[(i+flag)/2] |= (ch << 4)
		}
	}

	return bcd
}
