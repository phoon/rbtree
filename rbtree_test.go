package rbtree

import (
	"math/rand"
	"testing"
	"time"
)

// 100%
func TestCoverage(tt *testing.T) {
	t := NewRBT()
	arr1 := []int{1, 2, 3}
	for _, v := range arr1 {
		t.Insert(KeyTypeInt(v), 0)
	}
	t.Search(KeyTypeInt(1))
	t.Search(KeyTypeInt(0))
	for _, v := range arr1 {
		t.Remove(KeyTypeInt(v))
	}
	for i := len(arr1) - 1; i <= 0; i-- {
		t.Insert(KeyTypeInt(arr1[i]), 0)
	}
	for _, v := range arr1 {
		t.Remove(KeyTypeInt(v))
	}

	arr2 := []int{
		97, 46, 80, 28, 73, 21, 26, 34, 53, 64,
		75, 38, 94, 82, 11, 15, 85, 45, 2, 36,
		7, 58, 77, 14, 30, 74, 25, 17, 76, 91,
		68, 55, 95, 92, 3, 5, 50, 13, 32, 20,
		4, 72, 49, 8, 78, 62, 39, 93, 10,
		67, 54, 70, 90, 59, 40, 27, 66, 22,
		69, 18, 9, 12, 35, 99, 19, 87, 1, 24,
		23, 16, 48, 51, 86, 88, 84, 89, 98, 63,
		52, 81, 41, 37, 60, 79, 6, 42, 61, 47,
		57, 33, 56, 71, 100, 83, 96, 65, 29, 31,
		57, 33, // this 2 elements are duplicate values
	}

	for _, v := range arr2 {
		t.Insert(KeyTypeInt(v), 0)
	}

	if min(t.root).Key != KeyTypeInt(1) {
		tt.Error("error in min")
	}

	if max(t.root).Key != KeyTypeInt(100) {
		tt.Error("error in max")
	}

	if t.predecessor(t.search(t.root, KeyTypeInt(1)).Left) != nil {
		tt.Error("error in predecessor")
	}

	if t.predecessor(t.search(t.root, KeyTypeInt(22))).Key != KeyTypeInt(21) {
		tt.Error("error in predecessor")
	}

	if t.successor(t.search(t.root, KeyTypeInt(1)).Left) != nil {
		tt.Error("error in successor")
	}

	if t.successor(t.search(t.root, KeyTypeInt(12))).Key != KeyTypeInt(13) {
		tt.Error("error in successor")
	}

	if t.successor(t.search(t.root, KeyTypeInt(26))).Key != KeyTypeInt(27) {
		tt.Error("error in successor")
	}

	for _, v := range arr2 {
		t.Remove(KeyTypeInt(v))
	}
	if t.Size() != 0 {
		tt.Errorf("size is %d, not zero\n, something is wrong.", t.Size())
	}
}

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

	if t.Size() != 0 {
		tt.Errorf("size is %d, not zero\n, something is wrong.", t.Size())
	}
}
