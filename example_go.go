package example

import "github.com/codecov/example-go/awesome"

var result string

func Setup() {
    result = awesome.Smile()
}

func GetResult() string {
    return result
}
