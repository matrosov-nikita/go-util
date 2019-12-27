package cast

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_TryUInt8(t *testing.T) {
	// UInt8 error cases from string
	testTryUInt8Err(t, "-1", "256", "1024")

	// UInt8 no error cases from string
	testTryUInt8(t, "0", "1", "128", "255")

	// UInt8 error cases from int
	testTryUInt8Err(t, -1, 256, 1024)

	// UInt8 no error cases from int
	testTryUInt8(t, 0, 1, 128, 255)

	// UInt8 error cases from int64
	testTryUInt8Err(t, int64(-1), int64(256), int64(1024))

	// UInt8 no error cases from int64
	testTryUInt8(t, int64(0), int64(1), int64(128), int64(255))

	// UInt8 error cases from int32
	testTryUInt8Err(t, int32(-1), int32(256), int32(1024))

	// UInt8 no error cases from int32
	testTryUInt8(t, int32(0), int32(1), int32(128), int32(255))

	// UInt8 error cases from int16
	testTryUInt8Err(t, int16(-1), int16(256), int16(1024))

	// UInt8 no error cases from int16
	testTryUInt8(t, int16(0), int16(1), int16(128), int16(255))

	// UInt8 error cases from int8
	testTryUInt8Err(t, int8(-1))

	// UInt8 no error cases from int8
	testTryUInt8(t, int16(0), int16(1), int16(127))

	// UInt8 error cases from uint64
	testTryUInt8Err(t, uint64(256), uint64(1024))

	// UInt8 no error cases from uint64
	testTryUInt8(t, uint64(0), uint64(1), uint64(128), uint64(255))

	// UInt8 error cases from uint32
	testTryUInt8Err(t, uint32(256), uint32(1024))

	// UInt8 no error cases from uint32
	testTryUInt8(t, uint32(0), uint32(1), uint32(128), uint32(255))

	// UInt8 error cases from uint16
	testTryUInt8Err(t, uint16(256), uint16(1024))

	// UInt8 no error cases from uint16
	testTryUInt8(t, uint16(0), uint16(1), uint16(128), uint16(255))

	// UInt8 error cases from float64
	testTryUInt8Err(t, float64(-1), float64(256), float64(1024))

	// UInt8 no error cases from float64
	testTryUInt8(t, float64(0), float64(1), float64(128), float64(255))

	// UInt8 error cases from float32
	testTryUInt8Err(t, float32(-1), float32(256), float32(1024))

	// UInt8 no error cases from float32
	testTryUInt8(t, float32(0), float32(1), float32(128), float32(255))

	fromF64, err := TryUInt8(1.5)
	require.NoError(t, err)
	require.Equal(t, uint8(1), fromF64)

	fromF32, err := TryUInt8(float32(1.5))
	require.NoError(t, err)
	require.Equal(t, uint8(1), fromF32)

	_, err = TryUInt8(struct{}{})
	require.Error(t, err)
}

func Test_TryUInt16(t *testing.T) {
	// UInt16 error cases from string
	testTryUInt16Err(t, "-1", "65536", "123456")

	// UInt16 no error cases from string
	testTryUInt16(t, "0", "1", "1024", "65535")

	// UInt16 error cases from int
	testTryUInt16Err(t, -1, 65536, 123456)

	// UInt16 no error cases from int
	testTryUInt16(t, 0, 1, 1024, 65535)

	// UInt16 error cases from int64
	testTryUInt16Err(t, int64(-1), int64(65536), int64(123456))

	// UInt16 no error cases from int64
	testTryUInt16(t, int64(0), int64(1), int64(1024), int64(65535))

	// UInt16 error cases from int32
	testTryUInt16Err(t, int32(-1), int32(65536), int32(123456))

	// UInt16 no error cases from int32
	testTryUInt16(t, int32(0), int32(1), int32(1024), int32(65535))

	// UInt16 error cases from int16
	testTryUInt16Err(t, int16(-1))

	// UInt16 no error cases from int16
	testTryUInt16(t, int16(0), int16(1), int16(32767))

	// UInt16 error cases from int8
	testTryUInt16Err(t, int8(-1))

	// UInt16 no error cases from int8
	testTryUInt16(t, int8(0), int8(1), int8(127))

	// UInt16 error cases from uint64
	testTryUInt16Err(t, uint64(65536), uint64(262144))

	// UInt16 no error cases from uint64
	testTryUInt16(t, uint64(0), uint64(1), uint64(128), uint64(65535))

	// UInt16 error cases from uint32
	testTryUInt16Err(t, uint32(65536), uint32(262144))

	// UInt16 no error cases from uint32
	testTryUInt16(t, uint32(0), uint32(1), uint32(128), uint32(65535))

	// UInt16 error cases from float64
	testTryUInt16Err(t, float64(-1), float64(65536))

	// UInt16 no error cases from float64
	testTryUInt16(t, float64(0), float64(1), float64(128), float64(65535))

	// UInt16 error cases from float32
	testTryUInt16Err(t, float32(-1), float32(65536))

	// UInt16 no error cases from float32
	testTryUInt16(t, float32(0), float32(1), float32(128), float32(65535))

	fromF64, err := TryUInt16(1.5)
	require.NoError(t, err)
	require.Equal(t, uint16(1), fromF64)

	fromF32, err := TryUInt16(float32(1.5))
	require.NoError(t, err)
	require.Equal(t, uint16(1), fromF32)

	_, err = TryUInt16(struct{}{})
	require.Error(t, err)
}

