package quadtree

import "testing"
import "math"
import "math/rand"
import _ "fmt"



// Generates n BoundingBoxes in the range of frame with average width and height avgSize
func randomBoundingBoxes(n int, frame BoundingBox, avgSize float64) []BoundingBox {
  ret := make([]BoundingBox, n)
  
  for i := 0; i < len(ret); i++ {
    w := rand.NormFloat64() * avgSize
    h := rand.NormFloat64() * avgSize
    x := rand.Float64() * frame.SizeX() + frame.MinX
    y := rand.Float64() * frame.SizeY() + frame.MinY
    ret[i] = NewBoundingBox(x, math.Min(frame.MaxX, x+w), y, math.Min(frame.MaxY, y+h))
  }
  
  return ret
}
   

// Returns all elements of data which intersect query
func queryBruteForce(data []BoundingBox, query BoundingBox) (ret []BoundingBoxer) {
  for _, v := range data {
    if query.Intersects(v.BoundingBox()) {
      ret = append(ret, v)
    }
  }
  
  return ret
}


func compareBoundingBoxer(v1, v2 BoundingBoxer) bool {
  b1 := v1.BoundingBox()
  b2 := v2.BoundingBox()
  
  return b1.MinX == b2.MinX && b1.MaxX == b2.MaxX &&
         b1.MinY == b2.MinY && b2.MaxY == b2.MaxY
}


func lookupResults(r1, r2 []BoundingBoxer) int {
  for i, v1 := range r1 {
    found := false
    
    for _, v2 := range r2 {
      if compareBoundingBoxer(v1, v2) {
        found = true
        break
      }
    }
  
    if ! found {
      return i
    }
  }
  
  return -1
}
  
// World-space extends from -1000..1000 in X and Y direction
var world BoundingBox = NewBoundingBox(-1000.0, 1000.0, -1000.0, 1000.0)
  

// Compary correctness of quad-tree results vs simple look-up on set of random rectangles
func TestQuadTreeRects(t *testing.T) {
  var rects []BoundingBox = randomBoundingBoxes(100*1000, world, 5)        
  qt := NewQuadTree(world)
  
  for _, v := range rects {
    qt.Add(v)
  }

  queries := randomBoundingBoxes(1000, world, 10)
  
  for _, q := range queries {
    r1 := queryBruteForce(rects, q)
    r2 := qt.Query(q)
    
    if len(r1) != len(r2) {
      t.Errorf("r1 and r2 differ: %v   %v\n", r1, r2)
    }
  
    if i := lookupResults(r1, r2); i != -1 {
      t.Errorf("%v was not returned by QT\n", r1[i])
    }
    
    if i := lookupResults(r2, r1); i != -1 {
      t.Errorf("%v was not returned by brute-force\n", r2[i])
    }

  }
}


// Compary correctness of quad-tree results vs simple look-up on set of random points
func TestQuadTreePoints(t *testing.T) {
  var points []BoundingBox = randomBoundingBoxes(100*1000, world, 0)        
  qt := NewQuadTree(world)
  
  for _, v := range points {
    qt.Add(v)
  }

  queries := randomBoundingBoxes(1000, world, 10)
  
  for _, q := range queries {
    r1 := queryBruteForce(points, q)
    r2 := qt.Query(q)
    
    if len(r1) != len(r2) {
      t.Errorf("r1 and r2 differ: %v   %v\n", r1, r2)
    }
  
    if i := lookupResults(r1, r2); i != -1 {
      t.Errorf("%v was not returned by QT\n", r1[i])
    }
    
    if i := lookupResults(r2, r1); i != -1 {
      t.Errorf("%v was not returned by brute-force\n", r2[i])
    }

  }
}


// Benchmark insertion into quad-tree
func BenchmarkInsert(b *testing.B) {
  b.StopTimer()

  var values []BoundingBox = randomBoundingBoxes(b.N, world, 5)
  qt := NewQuadTree(world)
  
  b.StartTimer()
  
  for _, v := range values {
    qt.Add(v)
  }
}


// A set of 10 million randomly distributed rectangles of avg size 5
var boxes10M []BoundingBox = randomBoundingBoxes(10*1000*1000, world, 5)        

// Benchmark quad-tree on set of rectangles
func BenchmarkRectsQuery(b *testing.B) {
  b.StopTimer()
  rand.Seed(1)
  qt := NewQuadTree(world)
 
  for _, v := range boxes10M {
    qt.Add(v)
  }

  queries := randomBoundingBoxes(b.N, world, 10)
  
  b.StartTimer()
  for _, q := range queries {
    qt.Query(q)
  }
}


// Benchmark simple look up on set of rectangles
func BenchmarkRectsBruteForce(b *testing.B) {
  b.StopTimer()
  rand.Seed(1)
  queries := randomBoundingBoxes(b.N, world, 10)

  b.StartTimer()
  for _, q := range queries {
    queryBruteForce(boxes10M, q)
  }
}
   
   
// A set of 10 million randomly distributed points
var points10M []BoundingBox = randomBoundingBoxes(10*1000*1000, world, 0)        

// Benchmark quad-tree on set of points
func BenchmarkPointsQuery(b *testing.B) {
  b.StopTimer()
  rand.Seed(1)
  qt := NewQuadTree(world)
 
  for _, v := range points10M {
    qt.Add(v)
  }

  queries := randomBoundingBoxes(b.N, world, 10)
  
  b.StartTimer()
  for _, q := range queries {
    qt.Query(q)
  }
}


// Benchmark simple look up on set of points
func BenchmarkPointsBruteForce(b *testing.B) {
  b.StopTimer()
  rand.Seed(1)
  queries := randomBoundingBoxes(b.N, world, 10)

  b.StartTimer()
  for _, q := range queries {
    queryBruteForce(points10M, q)
  }
}
