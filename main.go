// https://www.codechef.com/problems/ZCOSCH
package main

import "fmt"

func main() {
    var rank int
    _, _ = fmt.Scan(&rank)

    if rank >= 1 && rank <= 50 {
        fmt.Println(100)
    } else if rank >= 51 && rank <= 100 {
        fmt.Println(50)
    } else {
        fmt.Println(0)
    }
}
