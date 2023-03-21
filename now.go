package epoch

import "time"

// NowGenerator is a Generator which is capable of generating time.Time values based on time.Now function of Go's standard library
var NowGenerator = GenerationFunc(time.Now)
