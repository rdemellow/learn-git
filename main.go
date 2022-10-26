package main

import ("fmt")

func isFactor(value int, factor int) bool {
    return value % factor == 0;
}

func main() {
    maxNum := 100

    for i := 0; i < maxNum; i++ {
        isFactorOf5 := isFactor(i, 5);
        isFactorOf3 := isFactor(i, 3);

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
