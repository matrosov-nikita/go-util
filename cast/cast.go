package cast

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	dateTimeFormat = "2006-01-02T15:04:05"
	dateFormat     = "2006-01-02"
)

var errNumericOverFlow = errors.New("desired type overflow")

func TryUInt8(value interface{}) (uint8, error) {
	value = indirect(value)

	switch castedVal := value.(type) {
	case string:
		v, err := strconv.ParseUint(castedVal, 0, 8)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v to uint8: %s", value, err)
		}
		return uint8(v), nil
	case int:
		if castedVal < 0 || castedVal > math.MaxUint8 {
			return 0, errNumericOverFlow
		}
		return uint8(castedVal), nil
	case int64:
		if castedVal < 0 || castedVal > math.MaxUint8 {
			return 0, errNumericOverFlow
		}
		return uint8(castedVal), nil
	case int32:
		if castedVal < 0 || castedVal > math.MaxUint8 {
			return 0, errNumericOverFlow
		}
		return uint8(castedVal), nil
	case int16:
		if castedVal < 0 || castedVal > math.MaxUint8 {
			return 0, errNumericOverFlow
		}
		return uint8(castedVal), nil
	case int8:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint8(castedVal), nil
	case uint:
		if castedVal > math.MaxUint8 {
			return 0, errNumericOverFlow
		}
		return uint8(castedVal), nil
	case uint64:
		if castedVal > math.MaxUint8 {
			return 0, errNumericOverFlow
		}
		return uint8(castedVal), nil
	case uint32:
		if castedVal > math.MaxUint8 {
			return 0, errNumericOverFlow
		}
		return uint8(castedVal), nil
	case uint16:
		if castedVal > math.MaxUint8 {
			return 0, errNumericOverFlow
		}
		return uint8(castedVal), nil
	case uint8:
		return castedVal, nil
	case float64:
		if castedVal < 0 || castedVal > math.MaxUint8 {
			return 0, errNumericOverFlow
		}
		return uint8(castedVal), nil
	case float32:
		if castedVal < 0 || castedVal > math.MaxUint8 {
			return 0, errNumericOverFlow
		}
		return uint8(castedVal), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", value, value)
	}
}

func TryUInt16(value interface{}) (uint16, error) {
	value = indirect(value)

	switch castedVal := value.(type) {
	case string:
		v, err := strconv.ParseUint(castedVal, 0, 16)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v to uint16: %s", value, err)
		}
		return uint16(v), nil
	case int:
		if castedVal < 0 || castedVal > math.MaxUint16 {
			return 0, errNumericOverFlow
		}
		return uint16(castedVal), nil
	case int64:
		if castedVal < 0 || castedVal > math.MaxUint16 {
			return 0, errNumericOverFlow
		}
		return uint16(castedVal), nil
	case int32:
		if castedVal < 0 || castedVal > math.MaxUint16 {
			return 0, errNumericOverFlow
		}
		return uint16(castedVal), nil
	case int16:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint16(castedVal), nil
	case int8:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint16(castedVal), nil
	case uint:
		if castedVal > math.MaxUint16 {
			return 0, errNumericOverFlow
		}
		return uint16(castedVal), nil
	case uint64:
		if castedVal > math.MaxUint16 {
			return 0, errNumericOverFlow
		}
		return uint16(castedVal), nil
	case uint32:
		if castedVal > math.MaxUint16 {
			return 0, errNumericOverFlow
		}
		return uint16(castedVal), nil
	case uint16:
		return castedVal, nil
	case uint8:
		return uint16(castedVal), nil
	case float64:
		if castedVal < 0 || castedVal > math.MaxUint16 {
			return 0, errNumericOverFlow
		}
		return uint16(castedVal), nil
	case float32:
		if castedVal < 0 || castedVal > math.MaxUint16 {
			return 0, errNumericOverFlow
		}
		return uint16(castedVal), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", value, value)
	}
}

