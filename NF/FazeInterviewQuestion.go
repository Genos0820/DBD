package main

import "fmt"

func main() {
	numCourses := 6
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}, {4, 2}, {5, 3}}
	fmt.Println(prerequisites)
	fmt.Println(finishCourse(numCourses, prerequisites))
}

func finishCourse(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	//setting prerquisites in graph
	for _, pre := range prerequisites {
		c1 := pre[0]
		c2 := pre[1]
		graph[c1] = append(graph[c1], c2) //to reach c1 you need to visit c2
	}

	for i := 0; i < numCourses; i++ {
		if graphHasCycle(graph, visited, i) {
			return false
		}
	}
	return true
}

func graphHasCycle(graph [][]int, visited []int, i int) bool {
	if visited[i] == 1 {
		return true
	}

	if visited[i] == -1 {
		return false
	}

	visited[i] = 1

	for _, val := range graph[i] {
		if graphHasCycle(graph, visited, val) {
			return true
		}
	}
	visited[i] = -1
	return false

}

//Approach: We are using the the DFS(depth first search) for the solution to detect the cycle.

// -Intitialy we are creating the graph with all the courses as node in the graph.
// -We are defining the dependency for the node to be visited using the prerequisites array.
// -Afterwords we are checking that if there is a cycle in graph or not using the graphHasCycle function.
// -If Graph has cycle then the course can not be completed and if the graph doesn't have cycle the courses can be completed.
// -The graphHasCycle function uses the DFS traversal to detect the cycle in the graph.
