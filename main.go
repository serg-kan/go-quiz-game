package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"time"
	// "sync"
)

/*
	Part 1: 

	1. научиться читать csv файл
	2. вытащить вопросы
	3. задать вопрос, получить ответ от юзера
	4. проверить ответ, записать результат

	5. разобраться, как передавать флаги в параметры при запуске 

*/
/*
	Part 2: 

	1. запустить таймер в начале игры
	2. задавать вопросы не обращая внимания на таймер
	3. по окончании таймера выключать программу
*/
/*
	Refactor

	1. read csv in a separate function 
	2. 
*/

type Question struct {
	question string
	answer string
}

// var wg = sync.WaitGroup{}

func main() {
	
	timerCh := make(chan bool) 
	// запись и чтение канала блокирует функцию

	
	fmt.Println("main func")

	rows, err := readFile("problems.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var questions []Question
	for _, row := range rows {
		questions = append(questions, Question{
			question: row[0],
			answer: row[1],
		})
	}

	// wg.Add(1)/
	go startTimer(timerCh)

	timerFlag := false
	correctAnswers := 0


	for index, question := range questions {
		if timerFlag {
			break
		}

		select {
			case <-timerCh:
				fmt.Println("Time is up")
				timerFlag = true
			
			default:
				fmt.Printf("Problem %v: %v = ", index + 1, question.question)
				
				var answer string
				fmt.Scan(&answer)

				if answer == question.answer {
					correctAnswers++
				}
		}
	}

	fmt.Printf("Correct answers: %v/%v\n", correctAnswers, len(rows))

	// wg.Wait()
}

func readFile(name string) ([][]string, error) {
	file, err := os.Open(name)

	if err != nil {
		fmt.Println("Error:", err)
	}

	defer file.Close() // executes before return in function

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func startTimer(ch chan bool) {
	// ch <- false 

	time.Sleep(5 * time.Second)
	ch <- true
	fmt.Println("timer finished")

	// wg.Done()
}
