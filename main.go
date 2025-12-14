package main

import (
	"fmt"
)

func main() {
	//s := "gopher"
	//fmt.Printf("Hello and welcome, %s!\n", s)
	//
	//for i := 1; i <= 5; i++ {
	//	//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
	//	// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
	//	fmt.Println("i =", 100/i)
	//}

	// Buy and Sell Stocks
	stockPrices1 := []int{1, 2, 3, 4, 5}
	maxProfitGreedy(stockPrices1)
	fmt.Printf("Stock Prices: %v\n", stockPrices1)
	fmt.Printf("Max profit: %d\n\n", maxProfitGreedy(stockPrices1))

	stockPrices2 := []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Stock Prices: %v\n", stockPrices2)
	fmt.Printf("Max profit: %d\n", maxProfitGreedy(stockPrices2))

	// Shortest word edit path
	// Example 1
	//source1 := "bit"
	//target1 := "dog"
	//words1 := []string{"but", "put", "big", "pot", "pog", "dog", "lot"}
	//result1 := shortestWordEditPath(source1, target1, words1)
	//fmt.Printf("Source: %s, Target: %s, Result: %d (Expected: 5)\n", source1, target1, result1)

	// Example 2
	//source2 := "no"
	//target2 := "go"
	//words2 := []string{"to"}
	//result2 := shortestWordEditPath(source2, target2, words2)
	//fmt.Printf("Source: %s, Target: %s, Result: %d (Expected: -1)\n", source2, target2, result2)

	// Example 3 (A case where source is not in words, but is the starting point)
	//source3 := "a"
	//target3 := "c"
	//words3 := []string{"b", "c"}
	//result3 := shortestWordEditPath(source3, target3, words3)
	//fmt.Printf("Source: %s, Target: %s, Result: %d (Expected: 2)\n", source3, target3, result3)

	// Busiest time in the mall
	//data := [][]int{
	//	{1487799425, 14, 1},
	//	{1487799425, 4, 1},
	//	{1487799425, 2, 1},
	//	{1487800378, 10, 1},
	//	{1487800478, 18, 1},
	//	{1487901013, 1, 1},
	//	{1487901211, 7, 1},
	//	{1487901211, 7, 1},
	//}
	//// Expected Output : 1487901211
	//fmt.Println(findBusiestPeriod(data))
	//// Deletion Distance
	//str1 := "frog"
	//str2 := "dog"
	//fmt.Println(deleteDistance(str1, str2))
}
