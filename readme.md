# random-str

random-str returns a random string.

### FUNCTIONS

```go
// WriteFromFormat generates a string that follow the given format
func WriteFromFormat(format string) string
```

```go
// WriteN generates a string with length n. if format is specified, the string
// will be of the given format
func WriteN(n int, format ...rune) string
    
```

### Format

`A`: return only uppercase `(ABC)`
`a`: return only lowercase `(abc)`
`d`: return only digits `(123)`
`s`: return only symbols `(&^%)`
`*`: return only any `(Aa1*)`

[examples](./example)