package interfaces

// Tracing modules interface, using for dynamic modules
type Tracing interface {
	New() Tracing
	Init(Option)
	Closing() (err error)
}
