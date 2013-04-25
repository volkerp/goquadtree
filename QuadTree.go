package quadtree

import _ "fmt"

// Number of entries until a quad is split
const MAX_ENTRIES_PER_TILE = 16

// Maximum depth of quad-tree (not counting the root node)
const MAX_LEVELS = 10

// some constants for tile-indeces, for clarity
const ( 
  _TOPRIGHT    = 0
  _TOPLEFT     = 1
  _BOTTOMLEFT  = 2
  _BOTTOMRIGHT = 3
)


// QuadTree exspects its values to implement the BoundingBoxer interface.
type BoundingBoxer interface {
  BoundingBox() BoundingBox
}


type QuadTree struct {
  root qtile
}


// Constructs an empty quad-tree
// bbox specifies the extends of the coordinate system.
func NewQuadTree(bbox BoundingBox) QuadTree {
  qt := QuadTree{ qtile{BoundingBox:bbox} }

  return qt
}


// quad-tile / node of the quad-tree
type qtile struct {
           BoundingBox
  level    int                // level this tile is at. root is level 0
  contents []BoundingBoxer    // values stored in this tile
  childs   [4]*qtile          // sub-tiles. none or four.
}


// Adds a value to the quad-tree by trickle down from the root node.
func (qb *QuadTree) Add(v BoundingBoxer) {
  qb.root.add(v)
}


// Returns all values which intersect the query box
func (qb *QuadTree) Query(bbox BoundingBox) (values []BoundingBoxer) {
  return qb.root.query(bbox, values)
}


func (tile *qtile) add(v BoundingBoxer) {
  // look for sub-tile directly below this tile to accomodate value.
  if i := tile.findChildIndex(v.BoundingBox()); i < 0 {
    // no suitable sub-tile for value found.
    // either this tile has no childs or
    // value does not fit in any subtile.
    // store value at this level.
    tile.contents = append(tile.contents, v)

    // tile is split if exceeds it max number of entries and
    // has not childs already and max tree depth for this sub-tree not reached.
    if len(tile.contents) > MAX_ENTRIES_PER_TILE && tile.childs[_TOPRIGHT] == nil && tile.level < MAX_LEVELS {
      tile.split()
    }
  } else {
    // suitable sub-tile for value found at index i.
    // recursivly add value.
    tile.childs[i].add(v)
  }
}


// return child index for BoundingBox
// returns -1 if quad has no children or BoundingBox does not fit into any child
func (tile *qtile) findChildIndex(bbox BoundingBox) int {
  if tile.childs[_TOPRIGHT] == nil {
    return -1
  }

  for i, child := range tile.childs {
    if child.Contains(bbox) {
      return i
    }
  }

  return -1
}


// create four child quads.
// distribute contents of this tiles on newly created childs.
func (tile *qtile) split() {
  mx := tile.MaxX/2.0 + tile.MinX/2.0
  my := tile.MaxY/2.0 + tile.MinY/2.0

  tile.childs[_TOPRIGHT]    = &qtile{ BoundingBox:NewBoundingBox(mx, tile.MaxX, my, tile.MaxY), level:tile.level+1 }
  tile.childs[_TOPLEFT]     = &qtile{ BoundingBox:NewBoundingBox(tile.MinX, mx, my, tile.MaxY), level:tile.level+1 }
  tile.childs[_BOTTOMLEFT]  = &qtile{ BoundingBox:NewBoundingBox(tile.MinX, mx, tile.MinY, my), level:tile.level+1 }
  tile.childs[_BOTTOMRIGHT] = &qtile{ BoundingBox:NewBoundingBox(mx, tile.MaxX, tile.MinY, my), level:tile.level+1 }

  // copy values to temporary slice
  var contentsBak []BoundingBoxer
  contentsBak = append(contentsBak, tile.contents...)

  // clear values on this tile
  tile.contents = []BoundingBoxer{}

  // reinsert from temporary slice
  for _,v := range contentsBak {
    tile.add(v)
  }
}


func (tile *qtile) query(qbox BoundingBox, ret []BoundingBoxer) []BoundingBoxer {
  // end recursion if this tile does not intersect the query range
  if ! tile.Intersects(qbox) {
    return ret
  }
  
  // return candidates at this tile
  for _, v := range tile.contents {
    if qbox.Intersects(v.BoundingBox()) {
      ret = append(ret, v)
    }
  }
  
  // recurse into childs (if any)
  if tile.childs[_TOPRIGHT] != nil {
    for _, child := range tile.childs {
      ret = child.query(qbox, ret)
    }
  } 

  return ret
}



