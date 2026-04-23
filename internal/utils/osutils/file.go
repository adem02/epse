package osutils

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"text/template"
)

func CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		errorMessage := fmt.Errorf("❌ error while creating file: %s", path)
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

func CreateFileFromTmpl(embeddedFS embed.FS, tmplFilePath, filePath string, data interface{}) error {
	if reflect.TypeOf(data).Kind() != reflect.Pointer {
		return fmt.Errorf("❌ data must be a pointer to a struct, got %T", data)
	}

	tmplContent, err := embeddedFS.ReadFile(tmplFilePath)
	if err != nil {
		return fmt.Errorf("❌ Failed to read embedded template %s: %w", tmplFilePath, err)
	}

	tmpl, err := template.New(filepath.Base(tmplFilePath)).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("❌ error parsing template: %s", tmplFilePath)
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
