package main

import (
	"bufio"
	"fmt"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
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
	var feeds []*gofeed.Feed
	for scanner.Scan() {
		// get all the title of each link and show as a list
		feed, err := fp.ParseURL(scanner.Text())
		feeds = append(feeds, feed)
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
	choice := 0
	fmt.Scanln(&choice)

	// parse url and show list of items
	feed := feeds[choice-1]
	fmt.Printf("\n%s\n", feed.Title)
	for i, item := range feed.Items {
		fmt.Println(i, item.Title)
	}

	// make user choose what to view
	fmt.Printf("What would you like to view? ")
	fmt.Scanln(&choice)

	// show item contents
	desc := markdown.Render(feed.Items[choice].Description, 80, 6)
	fmt.Printf("\n%v\n\n%v\n%v\n", feed.Items[choice].Title, string(desc), feed.Items[choice].Link)
}