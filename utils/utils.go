package utils

import "sync"

func MapCount(m sync.Map) int {
	var n int
	m.Range(func(key, value interface{}) bool {
		n++
		return true
	})
	return n
}
