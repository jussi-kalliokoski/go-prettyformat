[![Build Status](https://travis-ci.org/jussi-kalliokoski/go-prettyformat.svg?branch=master)](https://travis-ci.org/jussi-kalliokoski/go-prettyformat)
[![GoDoc](https://godoc.org/github.com/jussi-kalliokoski/go-prettyformat?status.svg)](http://godoc.org/github.com/jussi-kalliokoski/go-prettyformat)

# go-prettyformat

A [go](https://golang.org) package for stringifying go values, similar to [pretty-format](https://www.npmjs.com/package/pretty-format) for JavaScript.

Supports pretty much all go types except arbitrary pointer values, functions, interface instances and channels.

The primary use case for this package is to support snapshot testing of data structures.

## Installation

```
go get github.com/jussi-kalliokoski/go-prettyformat
```

## Example

```go
import (
    "github.com/jussi-kalliokoski/go-prettyformat"
    "fmt"
)

func main() {
    type foo struct {
        Foo interface{}
    }

    str, err := prettyformat.Format(
        [1]foo{
            foo{
                Foo: []*foo{
                    &foo{
                        Foo: map[string]interface{}{
                            "foo": foo{
                                Foo: "123",
                            },
                        },
                    },
                },
            },
        },
    )

    if err != nil {
        panic(err)
    }

    fmt.Println(str)
    // [1]foo{
    //   foo{
    //     Foo: []*foo{
    //       *foo{
    //         Foo: map[string]interface{}{
    //           "foo": foo{
    //             Foo: (string)"123",
    //           },
    //         },
    //       },
    //     },
    //   },
    // }
}
```

## License

The MIT License (MIT) - see LICENSE file for more details.
