package main

import (
	"fmt"
)

type testWrapper struct {
	TestCount int
	testData  []test
}

type test struct {
	bagCount      int
	candiesPerBag []int
	result        int
}

func main() {

	tw := testWrapper{}
	// read no. of tests
	_, err := fmt.Scanf("%d", &tw.TestCount)
	if err != nil {
		return
	}

	for i := 1; i <= tw.TestCount; i++ {
		t := test{}

		// read no. of bags
		_, err := fmt.Scanf("%d", &t.bagCount)
		if err != nil {
			return
		}

		length := t.bagCount
		t.candiesPerBag = make([]int, length)
		for i := 0; i < length; i++ {
			fmt.Scan(&t.candiesPerBag[i])
		}

		tw.testData = append(tw.testData, t)
	}

	// calculate result
	for idx, t := range tw.testData {
		sumCandies := 0
		for _, c := range t.candiesPerBag {
			sumCandies += c
		}

		avgCandies := sumCandies / len(t.candiesPerBag)
		for i := 0; i < len(t.candiesPerBag); i++ {
			if sumCandies%len(t.candiesPerBag) == 0 {
				if t.candiesPerBag[i] > avgCandies {
					t.result += (t.candiesPerBag[i] - avgCandies) * 2
				}
			} else {
				t.result = -1
			}
		}
		tw.testData[idx] = t
	}

	fmt.Println("Output:")
	for _, t := range tw.testData {
		fmt.Println(t.result)
	}

	return
}
