package utils

var (
	primeNumbers = buildMap([]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97})
)

func buildMap(values []int) map[int]struct{} {
	var resultMap = make(map[int]struct{})
	for _, value := range values {
		resultMap[value] = struct{}{}
	}

	return resultMap
}

func Filter(source []int) []int {
	var filtered = make([]int, 0)

	for _, value := range source {
		if _, contains := primeNumbers[value]; !contains {
			filtered = append(filtered, value)
		}
	}
	return filtered
}
