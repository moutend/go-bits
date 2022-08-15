package bits

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScan(t *testing.T) {
	data := []byte{0x78, 0x00, 0x03, 0xe8, 0x00, 0x00, 0x13, 0x88, 0x00}
	src := bytes.NewBuffer(data)
	buffer := New()

	bitsPerField, err := buffer.Scan(src, 5)

	require.NoError(t, err)
	require.Equal(t, uint64(15), bitsPerField)

	minX, err := buffer.Scan(src, int(bitsPerField))

	require.NoError(t, err)
	require.Equal(t, uint64(0), minX)

	maxX, err := buffer.Scan(src, int(bitsPerField))

	require.NoError(t, err)
	require.Equal(t, uint64(8000), maxX)

	minY, err := buffer.Scan(src, int(bitsPerField))

	require.NoError(t, err)
	require.Equal(t, uint64(0), minY)

	maxY, err := buffer.Scan(src, int(bitsPerField))

	require.NoError(t, err)
	require.Equal(t, uint64(10000), maxY)
}

func TestAppend(t *testing.T) {
	buffer := New()

	require.NoError(t, buffer.Append(uint64(15), 5))
	require.NoError(t, buffer.Append(uint64(0), 15))
	require.NoError(t, buffer.Append(uint64(8000), 15))
	require.NoError(t, buffer.Append(uint64(0), 15))
	require.NoError(t, buffer.Append(uint64(10000), 15))

	expected := []byte{0x78, 0x00, 0x03, 0xe8, 0x00, 0x00, 0x13, 0x88, 0x00}
	actual := buffer.Bytes()

	require.Equal(t, expected, actual)
}
