package epoch

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerationFunc_GenerateEpoch(t *testing.T) {
	tests := []struct {
		name string
		fn   GenerationFunc
		want Epoch
	}{
		{
			name: "Given a GenerationFunc, a new Epoch should be returned when GenerateEpoch is called",
			fn: func() Epoch {
				return time.Date(2022, 3, 21, 16, 53, 43, 123, time.UTC)
			},
			want: time.Date(2022, 3, 21, 16, 53, 43, 123, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fn.GenerateEpoch())
		})
	}
}
