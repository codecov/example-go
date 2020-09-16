package example

import "github.com/codecov/example-go/awesome"

var result string

func Setup() {
    // Comment
    result = awesome.Smile()

}

func GetResult() string {
    /*
    Comment
    */
    return result
}

func UncoveredTest() int {
    x := 1
    x = x + 1
    return x
}

func CoveredTest(num int) int {
    x := num + num
    return x
}
