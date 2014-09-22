package example

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "os"
    "testing"
    "github.com/onsi/ginkgo/reporters"
)

func TestAwesome(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecsWithDefaultAndCustomReporters(t, "Awesome Suite", []Reporter{junitReporter})
}
