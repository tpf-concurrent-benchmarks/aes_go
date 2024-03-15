# AES in Go

This is a Go implementation of the [Advanced Encryption Standard (AES)](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.197-upd1.pdf).

The objective of this project is to analyze the concurrent capabilities of Go, under various workloads, and to compare the performance between different languages.

## Requirements

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/) (optional, for running tests locally)

The project is containerized, so you don't need to have Go installed on your machine.

## Usage

### Setup

Creates needed directories and builds the image with the binary:

```bash
make setup
```

### Run

Runs the binary and monitoring system in docker:

```bash
make deploy
```

### Cleanup

Stops and removes the containers and the image:

```bash
make remove
```

### Tests

Runs the tests (also in a container)

```bash
cd src
go test ./...
```
