package cybuf

import "io"

const (
	bufSize = 1 << 7
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

func (dec *Decoder) Decode(v interface{}) error {
	var err error
	if err = dec.read(); err != nil {
		return err
	}

	

	return nil
}

func (dec *Decoder) read() error {
	leftSize := dec.bufLength - dec.bufPointer
	if dec.bufLength-dec.bufPointer < 4 {
		for i := 0; i < leftSize; i++ {
			dec.buf[i] = dec.buf[dec.bufPointer+i]
		}
	}
	n, err := dec.reader.Read(dec.buf[leftSize:])
	if err != nil {
		return err
	}

	dec.bufLength = leftSize + n
}
