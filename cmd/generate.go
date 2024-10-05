package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"guthub.com/vloldik/dbml-gen/internal/dbparse"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/generator"
)

type IParser interface {
	Parse(dbml string) (*models.DBML, error)
}

const (
	DefaultBackend = ""
	GORMBackend    = "gorm"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Генерирует модели из DBML файла",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile, _ := cmd.Flags().GetString("input")
		outputDir, _ := cmd.Flags().GetString("output")
		useGorm, _ := cmd.Flags().GetString("backend")

		dbml, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("Ошибка чтения входного файла: %v\n", err)
			os.Exit(1)
		}

		var parser IParser = dbparse.New()

		parsed, err := parser.Parse(string(dbml))

		if err != nil {
			fmt.Printf("Ошибка парсинга DBML: %v\n", err)
			os.Exit(1)
		}

		j, err := json.Marshal(parsed)
		if err != nil {
			println(err.Error())
		}
		println(string(j))

		gen := generator.New()

		err = gen.GenerateModels(parsed, outputDir, useGorm)
		if err != nil {
			fmt.Printf("Ошибка генерации моделей: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Модели успешно сгенерированы")
	},
}

func init() {
	generateCmd.Flags().StringP("input", "i", "", "Путь к входному DBML файлу")
	generateCmd.Flags().StringP("output", "o", ".", "Директория для сгенерированных файлов")
	generateCmd.Flags().StringP("backend", "g", "", "Выбор функционала для сгенерированных моделей")
	generateCmd.MarkFlagRequired("input")
	rootCmd.AddCommand(generateCmd)
}
