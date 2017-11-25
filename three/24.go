package three

import (
)

var visited map[string]bool

func Visit(url string) {
	if visited == nil {
		visited = make(map[string]bool)
	}
	visited[url] = true
}

func Visited(url string) bool {
	if visited == nil {
		visited = make(map[string]bool)
	}
	return visited[url]
}
