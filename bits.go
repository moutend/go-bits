package bits

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Buffer struct {
	bits string
}

func New() *Buffer {
	return &Buffer{}
}

func (b *Buffer) Scan(src io.Reader, n int) (uint64, error) {
	if n > 64 {
		return 0, fmt.Errorf("bits: must be n <= 64")
	}
	if len(b.bits) < n {
		dst := &bytes.Buffer{}
		requiredBytes := int64(1 + (n-len(b.bits))/8)

		readBytes, err := io.CopyN(dst, src, requiredBytes)

		if err != nil {
			return 0, fmt.Errorf("bits: failed to read: %w", err)
		}
		if readBytes != requiredBytes {
			return 0, fmt.Errorf("bits: failed to read %d bytes", requiredBytes)
		}
		for _, v := range dst.Bytes() {
			b.bits += fmt.Sprintf("%08b", v)
		}
	}

	i64, err := strconv.ParseInt(b.bits[:n], 2, 64)

	if err != nil {
		return 0, fmt.Errorf("bits: %w", err)
	}

	b.bits = b.bits[n:]

	return uint64(i64), nil
}

func (b *Buffer) Append(value uint64, n int) error {
	if n > 64 {
		return fmt.Errorf("bits: must be n <= 64")
	}

	b.bits += fmt.Sprintf("%064b", value)[64-n:]

	return nil
}

func (b *Buffer) Bytes() []byte {
	bits := b.bits

	if remainder := len(bits) % 8; remainder > 0 {
		for i := 0; i < 8-remainder; i++ {
			bits += "0"
		}
	}

	length := len(bits)
	data := make([]byte, 0, length/8)

	for i := 0; i < length; i += 8 {
		i64, err := strconv.ParseInt(bits[i:i+8], 2, 64)

		if err != nil {
			return nil
		}

		data = append(data, byte(i64))
	}

	return data
}
