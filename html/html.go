package html

import (
	"html/template"
	"htmlToPdf/common"
	"os"
)

func HtmlTemplate(filePath string, data interface{}) (string, error) {
	root := common.HtmlStruct{RootPath: "file"}

	t, err := template.ParseFiles(filePath)
	if err != nil {
		return "", err
	}

	fileName := root.RootPath + "/genie.html"

	fileWriter, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	if err := t.Execute(fileWriter, data); err != nil {
		return "", err
	}

	return fileName, nil
}
