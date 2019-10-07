# Enum

Auto generate your types and implement various interfaces by only creating
constants. All you need to do is add a `go generate` line above your constant
group.

```go
//go:generate go-enum --trim "Post" --format snake
const (
    PostCreate PostType = 0
    PostRead   PostType = 2
    PostUpdate PostType = 4
    PostDelete PostType = 8
)
```

If you just want an enum and don't care about the value, use Go's `iota`.

```go
//go:generate go-enum --trim "Direction" --format upper
const (
    DirectionUp Direction = iota
    DirectionDown
    DirectionLeft
    DirectionRight
)
```

See the [example](example/) folder for example generated code.

## Installation

```
go get -u github.com/bombsimon/enum/...
```

## Interfaces

The current interfaces that will be implemented are

* `FromString` - Convert a string to the enum type
* `String` - Get the string value for the enum
* `Valid` - Returns true or false if the enum is valid
* `MarshalJSON` - JSON marshal interface
* `UnmarshalJSON` - JSON unmarshal interface

## Flags

The following flags can be used.

* `json` - Don't generate JSON interface by setting to false
* `format` - The way to format the constant. `snake|camel|upper|lower`
* `trim` - What part of the constant to trim.
