# Go Actions

[![Test](https://github.com/ghacts/go/actions/workflows/test.yml/badge.svg)](https://github.com/ghacts/go/actions/workflows/test.yml)

<!-- action-docs-description -->

## Description

GitHub Actions for building and testing Go projects

<!-- action-docs-description -->

<!-- action-docs-inputs -->

## Inputs

| parameter              | description                            | required | default       |
| ---------------------- | -------------------------------------- | -------- | ------------- |
| project-path           | The path to the project                | `false`  |               |
| go-version             | The version of Go to use               | `false`  | stable        |
| golang-ci-lint-version | The version of golangci-lint to use    | `false`  | latest        |
| golang-ci-lint-args    | Additional arguments for golangci-lint | `false`  | --timeout=10m |
| skip-lint              | Check whether to skip lint step        | `false`  | false         |
| skip-build             | Check whether to skip build step       | `false`  | false         |
| skip-test              | Check whether to skip test step        | `false`  | false         |

<!-- action-docs-inputs -->

## Examples

```
jobs:
  check:
    name: Check
    strategy:
      matrix:
        os:
          - ubuntu-latest
          - macos-latest
          - windows-latest
        go-version:
          - '1.18'
          - '1.19'
          - '1.20'
    runs-on: ${{ matrix.os }}
    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - name: Build and test
        uses: ghacts/go@v1
        with:
          # don't need to pass this parameter if project directory is under root
          project-path: pkg/module-xyz
          go-version: ${{ matrix.go-version }}
```