func Test_TryUInt32(t *testing.T) {
	// UInt32 error cases from string
	testTryUInt32Err(t, "-1", "4294967296")

	// UInt32 no error cases from string
	testTryUInt32(t, "0", "1", "4294967295")

	// UInt32 error cases from int
	testTryUInt32Err(t, -1, 4294967296)

	// UInt32 no error cases from int
	testTryUInt32(t, 0, 1, 1024, 4294967295)

	// UInt32 error cases from int64
	testTryUInt32Err(t, int64(-1), int64(math.MaxInt64))

	// UInt32 no error cases from int64
	testTryUInt32(t, int64(0), int64(1), int64(1024), int64(math.MaxUint32))

	// UInt32 error cases from int32
	testTryUInt32Err(t, int32(-1))

	// UInt32 no error cases from int32
	testTryUInt32(t, int32(0), int32(1), int32(1024), int32(math.MaxInt32))

	// UInt32 error cases from int16
	testTryUInt32Err(t, int16(-1))

	// UInt32 no error cases from int16
	testTryUInt32(t, int16(0), int16(1), int16(math.MaxInt16))

	// UInt32 error cases from int8
	testTryUInt32Err(t, int8(-1))

	// UInt32 no error cases from int8
	testTryUInt32(t, int8(0), int8(1), int8(127))

	// UInt32 error cases from uint64
	testTryUInt32Err(t, uint64(math.MaxUint64))

	// UInt32 no error cases from uint64
	testTryUInt32(t, uint64(0), uint64(1), uint64(128), uint64(math.MaxUint32))

	// UInt32 error cases from float64
	testTryUInt32Err(t, float64(-1), float64(math.MaxInt64))

	// UInt32 no error cases from float64
	testTryUInt32(t, float64(0), float64(1), float64(128), float64(math.MaxUint32))

	// UInt32 error cases from float32
	testTryUInt32Err(t, float32(-1))

	// UInt32 no error cases from float32
	testTryUInt32(t, float32(0), float32(1), float32(128))

	fromF64, err := TryUInt32(1.5)
	require.NoError(t, err)
	require.Equal(t, uint32(1), fromF64)

	fromF32, err := TryUInt32(float32(1.5))
	require.NoError(t, err)
	require.Equal(t, uint32(1), fromF32)

	_, err = TryUInt32(struct{}{})
	require.Error(t, err)
}

func Test_TryUInt64(t *testing.T) {
	// UInt64 error cases from string
	testTryUInt64Err(t, "-1")

	// UInt64 no error cases from string
	testTryUInt64(t, "0", "1", "4294967295")

	// UInt64 error cases from int
	testTryUInt64Err(t, -1)

	// UInt64 no error cases from int
	testTryUInt64(t, 0, 1, 1024)

	// UInt64 error cases from int64
	testTryUInt64Err(t, int64(-1))

	// UInt64 no error cases from int64
	testTryUInt64(t, int64(0), int64(1), int64(1024), int64(math.MaxInt64))

	// UInt64 error cases from int32
	testTryUInt64Err(t, int32(-1))

	// UInt64 no error cases from int32
	testTryUInt64(t, int32(0), int32(1), int32(1024), int32(math.MaxInt32))

	// UInt64 error cases from int16
	testTryUInt64Err(t, int16(-1))

	// UInt64 no error cases from int16
	testTryUInt64(t, int16(0), int16(1), int16(math.MaxInt16))

	// UInt64 error cases from int8
	testTryUInt64Err(t, int8(-1))

	// UInt64 no error cases from int8
	testTryUInt64(t, int8(0), int8(1), int8(127))

	// UInt64 error cases from float64
	testTryUInt64Err(t, float64(-1))

	// UInt64 no error cases from float64
	testTryUInt64(t, float64(0), float64(1), float64(128), float64(math.MaxUint32))

	// UInt64 error cases from float32
	testTryUInt64Err(t, float32(-1))

	// UInt64 no error cases from float32
	testTryUInt64(t, float32(0), float32(1), float32(128))

	fromF64, err := TryUInt64(1.5)
	require.NoError(t, err)
	require.Equal(t, uint64(1), fromF64)

	fromF32, err := TryUInt64(float32(1.5))
	require.NoError(t, err)
	require.Equal(t, uint64(1), fromF32)

	_, err = TryUInt64(struct{}{})
	require.Error(t, err)
}

