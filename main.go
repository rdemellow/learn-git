package main

import ("fmt")

func main() {
    maxNum := 100

    // I'm a cool comment

    for i := 0; i < maxNum; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i % 3 == 0 {
            fmt.Println("Fizz");
        } else if i % 5 == 0 {
            fmt.Println("Buzz");
        } else {
            fmt.Println(i);
        }
    }
}
