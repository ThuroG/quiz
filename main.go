package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type QuizItem struct {
	input  string
	result int
}

func createQuiz(data [][]string) []QuizItem {
	var quiz []QuizItem
	for i, line := range data {
		i = i
		var rec QuizItem
		for j, field := range line {
			if j == 0 {
				rec.input = field
			} else if j == 1 {
				n, err := strconv.Atoi(field)
				if err != nil{
					fmt.Println(err)
				} else{
					rec.result = n
				}
			}
		}
		quiz = append(quiz, rec)
	}
	return quiz
}

func readCsvFile(csvFile string) [][]string{
	//open file
	//f, err := os.Open("problems.csv") //TODO: use csvFile input variable
	f, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}

	//close file at the end of the program - defer used when surrounding function returns stgh
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the CSV file: ")
	csvFile, _ := reader.ReadString('\n')
	fmt.Print("Read CSV File " + csvFile)
	//trim Input because windows terminater suckssssssss \r\n
	csvFile = strings.TrimSuffix(csvFile, "\n")
	csvFile = strings.TrimSuffix(csvFile, "\r")

	csv := readCsvFile(csvFile)
	//convert records to array of Quiz structs
	quiz := createQuiz(csv)

	//set counter of correct answered to 0
	var answeredCorrectInt = 0
	for i, question := range quiz{
		var n = i+1
		fmt.Print(strconv.Itoa(n) + ". Question: What is " + question.input + "?")
		var answer int
		_, err := fmt.Scanf("%d\n", &answer)
		if err != nil {
			log.Fatal(err)
		}
		if ( answer == question.result){
			answeredCorrectInt = answeredCorrectInt + 1
		}
	} 
	
	answeredCorrectStr := strconv.Itoa(answeredCorrectInt)
	fmt.Print("You had " + answeredCorrectStr + " Questions correct!")
	//debug quiz
	//fmt.Printf("%+v\n", quiz)

}
