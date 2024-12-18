package utils

type Array2D[T any] struct {
	array  []T
	Width  int
	Height int
}

func NewArray2D[T any](width, height int) Array2D[T] {
	return Array2D[T]{
		array:  make([]T, width*height),
		Width:  width,
		Height: height,
	}
}

func (arr *Array2D[T]) Contains(x, y int) bool {
	return x >= 0 && x < arr.Width && y >= 0 && y < arr.Height
}

func (arr *Array2D[T]) Get(x, y int) T {
	return arr.array[y*arr.Width+x]
}

func (arr *Array2D[T]) Set(x, y int, nv T) {
	arr.array[y*arr.Width+x] = nv
}