func TryUInt32(value interface{}) (uint32, error) {
	value = indirect(value)

	switch castedVal := value.(type) {
	case string:
		v, err := strconv.ParseUint(castedVal, 0, 32)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v to uint32: %s", value, err)
		}
		return uint32(v), nil
	case int:
		if castedVal < 0 || uint64(castedVal) > math.MaxUint32 {
			return 0, errNumericOverFlow
		}
		return uint32(castedVal), nil
	case int64:
		if castedVal < 0 || castedVal > math.MaxUint32 {
			return 0, errNumericOverFlow
		}
		return uint32(castedVal), nil
	case int32:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint32(castedVal), nil
	case int16:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint32(castedVal), nil
	case int8:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint32(castedVal), nil
	case uint:
		if castedVal > math.MaxUint32 {
			return 0, errNumericOverFlow
		}
		return uint32(castedVal), nil
	case uint64:
		if castedVal > math.MaxUint32 {
			return 0, errNumericOverFlow
		}
		return uint32(castedVal), nil
	case uint32:
		return castedVal, nil
	case uint16:
		return uint32(castedVal), nil
	case uint8:
		return uint32(castedVal), nil
	case float64:
		if castedVal < 0 || castedVal > math.MaxUint32 {
			return 0, errNumericOverFlow
		}
		return uint32(castedVal), nil
	case float32:
		if castedVal < 0 || castedVal > math.MaxUint32 {
			return 0, errNumericOverFlow
		}
		return uint32(castedVal), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", value, value)
	}
}

func TryUInt64(value interface{}) (uint64, error) {
	value = indirect(value)

	switch castedVal := value.(type) {
	case string:
		v, err := strconv.ParseUint(castedVal, 0, 64)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v to uint64: %s", value, err)
		}
		return v, nil
	case int:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint64(castedVal), nil
	case int64:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint64(castedVal), nil
	case int32:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint64(castedVal), nil
	case int16:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint64(castedVal), nil
	case int8:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint64(castedVal), nil
	case uint:
		return uint64(castedVal), nil
	case uint64:
		return castedVal, nil
	case uint32:
		return uint64(castedVal), nil
	case uint16:
		return uint64(castedVal), nil
	case uint8:
		return uint64(castedVal), nil
	case float32:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint64(castedVal), nil
	case float64:
		if castedVal < 0 {
			return 0, errNumericOverFlow
		}
		return uint64(castedVal), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", value, value)
	}
}

func TryInt8(value interface{}) (int8, error) {
	value = indirect(value)

	switch castedVal := value.(type) {
	case int:
		if castedVal < math.MinInt8 || castedVal > math.MaxInt8 {
			return 0, errNumericOverFlow
		}
		return int8(castedVal), nil
	case int64:
		if castedVal < math.MinInt8 || castedVal > math.MaxInt8 {
			return 0, errNumericOverFlow
		}
		return int8(castedVal), nil
	case int32:
		if castedVal < math.MinInt8 || castedVal > math.MaxInt8 {
			return 0, errNumericOverFlow
		}
		return int8(castedVal), nil
	case int16:
		if castedVal < math.MinInt8 || castedVal > math.MaxInt8 {
			return 0, errNumericOverFlow
		}
		return int8(castedVal), nil
	case int8:
		return castedVal, nil
	case uint:
		if castedVal > math.MaxInt8 {
			return 0, errNumericOverFlow
		}
		return int8(castedVal), nil
	case uint64:
		if castedVal > math.MaxInt8 {
			return 0, errNumericOverFlow
		}
		return int8(castedVal), nil
	case uint32:
		if castedVal > math.MaxInt8 {
			return 0, errNumericOverFlow
		}
		return int8(castedVal), nil
	case uint16:
		if castedVal > math.MaxInt8 {
			return 0, errNumericOverFlow
		}
		return int8(castedVal), nil
	case uint8:
		if castedVal > math.MaxInt8 {
			return 0, errNumericOverFlow
		}
		return int8(castedVal), nil
	case float64:
		if castedVal < math.MinInt8 || castedVal > math.MaxInt8 {
			return 0, errNumericOverFlow
		}
		return int8(castedVal), nil
	case float32:
		if castedVal < math.MinInt8 || castedVal > math.MaxInt8 {
			return 0, errNumericOverFlow
		}
		return int8(castedVal), nil
	case string:
		v, err := strconv.ParseInt(castedVal, 0, 0)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int8", value, value)
		}
		if v < math.MinInt8 || v > math.MaxInt8 {
			return 0, errNumericOverFlow
		}
		return int8(v), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", value, value)
	}
}

