// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This example show a basic menu similar to that found in the ncurses
 * examples from TLDP. Errors are intentionally discarded */

package main

import . "code.google.com/p/goncurses"
import ( 
    "speechtext"
    "github.com/jimlawless/cfg"
    "strconv"
    "os/exec"
    "fmt"
)
    
    
    
const (
	HEIGHT = 16
    WIDTH  = 40
)



func check(err error) {

if err != nil {
    fmt.Println("There was an error: ", err)
}



}


func main() {
	
    // config loader code
    storymap := make(map[string]string)
    err := cfg.Load("rcj.cfg",storymap)
    check(err)
    





    //
    
    
    
    
    
    var active int
    menu := []string{"Introduction", "Who is Sophia Roberts?", "Who is Giovanni Roberts?", "Ghostcrabs", "The Giant", "Nana's Nose", "Chicken James", "Roger's Mistakes", "Rubber Band Adventures", "Drinkin' Slosh", "There's a Degree In this Dumpster!", "Exit"}


    i := 0
	stdscr, _ := Init()
	defer End()

	Raw(true)
	Echo(false)
	Cursor(0)
	stdscr.Clear()
	stdscr.Keypad(true)

    
	rows, cols := stdscr.Maxyx()
	y, x := (((rows-HEIGHT)/2)+4), (cols-WIDTH)/2
    win, _ := NewWindow(HEIGHT, WIDTH, y, x)
	win.Keypad(true)
    stdscr.Print("\n o\"\"50;@V\a\n")
    stdscr.Print("                      ROBOT CHICKEN JAMES\n                         version 1.0")
     
    stdscr.Print("\n\n") 

    stdscr.Print("  \n                   VIRTUAL FAMILY HISTORIAN\n ")
	stdscr.Refresh()

	printmenu(&win, menu, active)

	for {
		ch := stdscr.GetChar()
		switch KeyString(ch) {
		case "q":
			return
		case "up":
			if active == 0 {
				active = len(menu) - 1
			} else {
				active -= 1
			}
		case "down":
			if active == len(menu)-1 {
				active = 0
			} else {
				active += 1
			}
	// The main action	
        case "enter":
			//stdscr.MovePrintf(23, 0, "Choice #%d: %s selected",
			////	active,
			////	menu[active])
            if active==11 {
            return
        }
            i++
            if i > 4 {

                
            i = 0 
            exec.Command("pkill","chrome").Run()


            }
            
            selection := strconv.Itoa(active)
            speechtext.SpeechText(storymap[selection])	
            
            // active is the number chosen
            // menu[active] is the menu entry
            // make a call to PlaySelection() 
            // here. 
            
            
            
            stdscr.ClearToEOL()
			stdscr.Refresh()
// 		default:
// 			stdscr.MovePrintf(23, 0, "Character pressed = %3d/%c",
// 				ch, ch)
// 			stdscr.ClearToEOL()
// 			stdscr.Refresh()
		}

		printmenu(&win, menu, active)
	}
}

func printmenu(w *Window, menu []string, active int) {
	y, x := 2, 2
	w.Box(0, 0)
	for i, s := range menu {
		if i == active {
			w.AttrOn(A_REVERSE)
			w.MovePrint(y+i, x, s)
			w.AttrOff(A_REVERSE)
		} else {
			w.MovePrint(y+i, x, s)
		}
	}
	w.Refresh()
}
