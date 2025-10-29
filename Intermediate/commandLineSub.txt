package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	stringFlag := flag.String("User", "Guest", "Name of the coder")
	flag.Parse()
	fmt.Println(stringFlag)

	subCommand1 := flag.NewFlagSet("info", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("lang", flag.ExitOnError)

	firstFlag := subCommand1.Bool("Processing", false, "Command processing status.")
	secondFlag := subCommand1.Int("Bytes", 1024, "Byte length of the result")

	flagsc2 := subCommand2.String("Language", "Go", "enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "info":
		subCommand1.Parse(os.Args[2:])
		fmt.Println("Processing:", *firstFlag)
		fmt.Println("Bytes:", *secondFlag)
	case "lang":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("Language:", *flagsc2)
	default:
		fmt.Println("NO valid commands are found")
		os.Exit(1)
	}
}
