package main

import (
	"fmt"
	"os"
	"encoding/csv"
)

/*

1. научиться читать csv файл
2. вытащить вопросы
3. задать вопрос, получить ответ от юзера
4. проверить ответ, записать результат

5. разобраться, как передавать флаги в параметры при запуске 

*/

type Question struct {
	question string
	answer string
}

func main() {
	file, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println("Error:", err)
	}

	defer file.Close() // executes before return in function

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	correctAnswers := 0

	var questions []Question

	for _, row := range rows {
		questions = append(questions, Question{
			question: row[0],
			answer: row[1],
		})
	}


	for index, question := range questions {
		fmt.Printf("Problem %v: %v = ", index + 1, question.question)
		
		var answer string
		fmt.Scan(&answer)

		if answer == question.answer {
			correctAnswers++
		}
	}

	fmt.Printf("Correct answers: %v/%v\n", correctAnswers, len(rows))
}