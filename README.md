## About

A quad-tree is a geometrical data structure which offers fast axis aligned range queries 
(i.e. bounding-box queries) on sets of rectangles and points.

goquadtree is a simple 2D implementation in golang to be used as a golang package.

Depending on size and distribution of data a query may be in orders of 
magnitudes faster than a linear search.

From the included benchmarks:

    BenchmarkRectsQuadtree               100          11309510 ns/op
    BenchmarkRectsLinear                   5         407917200 ns/op
    BenchmarkPointsQuadtree            50000             51042 ns/op
    BenchmarkPointsLinear                  5         407603000 ns/op

## Usage

The included tests and benchmarks give an example on how to use the package.

