package test

import (
	btree "DS_Stu/Tree"
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	tt := btree.NewBTree()
	tt.Insert(3)
	tt.Insert(4)
	tt.Insert(5)
	tt.Insert(2)
	tt.Insert(1)
	tt.Insert(9)
	tt.Insert(7)
	tt.Insert(8)

	if tt.Search(3).Val != 3 {
		t.Errorf("Expected 3 but got %v", tt.Search(3).Val)
	}
	if tt.Search(4).Val != 4 {
		t.Errorf("Expected 4 but got %v", tt.Search(4).Val)
	}
	if tt.Search(5).Val != 5 {
		t.Errorf("Expected 5 but got %v", tt.Search(5).Val)
	}
	if tt.Search(2).Val != 2 {
		t.Errorf("Expected 2 but got %v", tt.Search(2).Val)
	}
	if tt.Search(1).Val != 1 {
		t.Errorf("Expected 1 but got %v", tt.Search(1).Val)
	}
	if tt.Search(9).Val != 9 {
		t.Errorf("Expected 9 but got %v", tt.Search(9).Val)
	}
	if tt.Search(7).Val != 7 {
		t.Errorf("Expected 7 but got %v", tt.Search(7).Val)
	}
	if tt.Search(8).Val != 8 {
		t.Errorf("Expected 8 but got %v", tt.Search(8).Val)
	}
}

func TestBigData(t *testing.T) {
	tree := btree.NewBTree()
	for i := 0; i < 10000; i++ {
		tree.Insert(i)
	}

	for i := 0; i < 10000; i++ {
		a := rand.Intn(1000)
		tree.Search(a)
		if tree.Search(a).Val != a {
			t.Errorf("Expected %v but got %v", a, tree.Search(a).Val)
		}
	}
}

func TestInsertSearch(t *testing.T) {

}
