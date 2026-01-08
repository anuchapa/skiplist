package skiplist

import (
	"testing"
)

func TestInsertAndFind(t *testing.T) {

	sl := NewSkipList[int, string]()
	dataInsert := [4]struct {
		k int
		v string
	}{
		{1, "One"},
		{2, "Two"},
		{3, "Three"},
		{4, "Four"},
	}

	for _, d := range dataInsert {
		sl.Insert(d.k, d.v)
	}

	for _, d := range dataInsert {
		result := sl.Find(d.k)
		if result == nil {
			t.Fatalf("Expected to find key %d, but got nil", d.k)
		}

		if result.Value != d.v {
			t.Errorf("Expected value %s, but got %s", d.v, result.Value)
		}
	}
}

func TestUpdate(t *testing.T) {
	sl := NewSkipList[int, string]()
	key := 10
	value := "Hello Skip List"
	update := "Update Skip List"

	sl.Insert(key, value)
	result := sl.Find(key)

	result.Value = update
	if result.Value != update {
		t.Errorf("Expected update value %s, but got %s", update, result.Value)
	}
}

func TestRemove(t *testing.T) {
	sl := NewSkipList[int, string]()

	dataInsert := [4]struct {
		k int
		v string
	}{
		{1, "One"},
		{2, "Two"},
		{3, "Three"},
		{4, "Four"},
	}

	dataRemove := []int{1, 2, 5}

	for _, d := range dataInsert {
		sl.Insert(d.k, d.v)
	}

	for _, key := range dataRemove {
		sl.Remove(key)
		result := sl.Find(key)
		if result != nil {
			t.Fatalf("Expected to get nil, but found the key %d", key)
		}
	}
}

func TestKeyOrder(t *testing.T) {
	sl := NewSkipList[int, string]()
	dataInsert := [4]struct {
		k int
		v string
	}{
		{1, "One"},
		{2, "Two"},
		{3, "Three"},
		{4, "Four"},
	}

	for _, d := range dataInsert {
		sl.Insert(d.k, d.v)
	}

	keys := sl.Keys(0)
	for i := 0; i < len(keys)-1; i++ {
		if keys[i] > keys[i+1] {
			t.Errorf("Ordering mismatch at index %d: got key %v, want it to be greater than previous key %v", i+1, keys[i], keys[i+1])
		}
	}
}

func TestDisOrderInsert(t *testing.T) {
	sl := NewSkipList[int, string]()
	dataInsert := [5]struct {
		k int
		v string
	}{
		{9, "Nine"},
		{3, "Three"},
		{4, "Four"},
		{2, "Two"},
		{1, "One"},
	}

	for _, d := range dataInsert {
		sl.Insert(d.k, d.v)
	}

	keys := sl.Keys(0)
	for i := 0; i < len(keys)-1; i++ {
		if keys[i] > keys[i+1] {
			t.Errorf("Ordering mismatch at index %d: got key %v, want it to be greater than previous key %v", i+1, keys[i], keys[i+1])
		}
	}
}

func TestDuplicateInsert(t *testing.T) {
	sl := NewSkipList[int, string]()
	dataInsert := [5]struct {
		k int
		v string
	}{
		{9, "Nine"},
		{3, "Three"},
		{4, "Four"},
		{2, "Two"},
		{1, "One"},
	}

	for _, d := range dataInsert {
		sl.Insert(d.k, d.v)
	}

	dataDupInsert := [5]struct {
		k int
		v string
	}{
		{3, "One"},
		{9, "Three"},
		{4, "Two"},
		{1, "Four"},
		{2, "Nine"},
	}

	for _, d := range dataDupInsert {
		sl.Insert(d.k, d.v)
	}

	if sl.Size() != 5 {
        t.Errorf("Expected size 5 after duplicate insert, got %d", sl.Size())
    }

	for _, d := range dataDupInsert {
		result := sl.Find(d.k)
		if result.Value != d.v {
			t.Errorf("Expected to get [key,value] [%d,%v], but got [%d,%v]", d.k, d.v, result.Key, result.Value)
		}
	}
}

func TestRandomLevelDistribution(t *testing.T) {
    sl := NewSkipList[int, int]()
    insertCount := 1000
    
    for i := 0; i < insertCount; i++ {
        sl.Insert(i, i)
    }

    if sl.level == 0 {
        t.Error("Probabilistic error: Max level is still 0 after 1000 inserts. Check randomlevel() logic.")
    }

    counts := make(map[int]int)
    for l := 0; l <= sl.level; l++ {
        nodeCount := 0
        curr := sl.Head.Forward[l]
        for curr != nil {
            nodeCount++
            curr = curr.Forward[l]
        }
        counts[l] = nodeCount
    }

    t.Logf("Level distribution: %v", counts)

    for l := 0; l < sl.level; l++ {
        if counts[l] < counts[l+1] {
            t.Errorf("Level %d has fewer nodes (%d) than level %d (%d)", l, counts[l], l+1, counts[l+1])
        }
    }
}
