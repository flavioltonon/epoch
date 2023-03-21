package epoch

// Generator defines the behaviors of a epoch generator
type Generator interface {
	// GenerateEpoch returns a new Epoch
	GenerateEpoch() Epoch
}
