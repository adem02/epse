package osutils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"text/template"
)

func CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		errorMessage := fmt.Errorf("❌ erreur lors de la création du fichier: %s", path)
		return nil, errorMessage
	}

	return file, nil
}

func OpenFileWithWriteMode(path string, writeMode bool) (*os.File, error) {
	var file *os.File
	var err error

	if writeMode {
		file, err = os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			err = fmt.Errorf("❌ failed to open file in write mode: %s", path)
		}
	} else {
		file, err = os.Open(path)
		if err != nil {
			err = fmt.Errorf("❌ failed to open file in read mode: %s", path)
		}
	}

	return file, err
}

func WriteJSONToFile(file *os.File, data interface{}) error {
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("❌ error while encoding the JSON file: %s", file.Name())
	}

	return nil
}

func CreateFileFromTmpl(tmplFilePath, filePath string, data interface{}) error {
	if reflect.TypeOf(data).Kind() != reflect.Pointer {
		return fmt.Errorf("❌ data must be a pointer to a struct, got %T", data)
	}

	tmpl, err := template.ParseFiles(tmplFilePath)
	if err != nil {
		errorMessage := fmt.Errorf("❌ error parsing template: %s", tmplFilePath)
		return errorMessage
	}

	newFile, err := CreateFile(filePath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	return tmpl.Execute(newFile, data)
}

func ParseJSONToStruct(reader io.Reader, data interface{}) error {
	if reflect.TypeOf(data).Kind() != reflect.Pointer {
		return fmt.Errorf("❌ data must be a pointer to a struct, got %T", data)
	}

	return json.NewDecoder(reader).Decode(data)
}

func AppendToFile(path, content string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("❌ failed to open file for appending: %s", path)
	}
	defer file.Close()

	_, err = fmt.Fprint(file, content)
	if err != nil {
		return fmt.Errorf("❌ failed to append to file: %s", path)
	}

	return nil
}
