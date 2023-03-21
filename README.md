# epoch
Epoch is a lightweight library for time generation

## Motivation

Flows that depend on time values such as `time.Now()` are hard to test alone. Epoch provides a ready-to-use toolkit to make reduce the coupling between time and test scenarios, allowing the manipulation of time generation functions according to your use case.

## Installation

> go get github.com/flavioltonon/epoch

## Usage

You can use one of the available epoch.Generator implementations or implement your own:

```Go
// MyGenerator is a custom epoch.Generator implementation
type MyGenerator struct {}

func (g *MyGenerator) GenerateEpoch() epoch.Epoch {
    return time.Date(2023, 3, 21, 17, 5, 12, 0, time.Local)
}
```

After that, you can use epoch.Generator interface to inject your generator as a dependency and you are ready to go!

```Go
type MyStruct struct {
    epochGenerator epoch.Generator
}

func NewMyStruct(epochGenerator epoch.Generator) *MyStruct {
    return &MyStruct{epochGenerator: epochGenerator}
}

func (s *MyStruct) DoSomething() {
    date := s.epochGenerator.GenerateEpoch()
    
    // do something
}
```

## Other examples

### Using epoch.NowGenerator

```Go
// epoch.NowGenerator wraps time.Now into an epoch.GenerationFunc, implementing epoch.Generator interface
myGenerator := epoch.NowGenerator
```

### Implementing a custom epoch.GenerationFunc

```Go
myGenerator := epoch.GenerationFunc(func() epoch.Epoch {
    // do something
    // return some time.Time value
})
```

## Running tests

> make tests
