package pointer

// IndirectGet возвращает значение разыменованного указателя ptr, если он не nil, в противном случае вернет нулевое
// значение типа T.
func IndirectGet[T any](ptr *T) T {
	if ptr == nil {
		var z T

		return z
	}

	return *ptr
}

// Pointer возвращает указатель на значение
// значение типа T
func Pointer[T any](value T) *T {
	return &value
}

// IndirectGetWithDefault возвращает значение разыменованного указателя ptr, если он не nil, в противном случае вернет
// нулевое дефолтное значение deft.
func IndirectGetWithDefault[T any](ptr *T, deft T) T {
	if ptr == nil {
		return deft
	}

	return *ptr
}

// ApplyIfNotNil Применить действие с value, если он не nil.
func ApplyIfNotNil[R any, T any](value *T, f func(T) R) {
	if value != nil {
		f(*value)
	}
}

func ApplyNillable[R any, T comparable](value *T, setNillable func(*T) R, clear func() R) {
	if value != nil {
		var z T
		if *value == z {
			clear()
		} else {
			setNillable(value)
		}
	}
}
