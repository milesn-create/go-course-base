package main

import "fmt"

// type Alias - позваляет улучшить читабельность кода
type bookmarkMap = map[string]string

func main() {

	bookmarks := bookmarkMap{}

Menu: // создаем лейбл чтобы выйти из цикла for c помощью break Menu
	for {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("Меню:\n1.Посмотреть закладки\n2.Добавить закладку\n3.Удалить закладку\n4.Завершение")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

		choice := 0
		fmt.Print("Введите опцию из меню: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:

			show(bookmarks)

		case 2:
			var nameNewObject string
			fmt.Print("Введите название закладки которую хотите добаваить: ")
			fmt.Scanln(&nameNewObject)
			var addressNewObject string
			fmt.Print("Введите её адресс: ")
			fmt.Scanln(&addressNewObject)

			App(nameNewObject, addressNewObject, bookmarks)

		case 3:

			var nameObject string
			fmt.Print("Введите название закладки которую хотите удалить: ")
			fmt.Scanln(&nameObject)
			deleteObject(nameObject, bookmarks)

		case 4:
			break Menu

		}

	}

}

func show(m bookmarkMap) {
	if len(m) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for key, value := range m {
		fmt.Println(key, ":", value)
	}

}
func App(name, adress string, m bookmarkMap) {
	m[name] = adress

}
func deleteObject(name string, m bookmarkMap) {
	delete(m, name)

}
