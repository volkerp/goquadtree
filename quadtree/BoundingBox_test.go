package quadtree

import "testing"
import _ "math"

func TestBoundingBox(t *testing.T) {
  a := NewBoundingBox( 0, 10, 0, 10 )
  
  b := NewBoundingBox( 4, 6, 4, 6 )    // b completely within a 
  
  if ! a.Intersects(b) || ! b.Intersects(a) {
    t.Errorf("%v does not intersect %v", a, b) 
  }
  
  if ! a.Intersects(a) {
    t.Errorf("%v does not intersect itself", a) 
  }
  
  if ! a.Contains(b) {
    t.Errorf("%v does not contain %v", a, b)
  }

  if ! a.Contains(a) {
    t.Errorf("%v does not contain itself", a)
  }
  
  if b.Contains(a) {
    t.Errorf("%v contains %v", b, a)
  }

  c := NewBoundingBox( 10, 20, 0, 10 )
  
  if a.Intersects(c) {
    t.Errorf("%v does intersect %v", a, c) 
  }
  if c.Intersects(a) {
    t.Errorf("%v does intersect %v", c, a) 
  }

  if a.Contains(c) || c.Contains(a) {
    t.Errorf("%v contains %v (or vise versa)", a, c) 
  }
  
  d := NewBoundingBox( -10, 0, 0, 10 )
  
  if a.Intersects(d) {
    t.Errorf("%v does intersect %v", a, d) 
  }
  if d.Intersects(a) {
    t.Errorf("%v does intersect %v", d, a) 
  }
  
  e := NewBoundingBox( 9, 15, 9, 15 )
   
  if ! a.Intersects(e) || ! e.Intersects(a) {
    t.Errorf("%v does not intersect %v", a, e) 
  }
  
  f := NewBoundingBox( -10, 20, 4, 6 )
  
  if  ! a.Intersects(f) || ! f.Intersects(a) {
    t.Errorf("%v does not intersect %v", a, f) 
  }
}


