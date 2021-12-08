package main

import (
	"fmt"
	"htmlToPdf/common"
	"htmlToPdf/html"
	"htmlToPdf/pdf"
	"os"
)

func main() {

	dataHTML := common.Data{
		Name:    "YooGenie",
		Address: "서울",
		Mobile:  "000-0000-0000",
	}

	fileHtmlName, err := html.HtmlTemplate("template/test.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("HTML : ", fileHtmlName)

	defer os.Remove(fileHtmlName)
	filePDFName, err := pdf.HtmlToPdf(fileHtmlName)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("PDF : ", filePDFName)
	return
}
