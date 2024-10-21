package main

func main() {

}

type LRUCache struct {
    data map[int]int
	dataMaxCap int
	usedItemsKeys []int
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
		data: make(map[int]int, capacity),
		dataMaxCap: capacity,
		usedItemsKeys: make([]int, 0),
	}
}

func (l *LRUCache) Get(key int) int {
    if v, found := l.data[key]; found {
		// usedItemKeyFound := false
		for i, usedItemKey := range l.usedItemsKeys {
			if usedItemKey == key {
				part1 := l.usedItemsKeys[:i]
				part2 := l.usedItemsKeys[i+1:]
				l.usedItemsKeys = part1
				l.usedItemsKeys = append(l.usedItemsKeys, part2...)
				l.usedItemsKeys = append(l.usedItemsKeys, usedItemKey)
				// usedItemKeyFound = true

				break
			}
		}

		// if !usedItemKeyFound {
		// 	l.usedItemsKeys = append(l.usedItemsKeys, key)
		// }

		return v
	}

	return -1
}


func (l *LRUCache) Put(key int, value int) {
	if _, found := l.data[key]; !found {
		if len(l.data) == l.dataMaxCap {
			leastUsedItemKey := l.usedItemsKeys[0]
			delete(l.data, leastUsedItemKey)

			l.usedItemsKeys = l.usedItemsKeys[1:]
		}
	} else if len(l.data) == l.dataMaxCap {
		l.usedItemsKeys = l.usedItemsKeys[1:]
	}

	l.data[key] = value
	l.usedItemsKeys = append(l.usedItemsKeys, key)
}
