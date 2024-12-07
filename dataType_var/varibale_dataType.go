
package datatype_var

// Import extrnal packages
import "fmt"


// Global varibale can be asssigned and not used.
// can be used within the function or outside as here
var myint int = 10;
func DataType() {
    // Can not be set and not used
    // Also must be set localy to the function and can not be global
    myvar := 10

    // Print some data to the stdout
    fmt.Println("Hello, world.")
    fmt.Println(myvar);
    i := 10
    fmt.Println(i)
}
