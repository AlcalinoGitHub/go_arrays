package array

//Returns wether an array contains a certain element
func Contains[T comparable](array []T, element T) bool {
	for i := range array {
		if array[i] == element {
			return true
		}
	}
	return false
}

//Returns the amount of elements on the array equal to the one passed in
func Count[T comparable](array []T, element T) int {
	count := 0
	for i := range array {
		if array[i] == element {
			count++
		}
	}
	return count
}

//Return the index of the element if it's present. If element is absent returns null
func IndexOf[T comparable](array []T, element T) *int {
	for i := range array {
		if array[i] == element {
			return &i
		}
	}
	return nil
}

type Number interface {
	int | int8 | int16 | int64 | int32 | uint | uint16 | uint32 | uint64 | uint8
}

//Returns the sum of all elements in the array.
func Sum[T Number](array []T) T {
	var sum T = 0
	for _, num := range array {
		sum += num
	}
	return sum
}

//Returns the largest element in the array. If empty returns null
func Max[T Number](array []T) *T {
	if len(array) == 0 {
		return nil
	}
	var max = array[0]

	for _, num := range array {
		if num > max {
			max = num
		}
	}

	return &max
}

//Returns the smallest element in the array. If empty returns null
func Min[T Number](array []T) *T {
	if len(array) == 0 {
		return nil
	}
	var min = array[0]

	for _, num := range array {
		if num < min {
			min = num
		}
	}

	return &min
}


func Reverse[T any](array []T) []T {
	final := make([]T, 0)
	for i := len(array) -1; i >= 0; i-- {
		final = append(final, array[i])
	}
	return final
}


func Filter[T any](array []T, function func(T) bool) []T {
	result := make([]T, 0)
	for _, element := range array {
		if function(element) {result = append(result, element)}
	}
	return result
}

func Map[T any, K any](array []T, function func(any) K) []K {
	result := make([]K, 0)

	for _, element := range array {
		result = append(result, function(element))
	}
	return result
}