package asyncloop

import "math"

type Range struct {
	Start, End int
}

func Ranges(n, size int) []Range {
	if n <= 0 || size <= 0 {
		return nil
	}

	rs := make([]Range, 0, int(math.Ceil(float64(n)/float64(size))))
	for start := 0; start < n; start += size {
		end := start + size
		if end > n {
			end = n
		}
		rs = append(rs, Range{Start: start, End: end})
	}
	return rs
}
