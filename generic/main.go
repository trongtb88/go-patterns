package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Num interface {
	int | int8 | float64
}

func Add[T Num](a, b T) T {
	return a + b
}

type UserID int

func AddUserID[T ~int](a, b T) T {
	return a + b
}

//----------------------------------------------
// input : [1,2,3], (n) => n * 2
// output : [2,4,6]

func mapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

// --------------------------------------------
type CustomMap[T comparable, V int | string] map[T]V

func main() {
	result := Add(1.5, 3.5)
	fmt.Printf("result %+v\n", result)

	userID1 := UserID(1)
	userID2 := UserID(2)
	resultUserID := AddUserID(userID1, userID2)
	fmt.Printf("resultUserID %+v\n", resultUserID)

	//----------------------------------------
	resultMap := mapValues([]int{1, 2, 3}, func(i int) int {
		return i * 2
	})

	fmt.Printf("resultMap %+v\n", resultMap)

	resultMapFloat := mapValues([]string{"a", "b", "c"}, func(i string) string {
		return i + i
	})
	fmt.Printf("resultMap %+v\n", resultMapFloat)

	//-----------------------------------------
	m := make(CustomMap[int, string])
	m[1] = "abc"
	fmt.Printf("%+v", m)

}