func Test_TryInt8(t *testing.T) {
	// Int8 error cases from string
	testTryInt8Err(t, "-1000", "256", "1024")

	// Int8 no error cases from string
	testTryInt8(t, "-128", "-1", "127")

	// Int8 error cases from int
	testTryInt8Err(t, -1000, 256, 1024)

	// Int8 no error cases from int
	testTryInt8(t, math.MinInt8, -1, math.MaxInt8)

	// Int8 error cases from int64
	testTryInt8Err(t, int64(math.MinInt64), int64(256), int64(math.MaxInt64))

	// Int8 no error cases from int64
	testTryInt8(t, int64(math.MinInt8), int64(-1), int64(127), int64(math.MaxInt8))

	// Int8 error cases from int32
	testTryInt8Err(t, int32(math.MinInt32), int32(256), int32(math.MaxInt32))

	// Int8 no error cases from int32
	testTryInt8(t, int32(math.MinInt8), int32(-1), int32(127), int32(math.MaxInt8))

	// Int8 error cases from int16
	testTryInt8Err(t, int16(math.MinInt16), int32(256), int16(math.MaxInt16))

	// Int8 no error cases from int16
	testTryInt8(t, int16(math.MinInt8), int16(-1), int16(127), int16(math.MaxInt8))

	// Int8 error cases from uint64
	testTryInt8Err(t, uint64(256), uint64(math.MaxUint64))

	// Int8 no error cases from uint64
	testTryInt8(t, uint64(0), uint64(math.MaxInt8))

	// Int8 error cases from uint32
	testTryInt8Err(t, uint32(256), uint32(math.MaxUint32))

	// Int8 no error cases from uint32
	testTryInt8(t, uint32(0), uint32(1), uint32(math.MaxInt8))

	// Int8 error cases from uint16
	testTryInt8Err(t, uint16(256), uint16(math.MaxUint16))

	// Int8 no error cases from uint16
	testTryInt8(t, uint16(0), uint16(1), uint16(math.MaxInt8))

	// Int8 error cases from uint8
	testTryInt8Err(t, uint16(math.MaxUint8))

	// Int8 no error cases from uint8
	testTryInt8(t, uint8(0), uint8(1), uint8(math.MaxInt8))

	// Int8 error cases from float64
	testTryInt8Err(t, float64(256), float64(1024))

	// Int8 no error cases from float64
	testTryInt8(t, float64(math.MinInt8), float64(0), float64(1), float64(math.MaxInt8))

	// Int8 error cases from float32
	testTryInt8Err(t, float32(256), float32(1024))

	// Int8 no error cases from float32
	testTryInt8(t, float64(math.MinInt8), float32(math.MaxInt8))

	fromF64, err := TryInt8(1.5)
	require.NoError(t, err)
	require.Equal(t, int8(1), fromF64)

	fromF32, err := TryInt8(float32(1.5))
	require.NoError(t, err)
	require.Equal(t, int8(1), fromF32)

	_, err = TryInt8(struct{}{})
	require.Error(t, err)
}

