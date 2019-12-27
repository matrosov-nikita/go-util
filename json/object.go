package json

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/itimofeev/go-util/cast"
)

const (
	keySep            = "."
	flatSep           = "_"
	quotedFlatSep     = "__"
	defaultTimeFormat = "2006-01-02T15:04:05.999999999"
)

type Object map[string]interface{}

func NewObject() Object {
	return make(map[string]interface{})
}

func (jo Object) GetField(key string) interface{} {
	return jo.deepGet(parseDeepKey(key))
}

func (jo Object) PutField(key string, val interface{}) interface{} {
	return jo.deepPut(parseDeepKey(key), val)
}

func (jo Object) GetFieldAsString(key string) string {
	casted, err := cast.TryString(jo.GetField(key))
	if err == nil {
		return casted
	}
	return ""
}

func (jo Object) GetFieldAsInt(key string) int {
	casted, err := cast.TryInt64(jo.GetField(key))
	if err == nil {
		return int(casted)
	}
	return 0
}

func (jo Object) GetFieldAsTime(key string, format ...string) *time.Time {
	val := jo.GetField(key)
	if val != nil {
		switch fieldVal := val.(type) {
		case string:
			timeFormat := defaultTimeFormat
			if len(format) > 0 {
				timeFormat = format[0]
			}
			parsed, err := time.Parse(timeFormat, fieldVal)
			if err != nil {
				return nil
			}
			return &parsed
		case time.Time:
			return &fieldVal
		}
	}
	return nil
}

func (jo Object) GetFieldAsObject(key string) Object {
	val := jo.GetField(key)
	if val != nil {
		switch fieldVal := val.(type) {
		case Object:
			return fieldVal
		case map[string]interface{}:
			return fieldVal
		}
	}
	return nil
}

func (jo Object) GetFieldAsUUID(key string) string {
	casted, err := cast.TryUUID(jo.GetField(key))
	if err == nil {
		return casted
	}
	return ""
}

func (jo Object) GetFieldAsInt8(key string) int8 {
	casted, err := cast.TryInt8(jo.GetField(key))
	if err == nil {
		return casted
	}
	return 0
}

func (jo Object) GetFieldAsInt16(key string) int16 {
	casted, err := cast.TryInt16(jo.GetField(key))
	if err == nil {
		return casted
	}
	return 0
}

func (jo Object) GetFieldAsInt32(key string) int32 {
	casted, err := cast.TryInt32(jo.GetField(key))
	if err == nil {
		return casted
	}
	return 0
}

func (jo Object) GetFieldAsInt64(key string) int64 {
	casted, err := cast.TryInt64(jo.GetField(key))
	if err == nil {
		return casted
	}
	return 0
}

func (jo Object) GetFieldAsUint8(key string) uint8 {
	casted, err := cast.TryUInt8(jo.GetField(key))
	if err == nil {
		return casted
	}
	return 0
}

func (jo Object) GetFieldAsUint16(key string) uint16 {
	casted, err := cast.TryUInt16(jo.GetField(key))
	if err == nil {
		return casted
	}
	return 0
}

func (jo Object) GetFieldAsUint32(key string) uint32 {
	casted, err := cast.TryUInt32(jo.GetField(key))
	if err == nil {
		return casted
	}
	return 0
}

func (jo Object) GetFieldAsUint64(key string) uint64 {
	casted, err := cast.TryUInt64(jo.GetField(key))
	if err == nil {
		return casted
	}
	return 0
}

func (jo Object) GetFieldAsFloat32(key string) float32 {
	casted, err := cast.TryFloat32(jo.GetField(key))
	if err == nil {
		return casted
	}
	return 0
}

func (jo Object) GetFieldAsFloat64(key string) float64 {
	casted, err := cast.TryFloat64(jo.GetField(key))
	if err == nil {
		return casted
	}
	return 0
}

func (jo Object) MustGetFieldAsString(key string) (string, error) {
	return cast.TryString(jo.GetField(key))
}

func (jo Object) MustGetFieldAsInt(key string) (int, error) {
	field, err := cast.TryInt64(jo.GetField(key))
	return int(field), err
}

func (jo Object) MustGetFieldAsTime(key string) (time.Time, error) {
	return cast.TryDateTime(jo.GetField(key))
}

func (jo Object) MustGetFieldAsUUID(key string) (string, error) {
	return cast.TryUUID(jo.GetField(key))
}

func (jo Object) MustGetFieldAsInt8(key string) (int8, error) {
	return cast.TryInt8(jo.GetField(key))
}

func (jo Object) MustGetFieldAsInt16(key string) (int16, error) {
	return cast.TryInt16(jo.GetField(key))
}

func (jo Object) MustGetFieldAsInt32(key string) (int32, error) {
	return cast.TryInt32(jo.GetField(key))
}

func (jo Object) MustGetFieldAsInt64(key string) (int64, error) {
	return cast.TryInt64(jo.GetField(key))
}

