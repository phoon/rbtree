package rbtree

import (
	"math/rand"
	"testing"
	"time"
)

func TestRBT(tt *testing.T) {
	t := NewRBT()

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i <= 999999; i++ {
		ri := r.Intn(999999)
		t.Insert(KeyTypeInt(ri), 0)
	}

	for i := 80; i < 9999; i++ {
		t.Remove(KeyTypeInt(i))
	}

	for i := 0; i <= 80; i++ {
		t.Remove(KeyTypeInt(i))
	}

	for i := 9999; i <= 999999; i++ {
		t.Remove(KeyTypeInt(i))
	}

	if t.size != 0 {
		tt.Errorf("size is %d, not zero\n, something is wrong.", t.size)
	}
}
