package main

import (
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
