package graph

import (
	"fmt"
	"math/bits"
)

// A Clique is a subset of vertices of an undirected graph such that every two
// distinct vertices in the clique have same edge.
func IsClique[T any](vertices []T, haveEdge func(from, to T) bool) bool {
	for i, from := range vertices {
		for _, to := range vertices[i+1:] {
			if !haveEdge(from, to) {
				return false
			}
		}
	}
	return true
}

// A Clique is a subset of vertices of an undirected graph such that every two
// distinct vertices in the clique have same edge. A K-Clique is a clique that
// has K vertices.
func IsKClique[T any](vertices []T, k int, haveEdge func(from, to T) bool) bool {
	if len(vertices) != k {
		return false
	}

	return IsClique(vertices, haveEdge)
}

type AdjacentMask struct{ masks []uint32 }

func newAdjacentMask[T any](vertices []T, haveEdge func(from, to T) bool) AdjacentMask {
	n := len(vertices)
	adjMasks := make([]uint32, n)
	for i := range n {
		adjMasks[i] |= 1 << i
		for j := i + 1; j < n; j++ {
			if haveEdge(vertices[i], vertices[j]) {
				adjMasks[i] |= 1 << j
				adjMasks[j] |= 1 << i
			}
		}
	}
	for a, m := range adjMasks {
		fmt.Printf("adjMasks %d: %08b\n", a, m)
	}
	return AdjacentMask{masks: adjMasks}
}

func (am *AdjacentMask) intersection(mask uint32) uint32 {
	if mask == 0 {
		return 0
	}
	intersection := uint32(1<<len(am.masks) - 1)
	for i := 0; mask > 0; i++ {
		if mask&1 == 1 {
			intersection &= am.masks[i]
		}
		mask >>= 1
	}
	return intersection
}

// Find all K-Cliques in graph that has no more than 32 vertices
func FindKClique[T any](vertices []T, k int, haveEdge func(from, to T) bool) [][]T {
	n := len(vertices)
	if k <= 0 || n > 32 || n < k {
		return nil
	}

	cliques := [][]T{}
	am := newAdjacentMask(vertices, haveEdge)

	for mask := uint32(1); mask < (1 << n); mask++ {
		verts := bits.OnesCount(uint(mask))
		if verts != k {
			continue
		}

		if (am.intersection(mask) & mask) == mask {
			clique := make([]T, 0)
			for i := 0; i < n; i++ {
				if mask&(1<<i) != 0 {
					clique = append(clique, vertices[i])
				}
			}
			cliques = append(cliques, clique)
		}
	}
	return cliques
}

// Find all connected sub-graphs in graph that has no more than 32 vertices
func FindConnectedGraph[T any](vertices []T, haveEdge func(from, to T) bool) [][]T {
	n := len(vertices)
	if n > 32 {
		return nil
	}

	visited := make([]bool, n)
	conGraphs := [][]T{}

	for i := range vertices {
		if !visited[i] {
			conGraph := _bfs(i, vertices, haveEdge, visited)
			conGraphs = append(conGraphs, conGraph)
		}
	}

	return conGraphs
}

func _bfs[T any](startIndex int, vertices []T, haveEdge func(from, to T) bool, visited []bool) []T {
	queue := []int{startIndex}
	connected := make([]bool, len(vertices))
	connected[startIndex] = true
	visited[startIndex] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for index := range vertices {
			if !visited[index] && haveEdge(vertices[current], vertices[index]) {
				visited[index] = true
				connected[index] = true
				queue = append(queue, index)
			}
		}
	}

	conGraph := []T{}
	for i := range connected {
		if connected[i] {
			conGraph = append(conGraph, vertices[i])
		}
	}
	return conGraph
}
