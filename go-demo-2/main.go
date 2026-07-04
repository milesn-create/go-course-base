package main

import (
	"fmt"
)

func main() {
	tr := make([]string, 0, 2)
	// мы создали слайд и теперь програмам будет при каждом добавлении транзакции выделть память ,
	// чтобы это оптимизировать функцией make
	// предсоздав массив определенной ддлины и капесети
	// в данном случае длина 0,капасети 2,
	fmt.Println(len(tr), cap(tr)) //0 2
	tr = append(tr, "1")          //память не выделяется
	fmt.Println(len(tr), cap(tr)) //1 2
	tr = append(tr, "2")          //память не выделяется
	fmt.Println(len(tr), cap(tr)) //2 2
	tr = append(tr, "3")          //память выделяется, так как капесети 2 и мы добавляем 3 элемент
	fmt.Println(len(tr), cap(tr)) //3 4 - капесети удваивается когда выходим за граицу и если длина слайса меньше 256
	fmt.Println(tr)

	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~ Счетчик транзакций ~~~~~~~~~~~~~~~~~~~~~~~~~ ")
	fmt.Println("(когда транзакции законатся введите - '0' чтобы увидить итоговую сумму)\n\n ")
	transactions := []float64{}
	transactions, _ = inputTransactions(transactions)
	balance := calculateBalance(transactions)
	fmt.Printf("Ваш баланс: %.2f\n", balance)

}
func inputTransactions(transactions []float64) ([]float64, error) {

	for {
		fmt.Print("Введите транзакцию: ")
		simbol := 0.0

		fmt.Scan(&simbol)

		if simbol == 0 {
			return transactions, nil
		}
		transactions = append(transactions, simbol)

	}
}
func sumTrunsactions(transaction []float64) (float64, error) {
	var sumTrunsaction float64
	for i := 0; i < len(transaction); i++ {

		sumTrunsaction += transaction[i]
	}
	return sumTrunsaction, nil
}
func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
