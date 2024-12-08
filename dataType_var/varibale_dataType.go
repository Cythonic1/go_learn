
package datatype_var

// Import extrnal packages
import "fmt"


// Global varibale can be asssigned and not used.
// can be used within the function or outside as here
var myint_4 int = 10;
// Strings
var string_type string = "Hello World";
// Boolean
var bool_type bool = true;
// Pointer
var ptr *int = &myint_4;

// For a function to be exportable you need to set
// it as CamleCase
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

// Other types that are intresting.
// 4 or 8 bytes unsined int and it goes until 64 bytes
var myunit uint = 10;
var myunit_64 uint64 = 1000000000;

// For floating we have complex and float types
var myfloat float32 = 1.243456546;
//complex are reporesnt imaginary numbers
var mycomplex complex64 = 1.34547567;

// Other types
var mybyte byte = 255;

// rune same as int32
var myrune rune = 200;

var myptra uintptr;


