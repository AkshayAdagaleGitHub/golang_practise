package main

import (
	"fmt"
	_ "fmt"
)

func shortestWordEditPath(source string, target string, words []string) int {

	if len(source) != len(target) {
		return -1
	}

	wordSet := make(map[string]bool)
	for _, word := range words {
		if len(word) == len(source) {
			wordSet[word] = true
		}
	}

	if _, found := wordSet[target]; !found {
		return -1
	}

	if source == target {
		return 0
	}

	queue := []string{source}
	visited := make(map[string]bool)
	visited[source] = true
	distance := 0

	for len(queue) > 0 {
		distance++
		currentLevelSize := len(queue)

		for i := 0; i < currentLevelSize; i++ {
			currentWord := queue[0]
			queue = queue[1:]

			for j := 0; j < len(currentWord); j++ {
				originalChar := currentWord[j]

				for charCode := 'a'; charCode <= 'z'; charCode++ {
					if byte(charCode) == originalChar {
						continue
					}
					newWordBytes := []byte(currentWord)
					newWordBytes[j] = byte(charCode)
					newWord := string(newWordBytes)

					if newWord == target {
						return distance
					}
					if _, valid := wordSet[newWord]; valid {
						if _, wasVisited := visited[newWord]; !wasVisited {
							visited[newWord] = true
							queue = append(queue, newWord)
						}
					}
				}
			}
		}
	}
	return -1
}

// Helper function to run the examples
func main() {
	// Example 1
	source1 := "bit"
	target1 := "dog"
	words1 := []string{"but", "put", "big", "pot", "pog", "dog", "lot"}
	result1 := shortestWordEditPath(source1, target1, words1)
	fmt.Printf("Source: %s, Target: %s, Result: %d (Expected: 5)\n", source1, target1, result1)

	// Example 2
	source2 := "no"
	target2 := "go"
	words2 := []string{"to"}
	result2 := shortestWordEditPath(source2, target2, words2)
	fmt.Printf("Source: %s, Target: %s, Result: %d (Expected: -1)\n", source2, target2, result2)

	// Example 3 (A case where source is not in words, but is the starting point)
	source3 := "a"
	target3 := "c"
	words3 := []string{"b", "c"}
	result3 := shortestWordEditPath(source3, target3, words3)
	fmt.Printf("Source: %s, Target: %s, Result: %d (Expected: 2)\n", source3, target3, result3)
}
