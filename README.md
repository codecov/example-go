
# [Codecov][1] Go Example

This repository serves as an **example** on how to use [Codecov Global][4] for Go.

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
  - go test -race -coverprofile=coverage.txt -covermode=atomic

after_success:
  - bash <(curl -s https://codecov.io/bash)
```

> - All other CI you can simply run `bash <(curl -s https://codecov.io/bash)`.
> - `-race` is a suggestion, not required. Learn more at https://blog.golang.org/race-detector



## Caveats
### Private Repos
> Set `CODECOV_TOKEN` in your environment variables.

Add to your `.travis.yml` file.
```yml
after_success:
  - bash <(curl -s https://codecov.io/bash) -t uuid-repo-token
```
> Or you can set the environment variable `CODECOV_TOKEN=uuid-repo-token` and remove the `-t` flag
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


## Support
### Contact
- Intercom (in-app messanger)
- Email: support@codecov.io
- Slack: slack.codecov.io
- [gh/codecov/support](https://github.com/codecov/support)

1. More documentation at https://docs.codecov.io
2. Configure codecov through the `codecov.yml`  https://docs.codecov.io/docs/codecov-yaml
3. View source and learn more about [Codecov Global Uploader][4]


[1]: https://codecov.io/
[4]: https://github.com/codecov/codecov-bash
