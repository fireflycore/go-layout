package array

func Map[T any, R any](array []T, fn func(index int, item T) R) []R {
	var temp []R
	for index, item := range array {
		temp = append(temp, fn(index, item))
	}
	return temp
}

func Filter[T any](array []T, fn func(index int, item T) bool) []T {
	var temp []T
	for index, item := range array {
		if fn(index, item) {
			temp = append(temp, item)
		}
	}
	return temp
}

func Unique[T any](array []T, fn func(index int, item T) string) []T {
	set := make(map[string]T, len(array))
	flag := 0
	var key string

	for index, value := range array {
		key = fn(index, value)
		if _, ok := set[key]; !ok {
			set[key] = value
			array[flag] = value
			flag++
		}
	}

	return array
}
