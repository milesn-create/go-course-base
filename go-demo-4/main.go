package main

import (
	"demo/password/account"
	"demo/password/account/output"
	"demo/password/encrypter"
	"demo/password/files"
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleateAccount,
}
var menuVariants = []string{
	"1.Создать аккаунт",
	"2.Найти аккаунт по URL",
	"3.Найти аккаунт по логину",
	"4.Удалить аккаунт",
	"5.Выход\n",
	"Выберите вариант из меню",
}

// .env - файл с переменными окружения , обязательно добавляется в .gitignore  :
//( не пушится в репозиторий)
//	/.env*  - значит должно быть игнорировано полностью
//
// в файле .env  - можем прописывать заранее переменные используемые приложением, а не в ручную при заупске
// но прописав их в этом файле они не добавяться автоматически к остальным  переменным окружения
//  поэтому их предварительно стоит загрузить
// go get github.com/joho/godotenv
// после проверяем что в go.mod добавилась зависимость
// приводим зависимости в порядки с помощью  : go mod tidy
//.     — это как уборка в комнате: и так жить можно, но с порядком удобнее, безопаснее и приятнее. Д
// Даже если всё работало, теперь оно будет работать стабильнее и предсказуемее.

// Вообще хранить KEY шифрования в .env не очень безопасно
//
//	лучше просто передавать в командной строке , но для данногго проекта пойдет
func main() {
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось найти .env файл!")
	}
	fmt.Println("___________________Менеджер паролей_________________")
	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())

Menu:
	for {

		choice := promtData(menuVariants...)

		var menuFunc = menu[choice]

		if menuFunc == nil {
			break Menu
		}

		menuFunc(vault)

		// switch choice {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleateAccount(vault)
		// default:
		// 	break Menu
		// }
	}

}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promtData("Введите URL для поиска: ")

	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)

	})
	outputFinderAccounts(&accounts)

}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promtData("Введите логин для поиска: ")

	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)

	})
	outputFinderAccounts(&accounts)

}
func outputFinderAccounts(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунтов с данным логином не найдено")
	}
	for _, ak := range *accounts {
		ak.Output()
	}

}
func deleateAccount(vault *account.VaultWithDb) {
	url := promtData("Введите URL для удаления:")
	isDeleted := vault.DeleateAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено!!!")
	} else {
		output.PrintError("Не найдено!")
	}

}
func createAccount(vault *account.VaultWithDb) {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите URL: ")
	myAccount, err := account.NewAccount(login, password, url)

	if errors.Is(err, account.ErrInvalidURL) {

		output.PrintError("Неверный формат URL!!!")
		return
	}
	if errors.Is(err, account.ErrNoLogin) {
		output.PrintError("Обязательно введите логин!")

		return
	}

	vault.AddAccount(*myAccount)

}

func promtData(prompt ...string) string {
	for index, value := range prompt {
		if index == len(prompt)-1 {
			fmt.Printf("%v : ", value)

		} else {
			fmt.Println(value)

		}

	}
	var result string
	fmt.Scanln(&result)
	return result
}
