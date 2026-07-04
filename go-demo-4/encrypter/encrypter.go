package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр KEY в переменные окружения")
	}
	return &Encrypter{
		Key: key,
	}
}
func (enc *Encrypter) Encrypt(plainStr []byte) []byte {
	//будем использовать симетричное шифрование  - библиотека aes
	// aes.NewCipher  - принимает слайс байт поэтому обернем наш ключ в слайс байт
	// возвращает эта функция два значения :
	// 		1. Block - объекь представляющий симитричный блочный шифр
	// 		2.  err - ошибка если что-то вдруг не удалось
	block, err := aes.NewCipher([]byte(enc.Key)) // создает по ключу уникальный шифр  -
	// то есть смешивает ключ с разными алгоритмами и получается зашифрованый ключ
	if err != nil {
		panic(err.Error())
	}
	//gcm - разбирает данные на кусочки по 16 байт и шифрует каждый отдельно
	//Нюанс: GCM не просто разбивает на кусочки. Он использует режим счетчика (CTR):
	// Каждый блок (16 байт) шифруется отдельно с помощью AES
	// Но для каждого блока используется уникальный счетчик (counter), который меняется от блока к блоку
	// Это позволяет параллельно шифровать блоки (очень быстро!)
	aesGCM, err := cipher.NewGCM(block) // прверяется также аутентификация - проверка целостности
	if err != nil {
		panic(err.Error())
	}
	//      в nonce  хранится некоторое уникальное значение, которое мы можем использовать для шифрования
	//			- Это нужно, чтобы одинаковые сообщения давали разный зашифрованный результат!
	// 			- для gcm  nonceSize = 12 байт
	//       	- nonce = Number used ONCE (число, используемое один раз)
	//       	- Случайное уникальное значение, которое НИКОГДА НЕ ПОВТОРЯЕТСЯ для одного ключа
	nonce := make([]byte, aesGCM.NonceSize())
	//io.ReadFull читает ровно len(nonce) байт (12 байт) из rand.Reader
	// Возвращает: (количество_прочитанных_байт, ошибка)
	// Количество байт игнорируем, т.к. если ошибки нет, то прочитано ровно 12
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	//возвращается зашифрованные слайс байт
	return aesGCM.Seal(nonce, nonce, plainStr, nil)
	//метод seal - шифрует
	// первыей раз nonce - в качестве dst - destination - место назначения
	// , в конечном итоге зашифрованные данные будут начинасться с nonce  - потом уже сами данные с
	// втрой раз - nonce чтобы использовать его для шифрования, последний аргумент доп данные например подпись - сейчс без них

}

// оба процессы похожи
// 1. создаем шифровальный блок по ключу
// 2. потом  создаем gcm
// 3. с помощью gcm можем или зашифровать ( aesGCM.Seal) или ращифровать( aesGCM.Open),
// предварительно понадобить поработаь с nonce (дополнительной случайной примесью )
func (enc *Encrypter) Decrypt(encryptedStr []byte) []byte {
	// 1. Создаем блочный шифр по ключу
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	// 2. Создаем GCM-режим для аутентифицированного расшифрования
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	// 3. Узнаем размер nonce (для GCM это 12 байт)
	nonceSize := aesGCM.NonceSize()
	// 4. Разделяем зашифрованные данные на nonce и сам шифр
	nonce, cipherText := encryptedStr[:nonceSize], encryptedStr[nonceSize:]
	// 5. Расшифровываем
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	// 6. Возвращаем расшифрованные данные
	return plainText

}
