package global

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

func GenerateFile(filePath string, fileName string, tplPath string) error {
	err := os.MkdirAll(filePath, 0777)
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	mainFile, err := os.Create(filePath + "/" + fileName)
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	defer mainFile.Close()

	mainTemp, err := template.ParseFiles(tplPath)
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}

	var buf bytes.Buffer

	mainTemp.Execute(&buf, BootConfigYaml)

	_, err = mainFile.Write(buf.Bytes())
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	fmt.Printf("successfully generated: %v/%v\n", filePath, fileName)
	return nil
}

func GenerateText(filePath string, fileName string, tplText string) error {
	err := os.MkdirAll(filePath, 0777)
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	mainFile, err := os.Create(filePath + "/" + fileName)
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	defer mainFile.Close()

	var buf bytes.Buffer

	template.Must(template.New("").Parse(tplText)).Execute(&buf, BootConfigYaml)

	_, err = mainFile.Write(buf.Bytes())
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	fmt.Printf("successfully generated: %v/%v\n", filePath, fileName)
	return nil
}

func GenerateFileForData(filePath string, fileName string, tplPath string, data any) error {
	err := os.MkdirAll(filePath, 0777)
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	mainFile, err := os.Create(filePath + "/" + fileName)
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	defer mainFile.Close()

	mainTemp, err := template.ParseFiles(tplPath)
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	var buf bytes.Buffer

	mainTemp.Execute(&buf, data)

	_, err = mainFile.Write(buf.Bytes())
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	fmt.Printf("successfully generated: %v/%v\n", filePath, fileName)
	return nil
}

func GenerateTextForData(filePath string, fileName string, tplText string, data any) error {
	err := os.MkdirAll(filePath, 0777)
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	mainFile, err := os.Create(filePath + "/" + fileName)
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	defer mainFile.Close()

	var buf bytes.Buffer

	template.Must(template.New("").Parse(tplText)).Execute(&buf, data)

	_, err = mainFile.Write(buf.Bytes())
	if err != nil {
		fmt.Printf("generated: %v/%v is err:%v\n", filePath, fileName, err)
		return err
	}
	fmt.Printf("successfully generated: %v/%v\n", filePath, fileName)
	return nil
}
