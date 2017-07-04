# spherand

Random points on a sphere in Golang.

## The Problem

If you generate latitude and longitude uniformly in the ranges `[-90, 90]` and
`[-180, 180]` you will get distributions biased towards the poles.

![Distribution of broken generator](http://i.imgur.com/zLOdR0J.png)

This package [correctly picks points on a
sphere](http://mathworld.wolfram.com/SpherePointPicking.html), providing a
uniform distribution visualized below.

![Distribution of correct generator](http://i.imgur.com/2GikHe6.png)

## Diagrams

Generated with [globe](https://github.com/mmcloughlin/globe).
