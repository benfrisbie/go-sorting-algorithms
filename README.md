[![Build Status](https://img.shields.io/github/actions/workflow/status/benfrisbie/go-sorting-algorithms/test.yml?branch=main&label=test&logo=github&style=flat-square)](https://github.com/benfrisbie/go-sorting-algorithms/actions?workflow=test)
[![Codecov](https://img.shields.io/codecov/c/github/benfrisbie/go-sorting-algorithms?logo=codecov&style=flat-square)](https://codecov.io/gh/benfrisbie/go-sorting-algorithms)

# About
Sorting algorithms written in golang using generics.

# Testing
Run tests:
```bash
go test ./...
```

Run benchmarks:
```bash
go test -bench=. ./...
```
Specify slice length
```bash
go test -bench=. -slice_length=10000 ./...
```
