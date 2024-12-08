package learnbyexample

import "fmt"



func Calculater(a int , b int ,op rune) int {
    switch op {
    case '-':
        return a - b;
    case '+' : 
        return a + b;
    case '*' :
        return a * b;
    case '/' : 
        return a /b ;
    default :
        fmt.Println("Invalid operation");
        return -1;
}


}
