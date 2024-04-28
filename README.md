# AES in Go

## Objective

This is a Go implementation of the [Advanced Encryption Standard (AES)](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.197-upd1.pdf).

The objective of this project is to analyze the concurrent capabilities of Go, under various workloads, and to compare the performance between different languages.

## Deployment

### Requirements

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/) (optional, for running tests locally)

The project is containerized, so you don't need to have Go installed on your machine.

### Configuration

The configuration is done through the `src/.env` file (This file is built into the container). The following variables are available:

- `CORES`: Number of threads to be used in the encryption process
- `REPEAT`: Number of times the encryption/decryption process will be repeated
- `PLAIN_TEXT`: Path to the file with the data to be encrypted
- `ENCRYPTED_TEXT`: Path to the file where the encrypted data will be stored
- `DECRYPTED_TEXT`: Path to the file where the decrypted data will be stored

> Having a `PLAIN_TEXT` and `ENCRYPTED_TEXT` will mean encrypting the data, while having a `ENCRYPTED_TEXT` and `DECRYPTED_TEXT` will mean decrypting the data. Having all three will mean encrypting and decrypting the data.

### Commands

#### Setup

- `make setup`: Creates needed directories and builds the image with the binary
- `make dummy_file`: Creates a dummy with data to be encrypted

#### Run

- `make deploy`: Runs the binary and monitoring system in docker
- `make remove`: Stops and removes the containers and the image
- `make logs`: Shows the logs of the running container

#### Tests

To run the tests:

```bash
cd src
go test ./...
```

## Libraries

- [go-statsd-client](https://pkg.go.dev/github.com/cactus/go-statsd-client/v5@v5.1.0)
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv@v1.5.1)
