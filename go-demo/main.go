package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("\n~~~ Калькулятор индекса массы тела ~~~")
	for {

		userKG, userHight := getUserInput()
		IMT, err := clculateIMT(userKG, userHight)
		// IMT, _ := clculateIMT(userKG, userHight) если хотим проигнорить ошибку
		// обычно всегда печатаем ошибки
		if err != nil {
			fmt.Println(err)
			continue
		}
		//также есть метод panic() - показывает дполнительную отлажочную ифнормацию
		// резко завершает работу и выходит из программы с кодом больше 0
		// 	if err != nil {
		// 		panic("Не указан вес или высота")
		//	 }
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}

	}

}
func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("\nВведите ваш рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Printf("Введите свой вес : ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
func clculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("Не указан вес или высота")

	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil

}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела : %.0f", IMT)
	fmt.Printf(result)
	switch {
	case IMT < 16:
		fmt.Println("\nУ вас сильный дефицит массы тела.")

	case IMT < 18.5:
		fmt.Println("\nУ вас дефицит массы тела.")
	case IMT < 25:
		fmt.Println("\nУ вас масса тела в пределах нормы.")
	case IMT < 30:
		fmt.Println("\nУ вас 1-я степень ожирения.")
	case IMT < 35:
		fmt.Println("\nУ вас 1-я степень ожирения.")
	case IMT < 40:
		fmt.Println("\nУ вас 2-я степень ожирения.")
	default:
		fmt.Println("\nУ вас 3-я степень ожирения.")
	}

}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Вы хотите сделать еще рассчет? (y/n) ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true

	}
	return false

}

// if IMT < 16 {
// 		fmt.Println("\nУ вас сильный дефицит массы тела.")
// 	} else if IMT < 18.5 {

// 		fmt.Println("\nУ вас  дефицит массы тела.")
// 	} else if IMT < 25 {

// 		fmt.Println("\nУ вас масса тела в пределах нормы.")
// 	} else if IMT < 30 {

// 		fmt.Println("\nУ вас  избыточная масса тела.")
// 	} else if IMT < 35 {

// 		fmt.Println("\nУ вас 1-я степень ожирения.")
// 	} else if IMT <= 40 {

// 		fmt.Println("\nУ вас 2-я степень ожирения.")
// 	} else {

// 		fmt.Println("\nУ вас 3-я степень ожирения.")
// 	}
// 	fmt.Print(result)
