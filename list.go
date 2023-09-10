package list

type List[T any] []T

func (l List[T]) Map(fn func(v T, index int) T) []T {
	var newValue []T
	for index, value := range l {
		newValue = append(newValue, fn(value, index))
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

func (l *List[T]) Pop() {
	var newList []T
	for index, value := range *l {
		if index+1 == len(*l) {
			break
		}
		newList = append(newList, value)
	}
	*l = newList
}

func (l List[T]) First(fn func(v T, index int) bool) T {
	var firstValue T
	for index, value := range l {
		if fn(value, index) {
			firstValue = value
			break
		}
	}
	return firstValue
}
