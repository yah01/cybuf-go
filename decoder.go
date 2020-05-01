package cybuf

import (
	"io"
)

const (
	bufSize = 1 << 8
)

type Decoder struct {
	reader     io.Reader
	buf        []byte
	bufPointer int
	bufLength  int
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		reader: r,
		buf:    make([]byte, bufSize),
	}
}

//func (dec *Decoder) Decode(v interface{}) error {
//	var (
//		c              byte
//		quote          byte
//		leftBraceCount int
//		cybuf          []byte
//		err            error
//	)
//
//	// Skip space character
//	for c, err = dec.read(); err == nil && c != 0; c, err = dec.read() {
//		if unicode.IsSpace(rune(c)) {
//			continue
//		}
//	}
//	if err != nil {
//		return err
//	}
//
//	if c, err = dec.read(); err != nil {
//		return err
//	}
//	if c != '{' {
//		return DecodeError_Not_Found_Begin
//	}
//	leftBraceCount = 1
//	cybuf = append(cybuf, c)
//
//	for c, err = dec.read(); err == nil && c != 0; c, err = dec.read() {
//		if quote == 0 {
//			switch c {
//			case '"', '\'':
//				quote = c
//			case '{':
//				leftBraceCount++
//			case '}':
//				leftBraceCount--
//			}
//		} else {
//			switch c {
//			case quote:
//				quote = 0
//			}
//		}
//
//		cybuf = append(cybuf, c)
//		if leftBraceCount == 0 {
//			break
//		}
//	}
//	if err != nil {
//		return err
//	}
//
//	if cybuf[len(cybuf)-1] != '}' {
//		return DecodeError_Not_Found_End
//	}
//
//	if err = Unmarshal(cybuf, v); err != nil {
//		return err
//	}
//
//	return nil
//}

func (dec *Decoder) read() (byte, error) {
	end, err := dec.readBlock()
	if err != nil {
		return 0, err
	}
	if end {
		return 0, nil
	}
	firstByte := dec.buf[dec.bufPointer]
	dec.bufPointer++
	return firstByte, nil
}

func (dec *Decoder) readBlock() (bool, error) {
	if dec.bufPointer >= dec.bufLength {
		n, err := dec.reader.Read(dec.buf)
		if err != nil {
			return false, err
		}
		dec.bufPointer = 0
		dec.bufLength = n
	}

	if dec.bufLength == 0 {
		return true, nil
	}
	return false, nil
}
