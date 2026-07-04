package account

import (
	"demo/password/account/output"
	"demo/password/encrypter"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}
type ByteWriter interface {
	Write([]byte)
}
type Db interface {
	ByteReader
	ByteWriter
}
type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type VaultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now()},
			db:  db,
			enc: enc,
		}

	}
	data := enc.Decrypt(file)
	var vault Vault

	err = json.Unmarshal(data, &vault)
	color.Cyan("Найдено %d аккаунтов", len(vault.Accounts))
	if err != nil {
		output.PrintError("Не удалось разобрать файл data.vault")

		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now()},
			db:  db,
			enc: enc,
		}

	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}

}
func (vault *VaultWithDb) AddAccount(ak Account) {
	vault.Accounts = append(vault.Accounts, ak)
	vault.save()

}
func (vault *VaultWithDb) DeleateAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, ak := range vault.Accounts {
		isMatched := strings.Contains(ak.Url, url)
		if !isMatched {

			accounts = append(accounts, ak)
			continue

		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.save()
	return isDeleted

}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, ak := range vault.Accounts {

		isMatched := checker(ak, str)
		if isMatched {
			accounts = append(accounts, ak)

		}
	}
	return accounts

}
func (vault *Vault) ToBytes() ([]byte, error) {
	file, error := json.Marshal(vault)
	if error != nil {
		return nil, error
	}
	return file, nil
}
func (vault *VaultWithDb) save() {

	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	encData := vault.enc.Encrypt(data)
	if err != nil {
		output.PrintError("Не удалось преобразовать файл data.json")

	}
	vault.db.Write(encData)

}

func sum[T int | string](a, b T) T {
	switch t := any(a).(type) {
	case string:
		fmt.Println(t)

	}
	return a + b
}

type List[T any] struct {
	elements []T
}

func (l *List[T]) addElement() {}
