go-to-zzz
=========

*a clone of https://github.com/Tokette/zzz written in go*

This is a small exercise to get a small amount of practice with Go.

Running *go run zzz.go* will recreate the same default behavior of running
*zzz.py* from Tokette's project however commandline arguments are still
missing(see TODO for more info).

*go-to-zzz* is licensed under the MIT license because licenses are cool things
to have.

TODO:
-----
The only real thing left to do is implement command line args. The following
text is from running *zzz.py -h*

```
usage: zzz [-h] [-c SLEEP_CYCLES] [-d SLEEP DELAY] [-l CYCLE_LENGTH]

Calculates wake up times based on sleep cycles

optional arguments:
  -h --help       Show this help message and exit
  -c SLEEP_CYCLES, --sleep-cycles SLEEP CYCLES
                        number of sleep cycles to calculate
  -d SLEEP_DELAY, --sleep-delay SLEEP_DELAY
                        offset in minutes used to compensate for the time
                        taken to fall asleep
  -l CYCLE_LENGTH, --cycle-length CYCLE_LENGTH
                        length i nminutes to comoplete sleep cycle
```

The goal is to completely reecreate this behavior with go-to-zzz. It shouldn't
be hard to implement the commandline args.
