# imaging

A little image processing library written in Go

[![Go Reference](https://pkg.go.dev/badge/github.com/go-micah/imaging.svg)](https://pkg.go.dev/github.com/go-micah/imaging) [![Go](https://github.com/go-micah/imaging/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/go-micah/imaging/actions/workflows/go.yml)

## Dist

![](https://imaging.micahwalter.com/dist?width=800&height=400)

## Steps

![](https://imaging.micahwalter.com/steps?width=800&height=400)


## AWS Lambda

This repository also contains an AWS SAM template as well as AWS Lambda functions, which can be used to easily deployed to AWS. The SAM template builds the following resources:

- API Gateway
- AWS Lambda Functions for each endpoint

## CLI

You can compile and use this library as a CLI application. To do so, simply run the following command:

```
make imaging-cli
```

To use the CLI, follow the help by typing:

```
./bin/imaging-cli help
```
