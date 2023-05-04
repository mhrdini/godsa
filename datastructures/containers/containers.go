package containers

type Container[T any] interface {
	Name() string
	Size() int
	Empty() bool
	Values() []T
	String() string
	Reset()
}
