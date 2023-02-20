# imaging

A little image processing library written in Go

[![Go Reference](https://pkg.go.dev/badge/github.com/go-micah/imaging.svg)](https://pkg.go.dev/github.com/go-micah/imaging) [![Go](https://github.com/go-micah/imaging/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/go-micah/imaging/actions/workflows/go.yml)

## Overview

`imaging` is a small image processing library written in Go. You can use this library in your code by importing it as a package, or you can use the CLI directly by cloning this repo and compiling the program. There is also a template for deploying the code to AWS, which allows you to use the functions in a web browser.

## Dist

`dist` generates a "distance image" which is a grayscale image where each pixel's luminance is relative the center of the image. These types of images can be useful anytime you need a quick test image as a placeholder.

![](https://imaging.micahwalter.com/dist?width=800&height=400)

## Steps

`steps` generates a "step wedge" test image. These types of test images can be useful when calibrating a camera system.

![](https://imaging.micahwalter.com/steps?width=800&height=400)


## AWS Lambda

This repository also contains an AWS SAM template as well as AWS Lambda functions, which can be used to easily deploy to AWS. The SAM template builds the following resources:

- API Gateway
- AWS Lambda Functions for each endpoint

To deploy to AWS using AWS SAM, follow these commands:

```
make aws
sam deploy --guided
```

You can also use SAM to test locally by using this command:

```
sam local start-api
```

## CLI

You can compile and use this library as a CLI application. To do so, simply run the following command:

```
make imaging-cli
```

To use the CLI, follow the help by typing:

```
./bin/imaging-cli help
```
