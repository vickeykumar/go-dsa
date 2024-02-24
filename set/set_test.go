package set

import (
	"fmt"
	"reflect"
	"testing"
)

// slicesEqualIgnoringOrder checks if two slices are equal ignoring the order of elements
func slicesEqualIgnoringOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	seen := make(map[int]bool)
	for _, v := range a {
		seen[v] = true
	}
	for _, v := range b {
		if !seen[v] {
			return false
		}
	}
	return true
}

func TestNewSet(t *testing.T) {
	// Test NewSet with an integer slice containing repeated elements
	s := NewSet([]int{1, 2, 3, 2, 4, 4})
	expectedSlice := []int{1, 2, 3, 4}
	fmt.Printf("NewSet with repeated elements, expected: %v, got: %v", expectedSlice, s.ToSlice())
	if !slicesEqualIgnoringOrder(s.ToSlice().([]int), expectedSlice) {
		t.Errorf("NewSet with repeated elements failed, expected: %v, got: %v", expectedSlice, s.ToSlice())
	}

	// Test NewSet with an empty slice
	s = NewSet([]int{})
	if s.Len() != 0 {
		t.Errorf("NewSet with empty slice failed, expected length: 0, got: %d", s.Len())
	}
}

func TestSet_Len(t *testing.T) {
	s := NewSet([]int{1, 2, 3})
	expectedLen := 3
	if s.Len() != expectedLen {
		t.Errorf("Expected length: %d, got: %d", expectedLen, s.Len())
	}
}

func TestSet_Append(t *testing.T) {
	s := NewSet([]int{1, 2, 3})
	s.Append([]int{4, 5, 6})
	expectedLen := 6
	if s.Len() != expectedLen {
		t.Errorf("Expected length after Append: %d, got: %d", expectedLen, s.Len())
	}
}

func TestSet_GetType(t *testing.T) {
	s := NewSet([]int{1, 2, 3})
	expectedType := reflect.TypeOf([]int{})
	if s.GetType() != expectedType {
		t.Errorf("Expected type: %v, got: %v", expectedType, s.GetType())
	}
}

func TestSet_Update(t *testing.T) {
	s1 := NewSet([]int{1, 2, 3})
	s2 := NewSet([]int{4, 5, 6})
	s1.Update(s2)
	expectedLen := 6
	if s1.Len() != expectedLen {
		t.Errorf("Expected length after Update: %d, got: %d", expectedLen, s1.Len())
	}
}

func TestSet_Add(t *testing.T) {
	s := NewSet([]int{1, 2, 3})
	s.Add(4, 5, 6)
	expectedLen := 6
	if s.Len() != expectedLen {
		t.Errorf("Expected length after Add: %d, got: %d", expectedLen, s.Len())
	}
}

func TestSet_ToSlice(t *testing.T) {
	s := NewSet([]int{1, 2, 3})
	expectedSlice := []int{1, 2, 3}
	if !slicesEqualIgnoringOrder(s.ToSlice().([]int), expectedSlice) {
		t.Errorf("Expected slice: %v, got: %v", expectedSlice, s.ToSlice())
	}
}

func TestSet_Remove(t *testing.T) {
	s := NewSet([]int{1, 2, 3})
	err := s.Remove(2, 3)
	if err != nil {
		t.Errorf("Error removing elements: %v", err)
	}
	expectedLen := 1
	if s.Len() != expectedLen {
		t.Errorf("Expected length after Remove: %d, got: %d", expectedLen, s.Len())
	}
}

func TestSet_Clear(t *testing.T) {
	s := NewSet([]int{1, 2, 3})
	s.Clear()
	expectedLen := 0
	if s.Len() != expectedLen {
		t.Errorf("Expected length after Clear: %d, got: %d", expectedLen, s.Len())
	}
}

func TestCopy(t *testing.T) {
	s1 := NewSet([]int{1, 2, 3})
	s2 := Copy(s1)
	if !s1.Eq(s2) {
		t.Errorf("Copy failed, sets are not equal")
	}
}

func TestUnion(t *testing.T) {
	s1 := NewSet([]int{1, 2, 3})
	s2 := NewSet([]int{3, 4, 5})
	unionSet := Union(s1, s2)
	expectedSlice := []int{1, 2, 3, 4, 5}
	if !slicesEqualIgnoringOrder(unionSet.ToSlice().([]int), expectedSlice) {
		t.Errorf("Union failed, expected: %v, got: %v", expectedSlice, unionSet.ToSlice())
	}
}

func TestIntersection(t *testing.T) {
	s1 := NewSet([]int{1, 2, 3})
	s2 := NewSet([]int{3, 4, 5})
	intersectionSet := Intersection(s1, s2)
	expectedSlice := []int{3}
	if !slicesEqualIgnoringOrder(intersectionSet.ToSlice().([]int), expectedSlice) {
		t.Errorf("Intersection failed, expected: %v, got: %v", expectedSlice, intersectionSet.ToSlice())
	}
}

func TestDifference(t *testing.T) {
	s1 := NewSet([]int{1, 2, 3})
	s2 := NewSet([]int{3, 4, 5})
	differenceSet := Difference(s1, s2)
	expectedSlice := []int{1, 2}
	if !slicesEqualIgnoringOrder(differenceSet.ToSlice().([]int), expectedSlice) {
		t.Errorf("Difference failed, expected: %v, got: %v", expectedSlice, differenceSet.ToSlice())
	}
}

func TestSet_Eq(t *testing.T) {
	s1 := NewSet([]int{1, 2, 3})
	s2 := NewSet([]int{1, 2, 3})
	if !s1.Eq(s2) {
		t.Errorf("Eq failed, sets are not equal")
	}
}

func TestSet_Isdisjoint(t *testing.T) {
	s1 := NewSet([]int{1, 2, 3})
	s2 := NewSet([]int{4, 5, 6})
	if !s1.Isdisjoint(s2) {
		t.Errorf("Isdisjoint failed, sets are not disjoint")
	}
}

func TestSet_Issubset(t *testing.T) {
	s1 := NewSet([]int{1, 2, 3})
	s2 := NewSet([]int{1, 2, 3, 4, 5})
	if !s1.Issubset(s2) {
		t.Errorf("Issubset failed, s1 is not a subset of s2")
	}
}

func TestSet_Issuperset(t *testing.T) {
	s1 := NewSet([]int{1, 2, 3, 4, 5})
	s2 := NewSet([]int{1, 2, 3})
	if !s1.Issuperset(s2) {
		t.Errorf("Issuperset failed, s1 is not a superset of s2")
	}
}

func TestSet_Contains(t *testing.T) {
	s := NewSet([]int{1, 2, 3})
	if !s.Contains(1) {
		t.Errorf("Contains failed, set should contain the element")
	}
	if s.Contains(4) {
		t.Errorf("Contains failed, set should not contain the element")
	}
}
