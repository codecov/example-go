Codecov Go Example
==================

[![Join the chat at https://gitter.im/codecov/example-go](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/codecov/example-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

| [https://codecov.io][1] | [@codecov][2] | [hello@codecov.io][3] |
| ----------------------- | ------------- | --------------------- |

This repository serves as an **example** on how to use [Codecov Global][4] for Go.

> Note: use `-covermode=atomic` or `-covermode=count` to show how many times a statement was exected.

# Travis CI

Add to your `.travis.yml` file.
```yml
language: go

go:
  - 1.2

before_install:
  - go get github.com/onsi/gomega
  - go get github.com/onsi/ginkgo
  - go get code.google.com/p/go.tools/cmd/cover

script:
  - go test -coverprofile=coverage.txt -covermode=atomic

after_success:
  - bash <(curl -s https://codecov.io/bash)
```

> All other CI you can simply run `bash <(curl -s https://codecov.io/bash)`.

## Private Repos
> Set `CODECOV_TOKEN` in your environment variables.

Add to your `.travis.yml` file.
```yml
env:
  global:
    - CODECOV_TOKEN=:uuid-repo-token

after_success:
  - bash <(curl -s https://codecov.io/bash)
```

## Caveat: Multiple files 
> If you see this `cannot use test profile flag with multiple packages` then use this shell template to execute your tests and store coverage output

```shell
#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(find ./* -maxdepth 10 -type d); do
    if ls $d/*.go &> /dev/null; then
        go test  -coverprofile=profile.out -covermode=atomic $d
        if [ -f profile.out ]; then
            cat profile.out >> coverage.txt
            echo '<<<<<< EOF' >> coverage.txt
            rm profile.out
        fi
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
