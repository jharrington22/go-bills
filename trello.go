package main

import (
	"fmt"
	"log"

	"github.com/VojtechVitek/go-trello"
)

func main() {
	// New Trello Client
	appKey := ""
	token := ""

	trello, err := trello.NewAuthClient(appKey, &token)
	if err != nil {
		log.Fatal(err)
	}

	// User @trello
	user, err := trello.Member("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.FullName)

	// @trello Boards
	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}

	if len(boards) > 0 {
		board := boards[0]
		fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)

		// @trello Board Lists
		lists, err := board.Lists()
		if err != nil {
			log.Fatal(err)
		}

		for _, list := range lists {
			fmt.Println("   - ", list.Name)

			// @trello Board List Cards
			cards, _ := list.Cards()
			for _, card := range cards {
				fmt.Println("      + ", card.Name)
			}
		}
	}
}
