package main

import (
    //"cfg"
    "fmt"
    "os/exec"
)

// #include <stdio.h>
// #include <errno.h>
import "C"




func check(err error)  { 
    if err != nil {
        fmt.Println("Error occured: ", err)
    }
}


func speaktext(plustext string)  {

    var err error
    querystring := "http://tts-api.com/tts.mp3?q=" + plustext + " &"
    
    err = exec.Command("xdg-open", querystring).Run()
   check(err) 
   err = exec.Command("wmctrl","-a xterm").Start()
    check(err)

    
   }


func plussify_string(nonplussed string) string {

    
fmt.Println("hi")



return "hi"





}

// Add timer loop between replies
// If no keystroke for 10 seconds, say something
// in a list of ten comments
// like "Hey, have you kids gotten tired of that rubber
// band bracelete machine yet?
// It should only do one of those for every question
// not all of them... So if the user waits, it will 
// spontaneously ask a question.. but just one..
// "Weve come a long way, from the days of the dirt necklace...dont you agree?"
// every 10 questions it has to kill the chrome instance
// or things will slow down



   
func main() {

speaktext("Rogers+heart+is+evasive+...+like+a+ghostcrab")
ac := []byte("How are you doing?")

fmt.Println(ac)







}
