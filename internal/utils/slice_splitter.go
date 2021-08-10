package utils

import "math"

func SplitSlice(source []string, batch int) [][]string {
	if batch <= 0 {
		panic("Batch must be greater than 0.")
	}
	chunksCount := int(math.Ceil(float64(len(source)) / float64(batch)))
	if chunksCount == 0 {
		return [][]string{}
	}

	chunks := make([][]string, 0, chunksCount)
	for i := 0; i < len(source); i += batch {
		from := i
		to := from + batch
		if to > len(source) {
			to = len(source)
		}
		chunks = append(chunks, source[from:to])
	}
	return chunks
}
