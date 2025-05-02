package main

type TimeMap struct {
	tmap map[string][]pair
}

type pair struct {
	value     string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{make(map[string][]pair)}
}

func (t *TimeMap) Set(key, value string, timestamp int) {
	t.tmap[key] = append(t.tmap[key], pair{value, timestamp})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	if _, exists := t.tmap[key]; !exists {
		return ""
	}
	pairs := t.tmap[key]
	l, r := 0, len(pairs)-1

	for l <= r {
		mid := (l + r) / 2
		if pairs[mid].timestamp <= timestamp {
			if mid == len(pairs)-1 || pairs[mid+1].timestamp > timestamp {
				return pairs[mid].value
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ""
}

func main() {}
