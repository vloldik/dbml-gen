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

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Генерирует модели из DBML файла",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile, _ := cmd.Flags().GetString("input")
		outputDir, _ := cmd.Flags().GetString("output")
		useGorm, _ := cmd.Flags().GetBool("gorm")

		dbml, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("Ошибка чтения входного файла: %v\n", err)
			os.Exit(1)
		}

		var parser IParser = dbparse.New()

		parsed, err := parser.Parse(string(dbml))
		j, _ := json.Marshal(parsed)
		println(string(j))
		if err != nil {
			fmt.Printf("Ошибка парсинга DBML: %v\n", err)
			os.Exit(1)
		}

		err = generator.GenerateModels(parsed, outputDir, useGorm)
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
	generateCmd.Flags().BoolP("gorm", "g", false, "Использовать теги GORM")
	generateCmd.MarkFlagRequired("input")
	rootCmd.AddCommand(generateCmd)
}
