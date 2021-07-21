go-splittestgen
=======
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

go-splittestgen splits test cases into some subsets and print commands to run one of subsets for parallel testing.
* The parser code is based on [github.com/Songmu/gotesplit](https://github.com/Songmu/gotesplit) written by [Songmu](https://github.com/Songmu)

## Usage

`go-splittestgen` is parsing output of `go test -list` from stdin(in the usual case, it is piped),
generating test commands from the test cases divided by a number,
and then printing one of the commands by an index number.

if you want to run the generated test commands immediately, just pipe it to a shell.

```bash
# print test commands
$ go test ./... -list . | go-splittestgen -total 3 -index 2

# execute tests
$ go test ./... -list . | go-splittestgen -total 3 -index 2 | sh
```

### Options

```
-total uint
        total number of test processes (default 1)
-index uint
        index of test processes (default 0, must be less than the total number)
```

## Installation

```bash
# go install
$ go install github.com/minoritea/go-splittestgen/cmd/go-splittestgen
# or just run
$ go test ./... -list . | go run github.com/minoritea/go-splittestgen/cmd/go-splittestgen
```

## Example
### GitHub Actions

```yaml
name: CI
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    strategy:
      fail-fast: false
      matrix:
        parallelism: [3]
        index: [0,1,2]
    steps:
      - uses: actions/setup-go@v2
      - uses: actions/checkout@v2
      # Add github.com/minoritea/go-splittestgen to go.mod
      # and install modules before tests.
      - name: Run tests parallelly
        run: |
          go mod download
          go test ./... -list . | \
          go run github.com/minoritea/go-splittestgen/cmd/go-splittestgen \
            -total ${{ matrix.parallelism }} \
            -index ${{ matrix.index }} | \
          sed -e 's/$/ -v -count 1/g' | sh
```

## Author
[minoritea](https://github.com/minoritea)

## Original Author
[Songmu](https://github.com/Songmu)
