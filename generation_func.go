package epoch

// Ensure that GenerationFunc implements Generator interface
var _ Generator = (GenerationFunc)(nil)

// GenerationFunc is a function type that implements Generator interface
type GenerationFunc func() Epoch

func (fn GenerationFunc) GenerateEpoch() Epoch { return fn() }