func TryInt16(value interface{}) (int16, error) {
	value = indirect(value)

	switch castedVal := value.(type) {
	case int:
		if castedVal < math.MinInt16 || castedVal > math.MaxInt16 {
			return 0, errNumericOverFlow
		}
		return int16(castedVal), nil
	case int64:
		if castedVal < math.MinInt16 || castedVal > math.MaxInt16 {
			return 0, errNumericOverFlow
		}
		return int16(castedVal), nil
	case int32:
		if castedVal < math.MinInt16 || castedVal > math.MaxInt16 {
			return 0, errNumericOverFlow
		}
		return int16(castedVal), nil
	case int16:
		return castedVal, nil
	case int8:
		return int16(castedVal), nil
	case uint:
		if castedVal > math.MaxInt16 {
			return 0, errNumericOverFlow
		}
		return int16(castedVal), nil
	case uint64:
		if castedVal > math.MaxInt16 {
			return 0, errNumericOverFlow
		}
		return int16(castedVal), nil
	case uint32:
		if castedVal > math.MaxInt16 {
			return 0, errNumericOverFlow
		}
		return int16(castedVal), nil
	case uint16:
		if castedVal > math.MaxInt16 {
			return 0, errNumericOverFlow
		}
		return int16(castedVal), nil
	case uint8:
		return int16(castedVal), nil
	case float64:
		if castedVal < math.MinInt16 || castedVal > math.MaxInt16 {
			return 0, errNumericOverFlow
		}
		return int16(castedVal), nil
	case float32:
		if castedVal < math.MinInt16 || castedVal > math.MaxInt16 {
			return 0, errNumericOverFlow
		}
		return int16(castedVal), nil
	case string:
		v, err := strconv.ParseInt(castedVal, 0, 0)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int16", value, value)
		}
		if v < math.MinInt16 || v > math.MaxInt16 {
			return 0, errNumericOverFlow
		}
		return int16(v), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", value, value)
	}
}

func TryInt32(value interface{}) (int32, error) {
	value = indirect(value)

	switch castedVal := value.(type) {
	case int:
		if castedVal < math.MinInt32 || castedVal > math.MaxInt32 {
			return 0, errNumericOverFlow
		}
		return int32(castedVal), nil
	case int64:
		if castedVal < math.MinInt32 || castedVal > math.MaxInt32 {
			return 0, errNumericOverFlow
		}
		return int32(castedVal), nil
	case int32:
		return castedVal, nil
	case int16:
		return int32(castedVal), nil
	case int8:
		return int32(castedVal), nil
	case uint:
		if castedVal > math.MaxInt32 {
			return 0, errNumericOverFlow
		}
		return int32(castedVal), nil
	case uint64:
		if castedVal > math.MaxInt32 {
			return 0, errNumericOverFlow
		}
		return int32(castedVal), nil
	case uint32:
		if castedVal > math.MaxInt32 {
			return 0, errNumericOverFlow
		}
		return int32(castedVal), nil
	case uint16:
		return int32(castedVal), nil
	case uint8:
		return int32(castedVal), nil
	case float64:
		if castedVal < math.MinInt32 || castedVal > math.MaxInt32 {
			return 0, errNumericOverFlow
		}
		return int32(castedVal), nil
	case float32:
		if castedVal < math.MinInt32 || castedVal > math.MaxInt32 {
			return 0, errNumericOverFlow
		}
		return int32(castedVal), nil
	case string:
		v, err := strconv.ParseInt(castedVal, 0, 0)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int32", value, value)
		}
		if v < math.MinInt32 || v > math.MaxInt32 {
			return 0, errNumericOverFlow
		}
		return int32(v), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", value, value)
	}
}

func TryInt64(value interface{}) (int64, error) {
	value = indirect(value)

	switch castedVal := value.(type) {
	case int:
		return int64(castedVal), nil
	case int64:
		return castedVal, nil
	case int32:
		return int64(castedVal), nil
	case int16:
		return int64(castedVal), nil
	case int8:
		return int64(castedVal), nil
	case uint:
		return int64(castedVal), nil
	case uint64:
		if castedVal > math.MaxInt64 {
			return 0, errNumericOverFlow
		}
		return int64(castedVal), nil
	case uint32:
		return int64(castedVal), nil
	case uint16:
		return int64(castedVal), nil
	case uint8:
		return int64(castedVal), nil
	case float64:
		if castedVal < math.MinInt64 || castedVal > math.MaxInt64 {
			return 0, errNumericOverFlow
		}
		return int64(castedVal), nil
	case float32:
		return int64(castedVal), nil
	case string:
		v, err := strconv.ParseInt(castedVal, 0, 0)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int64", value, value)
		}
		//nolint:staticcheck
		if v < math.MinInt64 || v > math.MaxInt64 {
			return 0, errNumericOverFlow
		}
		return v, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", value, value)
	}
}