func Test_TryInt16(t *testing.T) {
	// Int16 error cases from string
	testTryInt16Err(t, "65536", "123456")

	// Int16 no error cases from string
	testTryInt16(t, "-1", "0", "1", "1024", "32767")

	// Int16 error cases from int
	testTryInt16Err(t, math.MinInt32, math.MaxInt32)

	// Int16 no error cases from int
	testTryInt16(t, math.MinInt16, 0, 1, 1024, math.MaxInt16)

	// Int16 error cases from int64
	testTryInt16Err(t, int64(math.MinInt64), int64(math.MaxInt64))

	// Int16 no error cases from int64
	testTryInt16(t, int64(math.MinInt16), int64(1), int64(1024), int64(math.MaxInt16))

	// Int16 error cases from int32
	testTryInt16Err(t, int32(math.MinInt32), int32(math.MaxInt32))

	// Int16 no error cases from int32
	testTryInt16(t, int32(math.MinInt16), int32(1), int32(1024), int32(math.MaxInt16))

	// Int16 error cases from uint64
	testTryInt16Err(t, uint64(math.MaxUint16), uint64(math.MaxInt32))

	// Int16 no error cases from uint64
	testTryInt16(t, uint64(0), uint64(1), uint64(math.MaxInt16))

	// Int16 error cases from uint32
	testTryInt16Err(t, uint32(math.MaxUint16), uint32(math.MaxInt32))

	// Int16 no error cases from uint32
	testTryInt16(t, uint32(0), uint32(1), uint32(math.MaxInt16))

	// Int16 error cases from uint16
	testTryInt16Err(t, uint16(math.MaxUint16))

	// Int16 no error cases from uint16
	testTryInt16(t, uint32(0), uint32(1), uint16(math.MaxInt16))

	// Int16 error cases from float64
	testTryInt16Err(t, float64(math.MinInt64), float64(math.MaxInt64))

	// Int16 no error cases from float64
	testTryInt16(t, float64(math.MinInt16), float64(math.MaxInt16))

	// Int16 error cases from float32
	testTryInt16Err(t, float32(math.MinInt32), float32(math.MaxInt32))

	// Int16 no error cases from float32
	testTryInt16(t, float32(math.MinInt16), float32(math.MaxInt16))

	fromF64, err := TryInt16(1.5)
	require.NoError(t, err)
	require.Equal(t, int16(1), fromF64)

	fromF32, err := TryInt16(float32(1.5))
	require.NoError(t, err)
	require.Equal(t, int16(1), fromF32)

	_, err = TryInt16(struct{}{})
	require.Error(t, err)
}

func Test_TryInt32(t *testing.T) {
	// Int32 error cases from string
	testTryInt32Err(t, "4294967296")

	// Int32 no error cases from string
	testTryInt32(t, "-2147483648", "0", "1", "2147483647")

	// Int32 error cases from int
	testTryInt32Err(t, -2147483649, 4294967296)

	// Int32 no error cases from int
	testTryInt32(t, -2147483648, 0, 1, 1024, 2147483647)

	// Int32 error cases from int64
	testTryInt32Err(t, int64(math.MinInt64), int64(math.MaxInt64))

	// Int32 no error cases from int64
	testTryInt32(t, int64(math.MinInt32), int64(0), int64(1), int64(1024), int64(math.MaxInt32))

	// Int32 error cases from uint64
	testTryInt32Err(t, uint64(math.MaxUint64))

	// Int32 no error cases from uint64
	testTryInt32(t, uint64(0), uint64(1), uint64(128), uint64(math.MaxInt32))

	// Int32 error cases from uint32
	testTryInt32Err(t, uint32(math.MaxUint32))

	// Int32 no error cases from uint32
	testTryInt32(t, uint32(0), uint32(1), uint32(128), uint32(math.MaxInt32))

	// Int32 error cases from float64
	testTryInt32Err(t, float64(math.MinInt64), float64(math.MaxInt64))

	// Int32 no error cases from float64
	testTryInt32(t, float64(math.MinInt32), float64(math.MaxInt32))

	// Int32 error cases from float32
	testTryInt32Err(t, float32(math.MaxUint32))

	// Int32 no error cases from float32
	testTryInt32(t, float32(0), float32(1), float32(128))

	fromF64, err := TryInt32(1.5)
	require.NoError(t, err)
	require.Equal(t, int32(1), fromF64)

	fromF32, err := TryInt32(float32(1.5))
	require.NoError(t, err)
	require.Equal(t, int32(1), fromF32)

	_, err = TryInt32(struct{}{})
	require.Error(t, err)
}

func Test_TryInt64(t *testing.T) {
	// Int64 error cases from string
	testTryInt64Err(t, "18446744073709551616")

	// Int64 no error cases from string
	testTryInt64(t, "0", "1", "4294967295")

	// Int64 error cases from uint64
	testTryInt64Err(t, uint64(math.MaxUint64))

	// Int64 no error cases from uint64
	testTryInt64(t, uint64(0), uint64(1), uint64(1024), uint64(math.MaxInt64))

	fromF64, err := TryInt64(1.5)
	require.NoError(t, err)
	require.Equal(t, int64(1), fromF64)

	fromF32, err := TryInt64(float32(1.5))
	require.NoError(t, err)
	require.Equal(t, int64(1), fromF32)

	_, err = TryInt64(struct{}{})
	require.Error(t, err)
}

