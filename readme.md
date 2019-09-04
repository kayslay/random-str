# Package random-str

[![CircleCI](https://circleci.com/gh/kayslay/random-str.svg?style=svg)](https://circleci.com/gh/kayslay/random-str)

Package random-str generates random strings that follow a format.

Example usage:

```go
s = str.WriteN(10, 'A') // s will be "NBHTFGREDC" (all caps)
```

### Format

random-str can create random strings that conform to a particular format. The available formats are:

`A`: return only uppercase `(ABC)`

`a`: return only lowercase `(abc)`

`d`: return only digits `(123)`

`s`: return only symbols `(&^%)`

`*`: return any char `(Aa1*)`

check [example](./example)


### FUNCTIONS

#####  func WriteFromFormat
```go
func WriteFromFormat(format string) string
```
WriteFromFormat generates a string that follow the given format.  `format` parameter is a string that says how the random string will be formated.

Example:

```go
s1:= str.WriteFromFormat("Aads") // this returns a string with a capital letter,
// small letter, digit and a string in order"Gf3("
```

##### WriteN
```go
func WriteN(n int, format ...rune) string
```
 WriteN generates a string with length n. if format is specified, the string
 will be of the given format. only the first format passed will be used.

 Example:

```go
s1:= str.WriteN(10,"A") // returns random uppercase string of len 10 "ANBCHTEDFO"
s2:= str.WriteN(10,"a") // returns random lowercase string of len 10 "gahdbyeheb"
s3:= str.WriteN(10,) // returns random string of len 10 "Ay45)-hBN*"
```
