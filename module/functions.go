package offshoot

// A mapper function is used to convert the type stored by the offshoot.
// The mapper function should take in a value and convert to the output type specified.
type MapperFunction[TInput any, TOutput any] func(input TInput) TOutput

// An ensure function is used to validate the status of the value stored by the [offshoot]
type EnsureFunction[TInput any] func(input TInput) bool

// A tap function is used as an opertunity to access the value stored inside the [offshoot]
type TapFunction[TValue any] func(value TValue)

// The error handle function provides the consumer an opportunity 
// to handle an error and return a new offshoot or either failure or success
type ErrorHandle[TValue any] func(error Error) Offshoot[TValue]

// The value handle function provides the consuler an opportunity 
// to handle an value and create a new Offshoot 
type ValueHandle[TValue any] func(value TValue) Offshoot[TValue]