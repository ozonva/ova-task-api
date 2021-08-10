package utils

func InvertMap(source map[string]string) map[string]string {
	invertedMap := make(map[string]string, len(source))
	for key, value := range source {
		if _, contains := invertedMap[value]; contains {
			panic("An item with the same key [" + value + "] has already been added.")
		}
		invertedMap[value] = key
	}
	return invertedMap
}
