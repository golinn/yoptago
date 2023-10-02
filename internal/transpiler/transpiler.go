package transpiler

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kvizyx/yoptago/internal/words"
)

const yoptaSuffix = ".yo"

func TranspileDirectory(path string) ([]string, error) {
	var transpiled []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), yoptaSuffix) {
			body, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			transpiledFile := transpileFile(string(body))

			tempFilePath := strings.TrimSuffix(path, yoptaSuffix) + ".go"
			err = os.WriteFile(tempFilePath, []byte(transpiledFile), 0644)
			if err != nil {
				return err
			}

			transpiled = append(transpiled, tempFilePath)
		}
		return nil
	})

	return transpiled, err
}

func transpileFile(body string) string {
	for original, poPatsanski := range words.Words {
		body = strings.ReplaceAll(body, poPatsanski, original)
	}

	return body
}
