package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vloldik/dbml-gen/internal/dbparse"
	"github.com/vloldik/dbml-gen/internal/dbparse/models"
	"github.com/vloldik/dbml-gen/internal/generator"
	"github.com/vloldik/dbml-gen/internal/generator/gormgen"
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
		tagStyle, _ := cmd.Flags().GetString("backend")
		module, _ := cmd.Flags().GetString("module")

		var structGen generator.IStructFromTableGenerator

		switch tagStyle {
		case "gorm":
			structGen = gormgen.NewStructGenerator()
		default:
			panic(fmt.Errorf("unsupported backend: %s", tagStyle))
		}

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

		gen := generator.New(outputDir, module, tagStyle, structGen)

		err = gen.GenerateModels(parsed)
		if err != nil {
			fmt.Printf("Ошибка генерации моделей: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Модели успешно сгенерированы")
	},
}

func init() {
	generateCmd.Flags().StringP("input", "i", "", "Путь к входному DBML файлу")
	generateCmd.Flags().StringP("module", "m", "", "Название модуля (e.g. gorm.io/gorm)")
	generateCmd.Flags().StringP("output", "o", ".", "Директория для сгенерированных файлов")
	generateCmd.Flags().StringP("backend", "g", "", "Выбор функционала для сгенерированных моделей")
	generateCmd.MarkFlagRequired("input")
	rootCmd.AddCommand(generateCmd)
}
