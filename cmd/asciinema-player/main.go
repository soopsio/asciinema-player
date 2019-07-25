package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/soopsio/asciinema-player/pkg/asciicast"
	"github.com/soopsio/asciinema-player/pkg/parser"
	"github.com/soopsio/asciinema-player/pkg/terminal"
	goterm "golang.org/x/crypto/ssh/terminal"
)

func errExit(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var (
	maxWait  time.Duration
	speed    float64
	filePath string
)

func init() {
	flag.DurationVar(&maxWait, "maxWait", 2*time.Second, "maximum time between frames")
	flag.Float64Var(&speed, "speed", 1, "speed adjustment: <1 - increase, >1 - decrease")
	flag.StringVar(&filePath, "f", "", "path to asciinema v2 file")
	flag.Parse()
}

func main() {
	if filePath == "" {
		fmt.Println("Please specify file\nUsage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Open(filePath)
	errExit(err)
	defer file.Close()

	parsed, err := parser.Parse(file)
	errExit(err)

	f, _ := os.Open("/dev/tty")
	state, _ := goterm.GetState(int(f.Fd()))
	defer goterm.Restore(int(f.Fd()), state)

	term, err := terminal.NewPty()
	errExit(err)
	errExit(term.ToRaw())
	defer term.Reset()
	//log.Println("abcd")
	tp := &asciicast.TerminalPlayer{Terminal: term}
	//log.Println("efgh")

	err = tp.Play(parsed, maxWait, speed)
	errExit(err)
}