func (jo Object) MustGetFieldAsUint8(key string) (uint8, error) {
	return cast.TryUInt8(jo.GetField(key))
}

func (jo Object) MustGetFieldAsUint16(key string) (uint16, error) {
	return cast.TryUInt16(jo.GetField(key))
}

func (jo Object) MustGetFieldAsUint32(key string) (uint32, error) {
	return cast.TryUInt32(jo.GetField(key))
}

func (jo Object) MustGetFieldAsUint64(key string) (uint64, error) {
	return cast.TryUInt64(jo.GetField(key))
}

func (jo Object) MustGetFieldAsFloat32(key string) (float32, error) {
	return cast.TryFloat32(jo.GetField(key))
}

func (jo Object) MustGetFieldAsFloat64(key string) (float64, error) {
	return cast.TryFloat64(jo.GetField(key))
}

func (jo Object) Remove(key string) {
	delete(jo, key)
}

func (jo Object) Put(key string, val interface{}) {
	jo[key] = val
}

func (jo Object) JSON() ([]byte, error) {
	return json.Marshal(jo)
}

func (jo Object) OmitEmpty() Object {
	for key, val := range jo {
		if val == nil {
			delete(jo, key)
		}
	}

	return jo
}

func (jo Object) OmitKey(key ...string) Object {
	for _, k := range key {
		delete(jo, k)
	}
	return jo
}

func (jo Object) Flatten(delim ...string) Object {
	if len(jo) == 0 {
		return jo
	}

	delimiter := flatSep
	if len(delim) > 0 {
		delimiter = delim[0]
	}

	flatten := make(map[string]interface{})
	for field, val := range jo {
		var oVal Object
		switch castedVal := val.(type) {
		case Object:
			oVal = castedVal
		case map[string]interface{}:
			oVal = castedVal
		default:
			flatten[screenDelimiter(field)] = val
			continue
		}
		flattenVal := oVal.Flatten(delimiter)
		for f, v := range flattenVal {
			flatten[screenDelimiter(field)+delimiter+f] = v
		}
	}

	return flatten
}

func (jo Object) Nested() Object {
	if len(jo) == 0 {
		return jo
	}

	nested := NewObject()
	for field, val := range jo {
		nested.deepPut(SplitFlatKey(field), val)
	}

	return nested
}

func FlattenField(field string) string {
	flatten := field
	if strings.Contains(flatten, flatSep) {
		flatten = strings.ReplaceAll(flatten, flatSep, quotedFlatSep)
	}

	return strings.ReplaceAll(flatten, ".", flatSep)
}

func screenDelimiter(field string) string {
	return strings.ReplaceAll(field, flatSep, "__")
}

func (jo Object) deepGet(key []string) interface{} {
	if len(key) > 1 {
		val, ok := jo[key[0]]
		if !ok {
			return nil
		}

		switch fieldVal := val.(type) {
		case Object:
			return fieldVal.deepGet(key[1:])
		case map[string]interface{}:
			return Object(fieldVal).deepGet(key[1:])
		default:
			return nil
		}
	}

	return jo[key[0]]
}

func (jo Object) deepPut(key []string, val interface{}) Object {
	if len(key) > 1 {
		var levelObj Object
		foundLevelObj, ok := jo[key[0]]
		if ok {
			switch castedLevelObj := foundLevelObj.(type) {
			case Object:
				levelObj = castedLevelObj
			case map[string]interface{}:
				levelObj = castedLevelObj
			default:
				levelObj = NewObject()
			}
		} else {
			levelObj = NewObject()
		}

		jo.Put(key[0], levelObj.deepPut(key[1:], val))
		return jo
	}

	jo.Put(key[0], val)
	return jo
}

func parseDeepKey(key string) []string {
	return strings.Split(key, keySep)
}

func SplitFlatKey(key string) []string {
	if !strings.Contains(key, flatSep) {
		return []string{key}
	}

	maxItemsCount := strings.Count(key, flatSep)
	indexes := make([]int, 0, maxItemsCount)
	prevIsSep := false
	for i := 0; i < len(key); i++ {
		if key[i] == '_' {
			if i == 0 {
				continue
			}
			if i == len(key)-1 {
				break
			}
			if key[i+1] != '_' && !prevIsSep {
				indexes = append(indexes, i)
			}
			prevIsSep = true
		} else {
			prevIsSep = false
		}
	}
	split := make([]string, 0, len(indexes)+1)
	start := 0
	for _, idx := range indexes {
		split = append(split, strings.ReplaceAll(key[start:idx], quotedFlatSep, flatSep))
		start = idx + 1
	}
	split = append(split, strings.ReplaceAll(key[start:], quotedFlatSep, flatSep))
	return split
}

// Value returns Object value marshalled to JSON for storing in the database
func (jo Object) Value() (driver.Value, error) {
	if len(jo) == 0 {
		return nil, nil
	}

	return jo.JSON()
}

// Scan scans Object from specified data from the database
func (jo *Object) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	return json.Unmarshal(bytes, &jo)
}
