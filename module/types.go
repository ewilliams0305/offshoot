package offshoot

// The convertable interface provides a function to convert a Type [TInput] to another type [TOutput]
type Mappable[TInput any, TOutput any] interface {
	Map(mapper Map[TInput, TOutput])
}
