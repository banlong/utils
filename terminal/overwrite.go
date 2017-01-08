package main

import (
	"fmt"
	"time"
	"github.com/jroimartin/gocui"
	"log"
)

func main() {
	//showCount()
	showMultipleCount()
	//HelloWorldBox()
}

func showCount()  {
	ticker := time.Tick(time.Second)
	fmt.Println("Counting down to launch...")
	for i := 1; i <= 10; i++ {
		<-ticker
		fmt.Printf("\rOn %d/10", i)   // use \r if you are running this in terminal
		//fmt.Printf("\x0cOn %d/10", i) // use \x0c for play.golang.org
	}
	fmt.Println("\nAll is said and done.")
}

func showMultipleCount(){
	fmt.Println("Counting down to launch...")
	ticker := time.Tick(time.Second)
	var i, j, total *int
	i = new(int)
	j = new(int)
	total = new(int)
	for ;*total < 20; {
		go func(A *int, B *int) {
			for ; *B <= 10;*B++ {
				<-ticker
				*total = *A + *B
				fmt.Printf("\rTotal: %d/20 , A:%d/10, B: %d/10", *total, *A, *B)
			}
		}(i, j)

		for ; *i <= 10; *i++{
			<-ticker
			*total = *i + *j
			fmt.Printf("\rTotal: %d/20 , A:%d/10, B: %d/10", *total, *i, *j)

		}
	}
	fmt.Println("\nAll is said and done.")
}

func HelloWorldBox() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+27, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world! This is a test")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}