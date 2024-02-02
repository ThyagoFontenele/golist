package golist

type List[T any] []T

func (l List[T]) Where(fn func(v T, index int) bool) []T {
	var newValue []T
	for index, value := range l {
		if fn(value, index) {
			newValue = append(newValue, value)
		}
	}
	return newValue
}

func (l List[T]) ForEach(fn func(v T, index int)) {
	for index, value := range l {
		fn(value, index)
	}
}

func (l List[T]) Filter(fn func(v T, index int) bool) []T {
	var newValue []T
	for index, value := range l {
		if fn(value, index) {
			newValue = append(newValue, value)
		}
	}
	return newValue
}

func (l *List[T]) Push(value T) {
	*l = append(*l, value)
}

func (l *List[T]) Concat(value []T) {
	for _, value := range value {
		*l = append(*l, value)
	}
}

func (l *List[T]) Pop() {
	*l = (*l)[0 : len(*l)-1]
}

func (l List[T]) First(fn func(v T, index int) bool) T {
	var defaultValue T

	for index, value := range l {
		if fn(value, index) {
			return value
		}
	}

	return defaultValue
}

func (l *List[T]) Splice(elementIndex int, n int) []T {
	var newList []T
	var elements []T
	for index, value := range *l {
		if index >= elementIndex && index < elementIndex+n {
			elements = append(elements, value)
			continue
		}
		newList = append(newList, value)
	}
	*l = newList

	return elements
}
