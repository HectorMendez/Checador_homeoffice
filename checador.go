//Package by Héctor Méndez
//Licensing under GPLv3
// TODO we need an GUI in order to be user friendly 
//need a scape routine in case of end journal for counting the hours by weekdays
//need some help to review and improve code with compatibility for windows and MAC

package main

import (
    "fmt"
    "os"
    "time"
    "os/exec"
)
//Main routine, with var declaration and creation of buffer for a signals
func main() {
    var working float64
    var relax   float64
    var switchact bool       = false
    var switch2relax bool    = true
    var switch2working bool  = false

    fmt.Println("Type any key for start the Workday")
    fmt.Println("The counter starts with relaxing type any key for start working")
    fmt.Println("You start workint at " ,time.Now())
    ch := make(chan string)
    go func(ch chan string) {
        // disable input buffering
        exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
        // do not display entered characters on the screen
        exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
        var b []byte = make([]byte, 1)
        //Routine for check for standard input with no wait for a key
        for {
            os.Stdin.Read(b)
            ch <- string(b)
        }
    }(ch)

    //The intention with this fork was do a Thiker and with this controll the flow of time
    for {
        select {
            
            //In case of a signal by key lanch this case, change status of working to relax
            //and show time elapsed
            case stdin, _ := <-ch:
                fmt.Println("\n \n Keys pressed:", stdin)
                if switch2relax == true {
                    fmt.Println("\n Now you are relaxing\n \n")
                }else {
                    fmt.Println("\n Now you are working \n \n")
                }
                fmt.Printf("working for %f Hours \n", ((working/60.0)/60.0))
                fmt.Printf("Relaxing for %f Hours \n", (relax/60.0)/60.0)

                switch2relax = !switch2relax
                switch2working = !switch2working
                switchact = true
            
            //By default the program is count time, if the routin is not counting time
            //working count time relaxing, this script has no pause
            default:
                if switchact == true {
                    if switch2relax == true {
                        working=1.0+working
                         if ((working/60.0)/60.0) == 8.00000 {
                            exec.Command("mplayer", "stopwork.mp3").Run()
                        }
                    //fmt.Println("Working for ", working)
                    }else{
                    relax=1.0+relax
                    //fmt.Println("relaxing for ", relax)
                    }
                }
        }
        time.Sleep(time.Second * 1)
    }
}
