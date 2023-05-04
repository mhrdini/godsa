package containers

type Container[T any] interface {
	Size() int
	Empty() bool
	Values() []T
	String() string
}
