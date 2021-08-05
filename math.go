package logger

//MaxInt32 returns the max of two int32
func MaxInt32(a, b int) int {
	r := a
	if b > a {
		r = b
	}
	return r
}

//MaxInt64 returns the max of two int64
func MaxInt64(a, b int64) int64 {
	r := a
	if b > a {
		r = b
	}
	return r
}
