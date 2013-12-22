package speechtext
import (
    "fmt"
    "os/exec"
    "strings"
)




func check(err error)  { 
    if err != nil {
        fmt.Println("Error occured: ", err)
    }
}


func speaktext(nonplustext string)  {
    plustext := plussify_string(nonplustext) 
    var err error
    querystring := "http://tts-api.com/tts.mp3?q=" + plustext + " &"
    
    err = exec.Command("xdg-open", querystring).Run()
   check(err) 
   
      
   err = exec.Command("wmctrl","-a","ROX").Run()
    check(err)

    
   }


func plussify_string(nonplussed string) string {

    plussedstring := strings.Replace(nonplussed, " ", "+",-1)



return plussedstring
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



   
func SpeechText(selection string) {


if selection=="9" {
    var err error
    err = exec.Command("google-chrome", "robotnature.mp3").Start()
    check(err) 

} else if selection=="10"{

    var err error
    err = exec.Command("google-chrome", "du.mp3").Start()
    check(err)
}else{
speaktext(selection)    
}





}
