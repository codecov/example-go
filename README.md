Go Example
=======
| [https://codecov.io/][1] | [@codecov][2] | [hello@codecov.io][3] |
| ------------------------ | ------------- | --------------------- |

> Example of how to integrate with [Codecov.io][1] for your **awesome** Go project!

## See this repos [Coveage Reports][4]


## Usage

```sh
sudo pip install codecov
codecov --token=<repo token>
```

## Require min coverage
```sh
codecov --min-coverage=75
```
> if coverage is under `75` codecov will trigger your build to fail

# [![travis-org](https://avatars2.githubusercontent.com/u/639823?v=2&s=50)](https://travis-ci.org) Travis C
> Append to your `.travis.yml`

```yml
before_install:
    sudo pip install codecov
after_success:
    codecov
```

> ### Start testing with [Travis](https://travis-ci.org/)

# [![codeship](https://avatars1.githubusercontent.com/u/2988541?v=2&s=50)](https://codeship.io/) Codeship
> Append to your `Test Commands` *after* your test commands

```sh
sudo pip install codecov
codecov --token=<repo token>
```

> ### Start testing with [Codeship](https://codeship.io/)


# [![circleci](https://avatars0.githubusercontent.com/u/1231870?v=2&s=50)](https://circleci.com/) Circle CI
> Append to your `circle.yml` file

```yml
test:
    post:
        - sudo pip install codecov
        - codecov --token=<repo token>
```
> ### Start testing with [Circle CI](https://circleci.com/)


### `(optional)` using [gocov](https://github.com/axw/gocov)
> We support uploading xml reports too, via this method below

```
install:
  - sudo pip install codecov
  - go get github.com/axw/gocov/gocov/...
  - go get github.com/AlekSi/gocov-xml
script:
  - gocov test github.com/your/project | gocov-xml > coverage.xml
after_success:
  - codecov
```


[1]: https://codecov.io/
[2]: https://twitter.com/codecov
[3]: mailto:hello@codecov.io
[4]: https://codecov.io/github/codecov/example-go
