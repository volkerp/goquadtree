## About

A quad-tree is a geometrical data structure which offers fast axis aligned range queries 
(i.e. bounding-box queries) on sets of rectangles and points.

quadtree is a simple implementation in golang to be used as a golang package.

## Usage

Depending on size and distribution of data a bounding-box query may be in order of 
magnitudes faster than a linear search.

From the included benchmarks:

    BenchmarkRectsQuery                  100          11309510 ns/op
    BenchmarkRectsBruteForce               5         407917200 ns/op
    BenchmarkPointsQuery               50000             51042 ns/op
    BenchmarkPointsBruteForce              5         407603000 ns/op


