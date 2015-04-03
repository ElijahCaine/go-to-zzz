package main

//import "os"
import "fmt"
import "time"

func main() {

    t := time.Now()
    s := t.Unix()

    cycle_length := 90
    sleep_delay := 15
    cycles := 6

    fmt.Printf("Current Time: %02d:%02d\n",
        time.Unix(s, 0).Hour(), time.Unix(s, 0).Minute())

    s += int64(sleep_delay * 60)

    fmt.Printf("Sleeping Begins in 15 Minutes At: %02d:%02d\n",
        time.Unix(s, 0).Hour(), time.Unix(s, 0).Minute())

    fmt.Printf("Using %d Minute Sleep Cycles\n", cycle_length)

    for i:=1; i <= cycles; i++ {
        s += 5400
        fmt.Printf("%d Cycle(s) %.2f Hours: Wake up at %02d:%02d\n", i, float32(i)*float32(1.5), time.Unix(s, 0).Hour(), time.Unix(s, 0).Minute())
    }

}

func args() {

//    argsWithProg := os.Args
//    argsWithoutProg := os.Args[1:]

//    arg := os.Args[3]

}
