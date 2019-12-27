package json

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestObject_Nested(t *testing.T) {
	inNestedStr := `{"type":"detector","event":{"detector":{"faceAppeared":{"age":44,"gender":2,"phase":"happened","quality":0.72228586673736572,"rectangle":{"h":0.56259259259259253,"index":78,"w":0.3164583333333334,"x":0.53187499999999999,"y":0.12888888888888889},"time_begin":{"datetime":"2019-03-27T11:10:14.640000","utc":"2019-03-27T08:10:14.640000"}},"type":{"id":"faceAppeared","name":"faceAppeared"}}},"version":1,"id":"0d49659f-1edc-49f2-872a-5ead1db8390a","time":{"datetime":"2019-03-27T11:10:14.920000","utc":"2019-03-27T08:10:14.920000"},"source":{"detector":{"id":"A-SHAULUKHOV/AVDetector.1/EventSupplier","name":""},"server":{"id":"A-SHAULUKHOV","name":""},"video":{"id":"A-SHAULUKHOV/DeviceIpint.1/SourceEndpoint.video:0:0","name":"Camera"}}}`
	var inNested Object
	require.NoError(t, json.Unmarshal([]byte(inNestedStr), &inNested))
	flatten := inNested.Flatten()
	outNested := flatten.Nested()
	outNestedStr, err := json.Marshal(outNested)
	require.NoError(t, err)
	require.JSONEq(t, inNestedStr, string(outNestedStr))
}

func TestSplitFlatKey(t *testing.T) {
	require.Equal(t, []string{""}, SplitFlatKey(""))
	require.Equal(t, []string{"event"}, SplitFlatKey("event"))
	require.Equal(t, []string{"_listed"}, SplitFlatKey("_listed"))
	require.Equal(t, []string{"listed_"}, SplitFlatKey("listed_"))
	require.Equal(t, []string{"event", "detector"}, SplitFlatKey("event_detector"))
	require.Equal(t, []string{"time_utc"}, SplitFlatKey("time__utc"))
	require.Equal(t, []string{"listed_face_detected"}, SplitFlatKey("listed__face__detected"))
	require.Equal(t, []string{"event", "detector", "faceAppeared", "time_begin", "datetime"}, SplitFlatKey("event_detector_faceAppeared_time__begin_datetime"))
	require.Equal(t, []string{"event", "detector", "listed_face_detected", "rectangle", "x"}, SplitFlatKey("event_detector_listed__face__detected_rectangle_x"))
}

func TestGetValueForField(t *testing.T) {
	obj := Object{"test": 10}
	require.Equal(t, int64(10), obj.GetFieldAsInt64("test"))
	require.Equal(t, int64(0), obj.GetFieldAsInt64("unknown field"))
}
