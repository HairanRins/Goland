// Custom calculator package
package calculator

import "fmt"

// Package-level variable (public)
var Version = "1.0.0"

// Package-level variable (private)
var internalCounter int = 0

// init() function runs automatically when package is imported
func init() {
	fmt.Println("Calculator package initialized!")
	internalCounter = 100
}

// Public functions (capitalized) - accessible from other packages

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// Subtract returns the difference of two integers
func Subtract(a, b int) int {
	return a - b
}

// Divide returns the quotient of two integers
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// PublicFunction demonstrates a public function
func PublicFunction() string {
	return "This is a public function from calculator package"
}

// GetVersion returns the package version
func GetVersion() string {
	return Version
}

// Private functions (lowercase) - not accessible from other packages

// privateFunction demonstrates a private function
func privateFunction() string {
	return "This is a private function - not accessible outside package"
}

// internalHelper shows internal helper function
func internalHelper(x int) int {
	return x * 2
}

// Public struct (capitalized) - accessible from other packages

// Person represents a person with name and age
type Person struct {
	Name string // Public field
	age  int    // Private field
}

// NewPerson creates a new Person with validation
func NewPerson(name string, age int) *Person {
	if name == "" {
		name = "Unknown"
	}
	if age < 0 {
		age = 0
	}
	return &Person{
		Name: name,
		age:  age,
	}
}

// GetName returns the person's name (public method)
func (p *Person) GetName() string {
	return p.Name
}

// GetAge returns the person's age (public method)
func (p *Person) GetAge() int {
	return p.age
}

// SetAge sets the person's age with validation (public method)
func (p *Person) SetAge(age int) {
	if age >= 0 {
		p.age = age
	}
}

// private method (lowercase) - not accessible from other packages
func (p *Person) privateMethod() string {
	return "This is a private method"
}

// Public interface (capitalized) - accessible from other packages

// Calculator defines calculator operations
type Calculator interface {
	Calculate(a, b int) int
}

// AdvancedCalculator implements Calculator interface
type AdvancedCalculator struct {
	operation string
}

// NewAdvancedCalculator creates a new advanced calculator
func NewAdvancedCalculator(operation string) *AdvancedCalculator {
	return &AdvancedCalculator{operation: operation}
}

// Calculate implements the Calculator interface
func (ac *AdvancedCalculator) Calculate(a, b int) int {
	switch ac.operation {
	case "add":
		return Add(a, b)
	case "multiply":
		return Multiply(a, b)
	case "subtract":
		return Subtract(a, b)
	default:
		return 0
	}
}

// GetOperation returns the current operation
func (ac *AdvancedCalculator) GetOperation() string {
	return ac.operation
}
