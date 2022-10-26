package main

import ("fmt")

func main() {
    maxNum := 100

    for i := 0; i < maxNum; i++ {
        if i % 3 == 0 {
            fmt.Println("Foo");
        } else {
            fmt.Println(i);
        }
    }
}
