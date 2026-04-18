package main

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers.
func Multiply(a, b int) int {
	return a * b
}

// IsEven reports whether n is an even number.
func IsEven(n int) bool {
	return n%2 == 0
}

func displayName(name string) string {
	return "Hello, " + name + "!"
}