func TryFloat32(value interface{}) (float32, error) {
	value = indirect(value)

	switch castedVal := value.(type) {
	case float64:
		if castedVal < -math.MaxFloat32 || castedVal > math.MaxFloat32 {
			return 0, errNumericOverFlow
		}
		return float32(castedVal), nil
	case float32:
		return castedVal, nil
	case int:
		if float64(castedVal) < -math.MaxFloat32 || float64(castedVal) > math.MaxFloat32 {
			return 0, errNumericOverFlow
		}
		return float32(castedVal), nil
	case int64:
		if float64(castedVal) < -math.MaxFloat32 || float64(castedVal) > math.MaxFloat32 {
			return 0, errNumericOverFlow
		}
		return float32(castedVal), nil
	case int32:
		if float64(castedVal) < -math.MaxFloat32 || float64(castedVal) > math.MaxFloat32 {
			return 0, errNumericOverFlow
		}
		return float32(castedVal), nil
	case int16:
		return float32(castedVal), nil
	case int8:
		return float32(castedVal), nil
	case uint:
		return float32(castedVal), nil
	case uint64:
		if float64(castedVal) > math.MaxFloat32 {
			return 0, errNumericOverFlow
		}
		return float32(castedVal), nil
	case uint32:
		if float64(castedVal) > math.MaxFloat32 {
			return 0, errNumericOverFlow
		}
		return float32(castedVal), nil
	case uint16:
		return float32(castedVal), nil
	case uint8:
		return float32(castedVal), nil
	case string:
		v, err := strconv.ParseFloat(castedVal, 32)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to float32", value, value)
		}
		if v < -math.MaxFloat32 || v > math.MaxFloat32 {
			return 0, errNumericOverFlow
		}
		return float32(v), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", value, value)
	}
}

func TryFloat64(value interface{}) (float64, error) {
	value = indirect(value)

	switch castedVal := value.(type) {
	case float64:
		return castedVal, nil
	case float32:
		return float64(castedVal), nil
	case int:
		return float64(castedVal), nil
	case int64:
		return float64(castedVal), nil
	case int32:
		return float64(castedVal), nil
	case int16:
		return float64(castedVal), nil
	case int8:
		return float64(castedVal), nil
	case uint:
		return float64(castedVal), nil
	case uint64:
		return float64(castedVal), nil
	case uint32:
		return float64(castedVal), nil
	case uint16:
		return float64(castedVal), nil
	case uint8:
		return float64(castedVal), nil
	case string:
		v, err := strconv.ParseFloat(castedVal, 64)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to float64", value, value)
		}
		if v < -math.MaxFloat64 || v > math.MaxFloat64 {
			return 0, errNumericOverFlow
		}
		return v, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", value, value)
	}
}

func TryString(value interface{}) (string, error) {
	value = indirectToStringerOrError(value)

	switch s := value.(type) {
	case string:
		return s, nil
	case []byte:
		return string(s), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(s, 10), nil
	case int:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int32:
		return strconv.FormatInt(int64(s), 10), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case time.Time:
		return s.Format(time.RFC3339), nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", value, value)
	}
}

func TryDate(value interface{}) (time.Time, error) {
	value = indirect(value)

	switch v := value.(type) {
	case time.Time:
		return v, nil
	case string:
		return time.Parse(dateFormat, v)
	default:
		return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to Time", value, value)
	}
}

func TryDateTime(value interface{}) (time.Time, error) {
	value = indirect(value)

	switch v := value.(type) {
	case time.Time:
		return v, nil
	case string:
		return time.Parse(dateTimeFormat, v)
	default:
		return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to Time", value, value)
	}
}

func TryUUID(value interface{}) (string, error) {
	uuidStr, err := TryString(value)
	if err != nil {
		return uuidStr, err
	}

	_, err = uuid.Parse(uuidStr)
	if err != nil {
		return "", err
	}

	return uuidStr, nil
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirectToStringerOrError returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
// or error,
func indirectToStringerOrError(a interface{}) interface{} {
	if a == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}
