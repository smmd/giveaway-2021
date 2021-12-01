package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"encoding/json"
	"math/rand"

	"github.com/smmd/giveaway-2021/repository"
)

func main()  {
	var filePath string
	userNames := make([]string, 0)

	fmt.Print("Enter the file name: ")

	_, err := fmt.Scanf("%s", &filePath)

	if err != nil {
		panic(fmt.Errorf("could not read Path: %w", err))
	}

	file, err := os.Open(filePath)

	if err != nil {
		panic(fmt.Errorf("failed to open file error: %v", err))
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	for _, line := range text {
		userData := &repository.Response{}

		byt := []byte(line)
		err = json.Unmarshal(byt, userData)

		if err != nil {
			panic(fmt.Errorf("could not process line file error: %v", err))
		} else {
			instructions := userData.Data.FavoritersTimeline.Timeline.Instructions

			for _, instruction := range instructions {
				for _, entry := range instruction.Entries {
					user := entry.Content.ItemContent.UserResults.Result.Legacy

					if len(user.ScreenName) > 0 && user.FollowedBy {
						userNames = append(userNames, user.ScreenName)
					}
				}
			}
		}
	}

	followers := len(userNames)

	rand.Seed(time.Now().Unix())
	n := rand.Intn(followers)

	fmt.Printf("\nNumber of users that meet the rule: %v\n\n", len(userNames))
	fmt.Printf("Users: %#v\n\n", userNames)
	fmt.Printf("Winner: %v\n\n", userNames[n])
}
