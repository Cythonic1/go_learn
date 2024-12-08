package basics

import "fmt";

func BasicFlowControl(){

    // The Lable is a word define some palce 
    // To start execution from there
    Lable1: var i int = 10;
    if i == 10{
        fmt.Println("The condition met");
    }else if i > 5 {
        fmt.Println("Some else if");
    }else{
        fmt.Println("Condition not met");

    }

     for i:= 0 ; i <= 6 ; i+=2 {
        if i == 6{
            println("Lable has been executed");
            // This how we use lable with goto key word
            goto Lable1;
        }
    }



}
