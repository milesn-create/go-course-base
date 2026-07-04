package files

import (
	"demo/password/account/output"
	"os"

	"github.com/fatih/color"
)

type JsonDBb struct {
	fileName string
}

func NewJsonDb(name string) *JsonDBb {
	return &JsonDBb{
		fileName: name,
	}

}
func (db JsonDBb) Read() ([]byte, error) {

	data, error := os.ReadFile(db.fileName)

	if error != nil {

		return nil, error

	}
	return data, nil

}

func (db JsonDBb) Write(content []byte) {
	file, error := os.Create(db.fileName)
	if error != nil {
		output.PrintError(error)

	}

	_, error = file.Write(content)

	defer file.Close()

	if error != nil {
		output.PrintError(error)

	}
	color.Green("Файл успешно записан!")

}
