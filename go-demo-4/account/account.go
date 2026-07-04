package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// метод для генерации пароля
func (ak *Account) generatePassword(n int) {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	ak.Password = string(result)
}

// метод для вывода структуры в консоль
func (ak *Account) Output() {
	color.Red("~~~~~~Данные вашего аккаунта~~~~~~")
	color.Cyan("Логин: %s\n", ak.Login)
	color.Cyan("Ваш Пароль: %s\n", ak.Password)
	color.Cyan("Ваш Url: %s\n", ak.Url)
	color.Cyan("Время создания: %s\n", ak.CreatedAt.Format("02.03.2006 15:04:05"))
	color.Cyan("Время обновления: %s\n", ak.UpdatedAt.Format("02.03.2006 15:04:05"))

}

// массив допустипых симоволов для генерации пароля
var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!_*")

// создадим переменые для ошибок
var (
	ErrInvalidURL = errors.New("Invalid URL")
	ErrNoLogin    = errors.New("Isn`t login!!!")
)

// функция для создания расширенной структуры
func NewAccount(login, password, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, ErrInvalidURL
	}
	if login == "" {
		return nil, ErrNoLogin
	}

	newAccount := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Password:  password,
		Url:       urlString,
	}

	if password == "" {
		newAccount.generatePassword(12)

	}

	return newAccount, nil

}
