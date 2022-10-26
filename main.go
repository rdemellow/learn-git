package main

import ("fmt")

func main() {
    maxNum := 100

    for i := 0; i < maxNum; i++ {
        isFactorOf5 := i % 5 == 0
        isFactorOf3 := i % 3 == 0

        if isFactorOf3 && isFactorOf5 {
            fmt.Println("FizzBuzz")
        } else if isFactorOf3 {
            fmt.Println("Fizz");
        } else if isFactorOf5 {
            fmt.Println("Buzz");
        } else {
            fmt.Println(i);
        }
    }
}
