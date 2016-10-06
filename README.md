# swagui

    go get -u github.com/codemodus/swagui

Package swagui simplifies serving an instance of Swagger-UI. It can be added to 
a multiplexer, or served directly. If using a multiplexer, the path prefix 
option must match the relevant route.

## Usage

```go
type Options
type Swagui
    func New(opts *Options) (*Swagui, error)
    func (s *Swagui) Handler() http.Handler
    func (s *Swagui) PathPrefix() string
```

### Setup

```go
import (
    // ...

    "github.com/codemodus/swagui"
)

func main() {
    // ...
    
    ui, err := swagui.New(nil)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    if err = http.ListenAndServe(":21234", ui.Handler()); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    // ...
}
```

### With A Multiplexer

```go
func main() {
    // ...

    m := http.NewServeMux()
    m.Handle("/some_path", someHandler)
    m.Handle(ui.PathPrefix(), ui.Handler())
    
    if err = http.ListenAndServe(":21234", m); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    	
    // ...
}
```

## More Info

### Swagger Version

If an invalid version is configured (0, or > current release), the current
release will be used.

## Documentation

View the [GoDoc](http://godoc.org/github.com/codemodus/swagui)

## Benchmarks

N/A
