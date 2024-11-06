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

	var correctAnswers int = 0 

	for index := 0; index < 3; index++ {
		fmt.Printf("Problem %v: %v = ", index + 1, rows[index][0])
		
		var answer string
		fmt.Scan(&answer)

		if answer == rows[index][1] {
			correctAnswers++
		}
	}

	fmt.Printf("Correct answers: %v/%v\n", correctAnswers, len(rows))
}