package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"time"
	"sync"
	"flag"
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

var wg = sync.WaitGroup{}

func main() {

	csvFilename := flag.String("csv", "problems.csv", "a csv file with questions")
	flag.Parse()
	
	timerCh := make(chan bool) 
	answerCh := make(chan string)
	// запись и чтение канала блокирует функцию

	fmt.Println("main func")

	rows, err := readFile(*csvFilename)
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

	wg.Add(1)
	go startTimer(timerCh)
	wg.Add(1)
	go quiz(questions, timerCh, answerCh)
	
	wg.Wait()
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

func quiz(questions []Question, timerCh chan bool, answerCh chan string) {
	correctAnswers := 0

	for index, question := range questions {
		fmt.Printf("Problem %v: %v = ", index + 1, question.question)

		go readAnswer(answerCh)

		select {
			case <-timerCh:
				fmt.Println("Time is up")
				return
			case answer:=<-answerCh:
				if answer == question.answer {
					correctAnswers++
				}
		}
	}

	fmt.Printf("Correct answers: %v/%v\n", correctAnswers, len(questions))
	wg.Done()
}

// func quiz2(questions []Question, timerCh chan bool, answerCh chan string) {
// 	correctAnswers := 0

// 	for index, question := range questions {
// 		select {
// 			case <-timerCh:
// 				fmt.Println("Time is up")
// 				return
// 			case answer:=<-answerCh:
// 				fmt.Println("Answer", answerCh)
// 				if answer == question.answer {
// 					correctAnswers++
// 				}
	
// 			default:
// 				fmt.Printf("Problem %v: %v = ", index + 1, question.question)
// 				readAnswer(answerCh)	 // если оставить без go, то блокируется вызов, но тк внутри функции пишем в канал, а читателей нет - дедлок
										// если вызывать как горутину, то выводится несколько вопросов сразу, тк select не ждет ответа от пользака
										// если вынести вызов горутины за select, select будет ждать ответа (заблокированно) либо от таймера, либо от ответа
// 		}
// 	}

// 	fmt.Printf("Correct answers: %v/%v\n", correctAnswers, len(questions))
// 	wg.Done()
// }



func startTimer(ch chan bool) {
	// ch <- false 

	time.Sleep(5 * time.Second)
	ch <- true
	fmt.Println("timer finished")

	wg.Done()
}

func readAnswer(answerCh chan string) {
	var answer string
	fmt.Scan(&answer)

	answerCh <- answer
}