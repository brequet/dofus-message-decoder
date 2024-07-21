package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/brequet/dofus-data-file-parser/pkg/generator"
	"github.com/brequet/dofus-data-file-parser/pkg/parser"
)

// purpose: generate golang types from dofus .d2o files
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "dofusDataFolderPath outputFolderPath")
		os.Exit(1)
	}

	dofusDataFolderPath := os.Args[1]
	outputFolderPath := os.Args[2]

	err := checkDofusDataFolder(dofusDataFolderPath)
	if err != nil {
		slog.Error("error with provided dofus data folder", "error", err)
		os.Exit(1)
	}

	err = os.MkdirAll(outputFolderPath, 0755)
	if err != nil {
		slog.Error("error creating output folder", "error", err)
		os.Exit(1)
	}

	err = processCommonFolder(filepath.Join(dofusDataFolderPath, "common"), outputFolderPath)
	if err != nil {
		slog.Error("error processing common folder", "error", err)
		os.Exit(1)
	}
}

func checkDofusDataFolder(dofusDataFolderPath string) error {
	err := checkFolderExists(dofusDataFolderPath)
	if err != nil {
		return err
	}

	err = checkFolderExists(filepath.Join(dofusDataFolderPath, "common"))
	if err != nil {
		return err
	}

	return nil
}

func checkFolderExists(folderPath string) error {
	folderInfo, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("folder does not exist: %w", err)
	}

	if !folderInfo.IsDir() {
		return fmt.Errorf("folder is not a directory: %w", err)
	}

	return nil
}

func processCommonFolder(commonFolderPath, outputFolderPath string) error {
	files, err := os.ReadDir(commonFolderPath)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	classes := map[string]map[string]parser.Class{}

	fileParsedCount := 0
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".d2o" {
			continue
		}

		d2oFilePath := filepath.Join(commonFolderPath, file.Name())
		data, err := parser.ProcessD2oFile(d2oFilePath)
		if err != nil {
			slog.Error("error parsing file", "error", err)
			continue
		}

		fileParsedCount++

		for _, class := range data.Classes {
			if classes[class.PackageName] == nil {
				classes[class.PackageName] = map[string]parser.Class{}
			}
			classes[class.PackageName][class.PackageClass] = class
		}
	}
	slog.Info("d2o files parsed", "count", fileParsedCount)

	err = exportClassTypesToGolang(classes, outputFolderPath)
	if err != nil {
		return fmt.Errorf("error exporting class types to golang: %w", err)
	}

	return nil
}

func exportClassTypesToGolang(classes map[string]map[string]parser.Class, outputFolderPath string) error {
	for packageName, classMap := range classes {

		classList := make([]parser.Class, 0)
		for _, class := range classMap {
			classList = append(classList, class)
		}

		goFileContent, err := generator.GenerateGoFromClasses(classList)
		if err != nil {
			return fmt.Errorf("error generating golang from classes: %w", err)
		}

		fileName := packageName[strings.LastIndex(packageName, ".")+1:] + ".go"

		goFilePath := filepath.Join(outputFolderPath, fileName)
		err = os.WriteFile(goFilePath, goFileContent, 0644)
		if err != nil {
			return fmt.Errorf("error writing file: %w", err)
		}
	}

	return nil
}
