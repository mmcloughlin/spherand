# spherand

Random points on a sphere in Golang.

[![GoDoc Reference](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](http://godoc.org/github.com/mmcloughlin/spherand)
[![Build status](https://img.shields.io/travis/mmcloughlin/spherand.svg?style=flat-square)](https://travis-ci.org/mmcloughlin/spherand)


## The Problem

If you generate latitude and longitude uniformly in the ranges `[-90, 90]` and
`[-180, 180]` you will get distributions biased towards the poles.

![Distribution of broken generator](http://i.imgur.com/zLOdR0J.png)

This package [correctly picks points on a
sphere](http://mathworld.wolfram.com/SpherePointPicking.html), providing a
uniform distribution visualized below.

![Distribution of correct generator](http://i.imgur.com/2GikHe6.png)

## Getting Started

Install with

```sh
$ go get -u github.com/mmcloughlin/spherand
```

Generate a geographical point with

```go
lat, lng := spherand.Geographical()
```

You can also generate spherical coordinates with `Spherical()`.

If you need to control the random source, use a generator

```go
g := spherand.NewGenerator(rand.New(rand.NewSource(42)))
lat, lng := g.Geographical()
```

See [godoc](https://godoc.org/github.com/mmcloughlin/spherand) for reference.

## Diagrams

Generated with [globe](https://github.com/mmcloughlin/globe).
