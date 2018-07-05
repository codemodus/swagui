# swagui

    go get -u github.com/codemodus/swagui

Package swagui simplifies serving an instance of Swagger-UI.

## Usage

```go
type Options
type Swagui
    func New(opts *Options) (*Swagui, error)
    func (s *Swagui) Handler() http.Handler
type Version
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

    if err = http.ListenAndServe(":21234", ui.Handler("")); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }

    // ...
}
```

### With A Multiplexer

```go
func main() {
    // ...

    m := http.NewServeMux()
    // ...

    m.Handle("/some_path", someHandler)
    m.Handle(
        "/docs/", 
        http.StripPrefix("/docs/", ui.Handler("/defs/swagger.json")),
    )
    
    if err = http.ListenAndServe(":21234", m); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
    	
    // ...
}
```

## More Info

N/A

## Documentation

View the [GoDoc](http://godoc.org/github.com/codemodus/swagui)

## Benchmarks

N/A
