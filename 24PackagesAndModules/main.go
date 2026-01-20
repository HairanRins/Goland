// 24. Packages and Modules
package main

import (
	"fmt"
	"math"
	"strings"

	// Import our custom packages
	"github.com/example/24packagesandmodules/calculator"
	"github.com/example/24packagesandmodules/utils"
)

// 1. Using standard library packages
func standardLibraryExample() {
	fmt.Println("=== Standard Library Packages ===")

	// math package
	fmt.Printf("Square root of 16: %.2f\n", math.Sqrt(16))
	fmt.Printf("Pi value: %.6f\n", math.Pi)

	// strings package
	text := "Hello, World!"
	fmt.Printf("Upper case: %s\n", strings.ToUpper(text))
	fmt.Printf("Contains 'World': %t\n", strings.Contains(text, "World"))
	fmt.Printf("Split by comma: %v\n", strings.Split(text, ","))
}

// 2. Using our custom packages
func customPackagesExample() {
	fmt.Println("\n=== Custom Packages ===")

	// Using calculator package
	fmt.Println("Calculator Package:")
	result := calculator.Add(10, 5)
	fmt.Printf("10 + 5 = %d\n", result)

	result = calculator.Multiply(4, 7)
	fmt.Printf("4 * 7 = %d\n", result)

	result = calculator.Subtract(20, 8)
	fmt.Printf("20 - 8 = %d\n", result)

	// Using utils package
	fmt.Println("\nUtils Package:")
	greeting := utils.Greet("Alice")
	fmt.Println(greeting)

	reversed := utils.ReverseString("hello")
	fmt.Printf("Reversed 'hello': %s\n", reversed)

	formatted := utils.FormatNumber(1234567.89)
	fmt.Printf("Formatted number: %s\n", formatted)
}

// 3. Package visibility and naming
func packageVisibilityExample() {
	fmt.Println("\n=== Package Visibility ===")

	// Public functions (capitalized) - accessible
	fmt.Println("Public function:", calculator.PublicFunction())

	// Private functions (lowercase) - not accessible from outside
	// calculator.privateFunction() // This would cause compile error

	// Public struct
	person := calculator.NewPerson("John", 30)
	fmt.Printf("Person: %+v\n", person)

	// Accessing public field
	fmt.Printf("Person name: %s\n", person.GetName())
	fmt.Printf("Person age: %d\n", person.GetAge())
}

// 4. Package initialization
func packageInitializationExample() {
	fmt.Println("\n=== Package Initialization ===")

	// The init() function in packages runs automatically
	// Check calculator package for init() function example

	fmt.Printf("Calculator version: %s\n", calculator.GetVersion())
	fmt.Printf("Utils config: %+v\n", utils.GetConfig())
}

// 5. Aliasing imports
func aliasingExample() {
	fmt.Println("\n=== Import Aliasing ===")

	// Using alias for imported package
	// import str "strings"
	// str.ToUpper("hello")

	// Using blank import for side effects
	// import _ "database/sql"

	fmt.Println("Import aliasing demonstrated in imports section")
}

// 6. Third-party packages example (simulated)
func thirdPartyExample() {
	fmt.Println("\n=== Third-party Packages ===")

	// Examples of popular third-party packages:
	fmt.Println("Popular third-party packages:")
	fmt.Println("- github.com/gin-gonic/gin (web framework)")
	fmt.Println("- github.com/gorilla/mux (HTTP router)")
	fmt.Println("- github.com/stretchr/testify (testing)")
	fmt.Println("- github.com/spf13/cobra (CLI framework)")
	fmt.Println("- gorm.io/gorm (ORM)")
	fmt.Println("- github.com/go-redis/redis (Redis client)")
}

func main() {
	fmt.Println("=== Go Packages and Modules ===")

	standardLibraryExample()
	customPackagesExample()
	packageVisibilityExample()
	packageInitializationExample()
	aliasingExample()
	thirdPartyExample()

	fmt.Println("\n=== Packages and Modules Complete ===")
}

/*
PACKAGES AND MODULES KEY CONCEPTS:

1. PACKAGES:
- Collection of related Go source files
- Provide namespace organization
- Control visibility through capitalization
- Enable code reuse and modularity

2. PACKAGE VISIBILITY:
- Capitalized names = public (exported)
- Lowercase names = private (unexported)
- Applies to functions, variables, types, constants
- Controls access across package boundaries

3. IMPORTING PACKAGES:
- import "package/path"
- Multiple imports: import ( "pkg1" "pkg2" )
- Aliasing: import alias "package/path"
- Blank import: import _ "package/path" (side effects only)

4. STANDARD LIBRARY:
- Extensive built-in packages
- Well-documented and tested
- Covers common programming tasks
- Examples: fmt, strings, math, net/http, encoding/json

5. CUSTOM PACKAGES:
- Create your own packages for organization
- Place in separate directories
- Package name = directory name (usually)
- Export only what's necessary

6. GO MODULES:
- Dependency management system
- Version control for packages
- go.mod file defines module
- go.sum file ensures integrity

7. MODULE COMMANDS:
- go mod init module-name
- go mod tidy (clean dependencies)
- go get package@version
- go mod download

8. BEST PRACTICES:
- Keep packages focused and small
- Use descriptive package names
- Minimize exported API
- Document public functions
- Follow Go naming conventions
- Use semantic versioning

9. PACKAGE STRUCTURE:
- Each package in its own directory
- All .go files in same package directory
- One package per directory
- Package declaration at top of each file

10. DEPENDENCY MANAGEMENT:
- Modules track exact versions
- go.mod lists direct dependencies
- go.sum records checksums
- Reproducible builds
*/
