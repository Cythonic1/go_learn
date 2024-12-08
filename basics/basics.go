package basics;


import "fmt"


func BasicLoops(){

    // For loops
    for i := 0 ; i < 10 ; i++ {
        fmt.Println("The number is : ", i);
    }

    // Infinite loop similar to while loops
    for {
        fmt.Println("Infinite Loop");
        break;
    }

    // Condition while loop
    for 1 < 1 {
        fmt.Println("Condition loop");
        break;
    }

    for _, j:= range "Hello World"{
        fmt.Println(string(j));
    }


}
