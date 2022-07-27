package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()

	// load config file containing the links
	home, _ := os.UserHomeDir()

	os.Mkdir(home+"/.config/gorss", 0755)
	file, err := os.OpenFile(home+"/.config/gorss/gorss", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// parse the config file
	scanner := bufio.NewScanner(file)
	index := 1
	for scanner.Scan() {
		// get all the title of each link and show as a list
		feed, err := fp.ParseURL(scanner.Text())
		if err != nil {
			panic(err)
		}
		fmt.Printf("[%v]: %s\n", index, feed.Title)
		index++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// make user choose what to view
	fmt.Printf("What would you like to view? ")
	var choice int
	fmt.Scanf("%2d", choice)

	// parse url and show content/entries

}