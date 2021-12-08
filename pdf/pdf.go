package pdf

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"htmlToPdf/common"
	"os"
)

func HtmlToPdf(htmlFile string) (string, error) {
	root := common.HtmlStruct{RootPath: "file"}

	f, err := os.Open(htmlFile)
	if err != nil {
		return "", err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return "", err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	if err := pdfg.Create(); err != nil {
		return "", err
	}

	fileName := root.RootPath + "/genie.pdf"

	if err := pdfg.WriteFile(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}
