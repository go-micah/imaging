[![Go](https://github.com/go-micah/imaging/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/go-micah/imaging/actions/workflows/go.yml)

# imaging

A little image processing library written in Go

## Overview

`imaging` is a small image processing library written in Go. You can use this library in your code by importing it as a package, or you can use the CLI directly by cloning this repo and compiling the program. There is also a template for deploying the code to AWS, which allows you to use the functions in a web browser.

## Dist

`dist` generates a "distance image" which is a grayscale image where each pixel's luminance is relative the center of the image. These types of images can be useful anytime you need a quick test image as a placeholder.

![](https://imaging.micahwalter.com/dist?width=800&height=400)

Try this in your browser [imaging.micahwalter.com/dist?width=800&height=400](https://imaging.micahwalter.com/dist?width=800&height=400). You can adjust the width and height.

## Steps

`steps` generates a "step wedge" test image. These types of test images can be useful when calibrating a camera system.

![](https://imaging.micahwalter.com/steps?width=800&height=400)

Try this in your browser [imaging.micahwalter.com/steps?width=800&height=400](https://imaging.micahwalter.com/steps?width=800&height=400). You can adjust the width and height.