func Test_TryFloat64(t *testing.T) {
	// Int64 error cases from string
	testTryInt64Err(t, "18446744073709551616")

	// Int64 no error cases from string
	testTryInt64(t, "0", "1", "4294967295")

	// Int64 error cases from uint64
	testTryInt64Err(t, uint64(math.MaxUint64))

	// Int64 no error cases from uint64
	testTryInt64(t, uint64(0), uint64(1), uint64(1024), uint64(math.MaxInt64))

	fromF64, err := TryInt64(1.5)
	require.NoError(t, err)
	require.Equal(t, int64(1), fromF64)

	fromF32, err := TryInt64(float32(1.5))
	require.NoError(t, err)
	require.Equal(t, int64(1), fromF32)

	_, err = TryInt64(struct{}{})
	require.Error(t, err)
}

func Test_TryString(t *testing.T) {
	expected, err := TryString("expected")
	require.NoError(t, err)
	require.Equal(t, "expected", expected)

	expected, err = TryString([]byte("expected"))
	require.NoError(t, err)
	require.Equal(t, "expected", expected)

	expected, err = TryString(123)
	require.NoError(t, err)
	require.Equal(t, "123", expected)

	expected, err = TryString(123.5)
	require.NoError(t, err)
	require.Equal(t, "123.5", expected)

	_, err = TryString(struct{}{})
	require.Error(t, err)
}

func Test_TryUUID(t *testing.T) {
	_, err := TryUUID("not an uuid")
	require.Error(t, err)

	_, err = TryUUID([]byte("not an uuid"))
	require.Error(t, err)

	_, err = TryUUID(123)
	require.Error(t, err)

	encoded, err := TryUUID("523bbaf3-7ef1-4f4a-8713-ab8217a8f182")
	require.NoError(t, err)
	require.Equal(t, "523bbaf3-7ef1-4f4a-8713-ab8217a8f182", encoded)
}

func Test_TryDateTime(t *testing.T) {
	now := time.Now()
	parsed, err := TryDateTime(now)
	require.NoError(t, err)
	require.Equal(t, now, parsed)

	timeStr := now.Format(dateTimeFormat)
	parsed, err = TryDateTime(timeStr)
	require.NoError(t, err)
	require.Equal(t, timeStr, parsed.Format(dateTimeFormat))

	timeStr = now.Format(time.RFC3339Nano)
	parsed, err = TryDateTime(timeStr)
	require.Error(t, err)
}

func Test_TryDate(t *testing.T) {
	now := time.Now()
	parsed, err := TryDate(now)
	require.NoError(t, err)
	require.Equal(t, now, parsed)

	dateStr := now.Format(dateFormat)
	parsed, err = TryDate(dateStr)
	require.NoError(t, err)
	require.Equal(t, dateStr, parsed.Format(dateFormat))

	dateStr = now.Format(time.RFC3339Nano)
	parsed, err = TryDate(dateStr)
	require.Error(t, err)
}

func testTryUInt8Err(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryUInt8(inVal)
		require.Error(t, err)
	}
}

func testTryUInt8(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryUInt8(inVal)
		require.NoError(t, err)
	}
}

func testTryUInt16Err(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryUInt16(inVal)
		require.Error(t, err)
	}
}

func testTryUInt16(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryUInt16(inVal)
		require.NoError(t, err)
	}
}

func testTryUInt32Err(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryUInt32(inVal)
		require.Error(t, err)
	}
}

func testTryUInt32(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryUInt32(inVal)
		require.NoError(t, err)
	}
}

func testTryUInt64Err(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryUInt64(inVal)
		require.Error(t, err)
	}
}

func testTryUInt64(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryUInt64(inVal)
		require.NoError(t, err)
	}
}

func testTryInt8Err(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryInt8(inVal)
		require.Error(t, err)
	}
}

func testTryInt8(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryInt8(inVal)
		require.NoError(t, err)
	}
}

func testTryInt16Err(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryInt16(inVal)
		require.Error(t, err)
	}
}

func testTryInt16(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryInt16(inVal)
		require.NoError(t, err)
	}
}

func testTryInt32Err(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryInt32(inVal)
		require.Error(t, err)
	}
}

func testTryInt32(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryInt32(inVal)
		require.NoError(t, err)
	}
}

func testTryInt64Err(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryInt64(inVal)
		require.Error(t, err)
	}
}

func testTryInt64(t *testing.T, in ...interface{}) {
	for _, inVal := range in {
		_, err := TryInt64(inVal)
		require.NoError(t, err)
	}
}
