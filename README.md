# [Codecov](https://codecov.io) Go Example

[![codecov](https://codecov.io/gh/codecov/example-go/branch/master/graph/badge.svg)](https://codecov.io/gh/codecov/example-go)
[![Build Status](https://travis-ci.org/codecov/example-go.svg?branch=master)](https://travis-ci.org/codecov/example-go)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fcodecov%2Fexample-go.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fcodecov%2Fexample-go?ref=badge_shield)

> Note: use `-covermode=atomic` or `-covermode=count` to show how many times a statement was executed.

## Builld

### Travis Setup

Add to your `.travis.yml` file.
```yml
language: go

go:
  - 1.8.x
  - tip

before_install:
  - go get -t -v ./...

script:
  - go test ./... -race -coverprofile=coverage.txt -covermode=atomic

after_success:
  - bash <(curl -s https://codecov.io/bash)
```

> - All other CI you can simply run `bash <(curl -s https://codecov.io/bash)`.
> - `-race` is a suggestion, not required. Learn more at https://blog.golang.org/race-detector

### Makefile setup
```
GO111MODULE = on
CODECOV_TOKEN="YOUR_TOKEN_HERE"

deps:
	go mod download

format:
	go fmt ./...

test: deps format #it's recommended to run go fmt for consistency in coverage
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic

codecov: test
	CODECOV_TOKEN=$(CODECOV_TOKEN) bash -c 'bash <(curl -s https://codecov.io/bash)'

```

With this makefile, you can run `make codecov` to run your tests and have the results uploaded to codecov.


### Multiple files
> If you see this `cannot use test profile flag with multiple packages` then use this shell template to execute your tests and store coverage output

```shell
#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
```

Then run this file as your test (ex. `./test.sh`)

> Reference http://stackoverflow.com/a/21142256/2055281

## Caveats

### Private Repo
Repository tokens are required for (a) all private repos, (b) public repos not using Travis-CI, CircleCI or AppVeyor. Find your repository token at Codecov and provide via appending `-t <your upload token>` to you where you upload reports.

## Links
- [Community Boards](https://community.codecov.io)
- [Support](https://codecov.io/support)
- [Documentation](https://docs.codecov.io)


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fcodecov%2Fexample-go.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fcodecov%2Fexample-go?ref=badge_large)
