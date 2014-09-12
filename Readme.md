# go-gravatar

[![Build Status](https://travis-ci.org/nowk/go-gravatar.svg?branch=master)](https://travis-ci.org/nowk/go-gravatar)

Go Gravtar

## Example

    import "github.com/nowk/go-gravatar"

    grav := gravatar.Lookup("email@example.com")

    grav.URL.String() // => http://www.gravatar.com/avatar/xxxxXXXXxxxx

## License

MIT

