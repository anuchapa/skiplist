package skiplist

import (
	"cmp"
	"fmt"
	"math/rand"
	"time"
)

type SkipNode[K cmp.Ordered, V any] struct {
	Key     K
	Value   V
	Forward []*SkipNode[K, V]
}

func NewSkipNode[K cmp.Ordered, V any](key K, value V, level int) *SkipNode[K, V] {
	return &SkipNode[K, V]{Key: key, Value: value, Forward: make([]*SkipNode[K, V], level+1)}
}

type SkipList[K cmp.Ordered, V any] struct {
	Head   *SkipNode[K, V]
	level  int
	size   int
	random *rand.Rand
}

func (s *SkipList[K, V]) Level() int {
	return s.level
}

func (s *SkipList[K, V]) Size() int {
	return s.size
}

func NewSkipList[K cmp.Ordered, V any]() *SkipList[K, V] {
	var zeroKey K
	var zeroValue V
	return &SkipList[K, V]{
		Head:   NewSkipNode(zeroKey, zeroValue, 0),
		level:  -1,
		size:   0,
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func NewSkipListMap[K cmp.Ordered, V any](maps map[K]V) *SkipList[K, V] {
	list := NewSkipList[K, V]()
	for k, v := range maps {
		list.Insert(k, v)
	}
	return list
}

func (s *SkipList[K, V]) Get(index int, level int) (any, error) {
	if level > s.level {
		return nil, fmt.Errorf("level %v out of range %v", level, s.level)
	}

	x := s.Head
	for i := 0; i >= index; i-- {
		for x.Forward[level] != nil {
			x = x.Forward[level]
		}
	}

	return x, nil

	// if x != nil {
	// } else {
	// 	return nil, fmt.Errorf("This index value is nil")
	// }

}

func (s *SkipList[K, V]) Keys(level int) []K {
	if level > s.level {
		panic(fmt.Sprintf("level %v out of range %v", level, s.level))
	}

	keys := make([]K, 0)
	x := s.Head
	for cur := x.Forward[level]; cur != nil; {
		keys = append(keys, cur.Key)
		cur = cur.Forward[level]
	}

	return keys
}

func (s *SkipList[K, V]) Values(level int) []V {
	if level > s.level {
		panic(fmt.Sprintf("level %v out of range %v", level, s.level))
	}

	values := make([]V, 0)
	x := s.Head
	for cur := x.Forward[level]; cur != nil; {
		values = append(values, cur.Value)
		cur = cur.Forward[level]
	}

	return values
}

func (s *SkipList[K, V]) ToMap(level int) map[K]V {
	if level > s.level {
		panic(fmt.Sprintf("level %v out of range %v", level, s.level))
	}

	maps := map[K]V{}
	x := s.Head
	for cur := x.Forward[level]; cur != nil; {
		maps[cur.Key] = cur.Value
		cur = cur.Forward[level]
	}

	return maps
}

func (s *SkipList[K, V]) Find(key K) *SkipNode[K, V] {
	x := s.Head
	for i := s.level; i >= 0; i-- {
		for x.Forward[i] != nil && x.Forward[i].Key < key {
			x = x.Forward[i]
		}
	}
	x = x.Forward[0]
	if x != nil && x.Key == key {
		return x
	} else {
		return nil
	}
}

func (s *SkipList[K, V]) Insert(key K, value V) {
	newlevel := s.randomlevel()

	if newlevel > s.level {
		s.adjusthead(newlevel)
	}

	update := make([]*SkipNode[K, V], s.level+1)
	x := s.Head
	for i := s.level; i >= 0; i-- {
		for x.Forward[i] != nil && x.Forward[i].Key < key {
			x = x.Forward[i]
		}
		update[i] = x
	}

	x = x.Forward[0]
	if x != nil && x.Key == key {
		x.Value = value
		return
	}

	x = NewSkipNode(key, value, newlevel)

	for i := 0; i <= newlevel; i++ {
		x.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = x
	}

	s.size++
}

func (s *SkipList[K, V]) Remove(key K) {
	x := s.Head
	update := make([]*SkipNode[K, V], s.level+1)
	for i := s.level; i >= 0; i-- {
		for x.Forward[i] != nil && x.Forward[i].Key < key {
			x = x.Forward[i]
		}
		update[i] = x

		// if x.forward[i] != nil && *(x.forward[i].key) == key {
		// }
	}

	for i := 0; i < s.level; i++ {
		curr := update[i].Forward[i]
		if curr == nil {
			continue
		}

		if curr.Key != key {
			continue
		}

		update[i].Forward[i] = curr.Forward[i]
	}
}

func (s *SkipList[K, V]) randomlevel() int {
	var lev int
	for lev = 0; s.random.Int()%2 == 0; lev++ {
	}
	return lev
}

func (s *SkipList[K, V]) adjusthead(newlevel int) {
	head := s.Head
	temp := make([]*SkipNode[K, V], newlevel+1)
	copy(temp, head.Forward)
	head.Forward = temp
	s.level = newlevel
}

func (s *SkipList[K, V]) Show() {
	x := s.Head
	for x.Forward[0] != nil {
		fmt.Print(x.Forward[0].Value, "->")
		x = x.Forward[0]
	}
	fmt.Print("nil\n")
}
