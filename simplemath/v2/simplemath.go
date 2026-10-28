// Package simplemath is my solution to the chapter 10
// of Bodner's Learning Go book (modules)
package simplemath

type addArgument interface {
  int | float64
}

// Add adds two ints
// 
// it has two parameters: integers being added
// it returns their sum
func Add[T addArgument](a, b T) T {
  return a + b
}
