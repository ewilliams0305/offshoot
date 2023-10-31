package offshoot

// Mappable is an interface that defines the behavior for mapping elements from one type to another.
// A Mappable must be able to convert an input type to an output type.
type Mappable[TInput any, TOutput any] interface {
	Map(input TInput) TOutput
}
