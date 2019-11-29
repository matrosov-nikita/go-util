package util

// StringPtr return pointer to string s
func StringPtr(s string) *string {
	return &s
}

// Int64Ptr return pointer to int64 i
func Int64Ptr(i int64) *int64 {
	return &i
}

// Float64Ptr return pointer to int64 d
func Float64Ptr(f float64) *float64 {
	return &f
}
