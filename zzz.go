package main

import "os"
import "fmt"
import "time"
import "strconv"
import "reflect"

func main() {

    t := time.Now()
    s := t.Unix()

    vals := args()

    if vals[3] == 0 {
        cycle_length := 90
        sleep_delay := 15
        cycles := 6

        if vals[0] != 0 {
            cycle_length = vals[0]
        }
        if vals[1] != 0 {
            sleep_delay = vals[1]
        }
        if vals[2] != 0 {
            cycles = vals[2]
        }

        fmt.Printf("Current Time: %02d:%02d\n",
            time.Unix(s, 0).Hour(), time.Unix(s, 0).Minute())

        s += int64(sleep_delay * 60)

        fmt.Printf("Sleeping Begins in 15 Minutes At: %02d:%02d\n",
            time.Unix(s, 0).Hour(), time.Unix(s, 0).Minute())

        fmt.Printf("Using %d Minute Sleep Cycles\n", cycle_length)

        for i:=1; i <= cycles; i++ {
            s += 5400
            fmt.Printf("%d Cycle(s) %.2f Hours: Wake up at %02d:%02d\n", i, float32(i)*float32(cycle_length/60), time.Unix(s, 0).Hour(), time.Unix(s, 0).Minute())
        }
    }

}

func args() [4]int {

    args := os.Args
    ret := [4]int{0,0,0,0}

    if len(args) > 1 {
        for i := 1; i < len(args); i+=2 {
            if args[i] == "-h" {
                help_msg()
                ret[3] = 1
            } else if args[i] == "--help" {
                help_msg()
                ret[3] = 1
            } else if args[i] == "-c" {
                val, _ := strconv.Atoi(args[i+1])
                if reflect.TypeOf(val).Kind() == reflect.Int {
                    ret[2] = val
                }
            } else if args[i] == "--sleep-cycles" {
                val, _ := strconv.Atoi(args[i+1])
                if reflect.TypeOf(val).Kind() == reflect.Int {
                    ret[2] = val
                }
            } else if args[i] == "-d" {
                val, _ := strconv.Atoi(args[i+1])
                if reflect.TypeOf(val).Kind() == reflect.Int {
                    ret[1] = val
                }
            } else if args[i] == "--sleep-delay" {
                val, _ := strconv.Atoi(args[i+1])
                if reflect.TypeOf(val).Kind() == reflect.Int {
                    ret[1] = val
                }
            } else if args[i] == "-l" {
                val, _ := strconv.Atoi(args[i+1])
                if reflect.TypeOf(val).Kind() == reflect.Int {
                    ret[0] = val
                }
            } else if args[i] == "--cycle-length" {
                val, _ := strconv.Atoi(args[i+1])
                if reflect.TypeOf(val).Kind() == reflect.Int {
                    ret[0] = val
                }
            }
        }
    } else { return ret }
    return ret

}

func help_msg() {

    fmt.Println("usage: zzz [-h] [-c SLEEP_CYCLES] [-d SLEEP_DELAY] [-l CYCLE_LENGTH]")
    fmt.Println(" ")
    fmt.Println("Calculates wake up times based on sleep cycles")
    fmt.Println(" ")
    fmt.Println("optional arguments:")
    fmt.Println("  -h, --help            show this help message and exit")
    fmt.Println("  -c SLEEP_CYCLES, --sleep-cycles SLEEP_CYCLES")
    fmt.Println("                        number of sleep cycles to calculate")
    fmt.Println("  -d SLEEP_DELAY, --sleep-delay SLEEP_DELAY")
    fmt.Println("                        offset in minutes used to compensate for the time")
    fmt.Println("                        taken to fall asleep")
    fmt.Println("  -l CYCLE_LENGTH, --cycle-length CYCLE_LENGTH")
    fmt.Println("                        length in minutes of a complete sleep cycle")

}
