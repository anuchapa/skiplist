# SkipList (Generic Skip List for Go)

A generic Skip List implementation in Go, supporting ordered keys
Designed as a lightweight, in-memory ordered map with logarithmic-time operations.

This project is primarily for learning and experimentation, but the implementation is complete and usable.

## Features

- Generic key/value support (Go 1.20+)
- Ordered keys via cmp.Ordered
- Average O(log n) time complexity
- Insert / Find / Remove
- Level-based traversal
- Convert to slice of keys, or values
- Deterministic structure with randomized levels


## Usage

### Create a Skip List 
```
list := NewSkipList[int, string](DefaultCmp[int])
```

### Insert
```
list.Insert(10, "ten")
list.Insert(5, "five")
list.Insert(20, "twenty")
```

### Find
```
node := list.Find(10)
if node != nil {
    fmt.Println(node.Value)
}
```
### Remove
```
list.Remove(5)
```

