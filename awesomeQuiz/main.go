package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func load_problems(filename string) []problem {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	data, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	problems := make([]problem, len(data))

	for i, row := range data {
		problems[i] = problem{row[0], row[1]}
		// row[0] = czy kot, row[1] = nie
		// problems[5] = problem {"czy kot", "nie"}
	}

	return problems
}

func ask_user(i int, p problem) bool {
	fmt.Printf("\n%v. %v\n", i+1, p.question)

	var user_answer string
	_, err := fmt.Scan(&user_answer)

	if err != nil {
		log.Fatal(err)
	}

	return strings.ToLower(p.answer) == strings.ToLower(user_answer)
}

const QUIZ_TIME_S = 30
const NUMBER_OF_QUESTIONS = 10

func main() {
	// csv - pyt i odp, counter- R,
	problems := load_problems("problems.csv")

	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})

	fmt.Printf("Na quiz masz %v sekund. Jest %v pytan.\n", QUIZ_TIME_S, NUMBER_OF_QUESTIONS)
	fmt.Println("Nacisnij enter by zaczac...")
	fmt.Scanln()

	start_time := time.Now()
	end_time := start_time.Add(QUIZ_TIME_S * time.Second)

	true_counter := 0

time_loop:
	for i, p := range problems[:NUMBER_OF_QUESTIONS] {
		fmt.Printf("Zostalo ci %.0f sekund\n", end_time.Sub(time.Now()).Seconds())

		is_answered := false
		answer_ok := false

		go func() {
			answer_ok = ask_user(i, p)
			is_answered = true
		}()

		for !is_answered {
			if end_time.Before(time.Now()) {
				fmt.Println("Czas dobiegl konca!")
				break time_loop
			}
		}

		if answer_ok {
			true_counter++
		}
	}

	fmt.Printf("Brawo! Wynik to: %v / %v", true_counter, NUMBER_OF_QUESTIONS)
}
