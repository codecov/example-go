Codecov Go Example
==================

| [https://codecov.io][1] | [@codecov][2] | [hello@codecov.io][3] |
| ----------------------- | ------------- | --------------------- |

This repository serves as an **example** on how to use [Codecov Global][4] for Go.

> Note: use `-covermode=atomic` or `-covermode=count` to show how many times a statement was executed.

# Travis CI

Add to your `.travis.yml` file.
```yml
language: go

go:
  - 1.7

before_install:
  - go get -t -v ./...

script:
  - go test -race -coverprofile=coverage.txt -covermode=atomic

after_success:
  - bash <(curl -s https://codecov.io/bash)
```

> - All other CI you can simply run `bash <(curl -s https://codecov.io/bash)`.
> - `-race` is a suggestion, not required. Learn more at https://blog.golang.org/race-detector

## Private Repos
> Set `CODECOV_TOKEN` in your environment variables.

Add to your `.travis.yml` file.
```yml
after_success:
  - bash <(curl -s https://codecov.io/bash) -t uuid-repo-token
```
> Or you can set the environment variable `CODECOV_TOKEN=uuid-repo-token` and remove the `-t` flag

## Caveat: Multiple files
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


View source and learn more about [Codecov Global Uploader][4]

[1]: https://codecov.io/
[2]: https://twitter.com/codecov
[3]: mailto:hello@codecov.io
[4]: https://github.com/codecov/codecov-bash
