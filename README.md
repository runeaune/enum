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

See the [example](example/) folder for example generated code. There you'll also
find information about how to create string enums and accept both their string
representation and the actual value.

## Installation

```
go get -u github.com/bombsimon/enum/...
```

## Interfaces

The current interfaces that will be implemented are

* `FromString` - Convert a string to the enum type
* `String` - Get the string value for the enum
* `Value` - Get the actual value of the enum
* `Valid` - Returns true or false if the enum is valid
* `MarshalJSON` - JSON marshal interface
* `UnmarshalJSON` - JSON unmarshal interface

## Flags

The following flags can be used.

* `json` - Don't generate JSON interface by setting to false
* `value` - Set to true to generate `Value()` and allow `FromString()` to accept
  the enums actual value as well
* `format` - The way to format the constant.
* `trim` - What part of the constant to trim.

## Format functions

The string representation of the enum can beformatted in multiple ways. The
value will be the name of the constant mins the part that's been trimmed off by
setting `--trim`. Below are examples assuming `--trim` is set to `Prefix`.

| Method                | Flag name        | Constant                | `String()` value            |
| --------------------- | ---------------- | ----------------------- | --------------------------- |
| Snake case            | snake            | `PrefixThisString`      | `this_string`               |
| Camel case            | camel            | `PrefixSomeValue`       | `SomeValue`                 |
| Upper                 | upper            | `PrefixMyValue`         | `MYVALUE`                   |
| Lower                 | lower            | `PrefixMoreConstants`   | `moreconstants`             |
| First letter          | first            | `PrefixJustFirstLetter` | `J` (will preserve casing)  |
| First letter upper    | first-upper      | `PrefixSomeValue`       | `S` (will convert to upper) |
| First letter lower    | first-lower      | `PrefixANotherValue`    | `a` (will convert to lower) |
| Capitalize first word | capitalize-first | `PrefixThisIsWords`     | `This is words`             |
| Capitalzie all words  | capitalize-all   | `PrefixThisIsAlsoWords` | `This Is Also Words`        |
