package main

import (
	"slices"
	"sync"
)

// ConcurrentBFSQueries concurrently processes BFS queries on the provided graph.
// - graph: adjacency list, e.g., graph[u] = []int{v1, v2, ...}
// - queries: a list of starting nodes for BFS.
// - numWorkers: how many goroutines can process BFS queries simultaneously.
//
// Return a map from the query (starting node) to the BFS order as a slice of nodes.
// YOU MUST use concurrency (goroutines + channels) to pass the performance tests.

type BFSPath struct {
	Initial int
	Path    []int
}

var wg sync.WaitGroup

func ConcurrentBFSQueries(graph map[int][]int, queries []int, numWorkers int) map[int][]int {
	// TODO: Implement concurrency-based BFS for multiple queries.
	// Return an empty map so the code compiles but fails tests if unchanged.
	queriesChan := make(chan int, len(queries))
	output := make(chan *BFSPath, len(queries))
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go SearchWorker(graph, queriesChan, output)
	}
	for _, query := range queries {
		queriesChan <- query
	}
	close(queriesChan)
	wg.Wait()
	close(output)
	result := make(map[int][]int)
	for queryResult := range output {
		result[queryResult.Initial] = queryResult.Path
	}
	return result
}

func SearchWorker(graph map[int][]int, queryInitialNode chan int, output chan<- *BFSPath) {
	defer wg.Done()
	for {
		initialNode, ok := <-queryInitialNode
		if !ok {
			break
		}
		var visitedNodes []int
		nodeQueue := []int{initialNode}
		for len(nodeQueue) > 0 {
			currentNode := nodeQueue[0]
			visitedNodes = append(visitedNodes, currentNode)
			var nextNodes []int
			for _, n := range graph[currentNode] {
				if !slices.Contains(nodeQueue, n) && !slices.Contains(visitedNodes, n) {
					nextNodes = append(nextNodes, n)
				}
			}
			nodeQueue = append(nodeQueue[1:], nextNodes...)

		}
		result := BFSPath{initialNode, visitedNodes}
		output <- &result
	}
}
func main() {
	// You can insert optional local tests here if desired.
}
