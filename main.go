package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func getInput(c chan string) {
	for {
		answer := ""
		_, err := fmt.Scanf("%s", &answer)
		if err != nil {
			log.Panic(err)
		}
		c <- answer
	}
}

func main() {
	csvFile, err := os.Open("problems.csv")

	if err != nil {
		log.Panic(err)
	}

	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	csvFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	inputChan := make(chan string)
	go getInput(inputChan)

	correctAnswers := 0
Exit:
	for _, record := range records {
		fmt.Printf("Question: %s:", record[0])
		answer := ""

		timer := time.NewTimer(time.Duration(3 * time.Second))

		select {
		case <-timer.C:
			fmt.Printf("\nTimer expired!\n")
			break Exit
		case v := <-inputChan:
			answer = v
		}

		if answer == record[1] {
			correctAnswers++
		}
	}

	fmt.Printf("Score: %d / %d\n", correctAnswers, len(records))

}
