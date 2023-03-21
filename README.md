# epoch
Epoch is a lightweight library for time generation

## Installation:

> go get github.com/flavioltonon/epoch

## Usage:

You can use one of the available epoch.Generator implementations or implement your own. Example:

```Go
type MyGenerator struct {}

func (g *MyGenerator) GenerateEpoch() epoch.Epoch {
    return time.Date(2023, 3, 21, 17, 5, 12, 0, time.Local)
}
```

## Running tests:

> make tests