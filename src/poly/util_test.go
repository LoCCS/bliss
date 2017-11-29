package poly

import (
  "testing"
)

func TestAddMod(t *testing.T) {
  tests := []int32{0,0,5,0, 0,1,2,1, 1,0,5,1, 3,4,7,0, 4,4,7,1}
  for i := 0; 4 * i < len(tests); i++ {
    a,b,q,r := tests[4*i],tests[4*i+1],tests[4*i+2],tests[4*i+3]
    if c := addMod(a,b,uint32(q)); c != r {
      t.Errorf("Error in computing addMod(%d,%d,%d)! Expected %d, got %d",
        a, b, q, r, c)
    }
  }
}

func TestSubMod(t *testing.T) {
  tests := []int32{0,1,5,4, 0,1,2,1, 1,0,5,1, 3,4,7,6, 4,4,7,0}
  for i := 0; 4 * i < len(tests); i++ {
    a,b,q,r := tests[4*i],tests[4*i+1],tests[4*i+2],tests[4*i+3]
    if c := subMod(a,b,uint32(q)); c != r {
      t.Errorf("Error in computing subMod(%d,%d,%d)! Expected %d, got %d",
        a, b, q, r, c)
    }
  }
}

func TestMulMod(t *testing.T) {
  tests := []int32{0,1,5,0, -1,1,2,1, 3,3,5,4, -3,-3,7,2, -3,2,7,1}
  for i := 0; 4 * i < len(tests); i++ {
    a,b,q,r := tests[4*i],tests[4*i+1],tests[4*i+2],tests[4*i+3]
    if c := mulMod(a,b,uint32(q)); c != r {
      t.Errorf("Error in computing mulMod(%d,%d,%d)! Expected %d, got %d",
        a, b, q, r, c)
    }
  }
}

func TestExpMod(t *testing.T) {
  tests := []int32{2,3,5,3, 3,6,7,1, 3,0,7,1, 3,5,7,5, 3,7,7,3}
  for i := 0; 4 * i < len(tests); i++ {
    a,b,q,r := tests[4*i],tests[4*i+1],tests[4*i+2],tests[4*i+3]
    if c := expMod(a,uint32(b),uint32(q)); c != r {
      t.Errorf("Error in computing expMod(%d,%d,%d)! Expected %d, got %d",
        a, b, q, r, c)
    }
  }
}
