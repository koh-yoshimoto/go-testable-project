package testutils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	_ "github.com/go-sql-driver/mysql"
)

// LoadFixtures は `./fixtures/testdata/*.yml` の全データを MySQL にロードする
func LoadFixtures(db *sql.DB, dirPath string) error {
	// 指定ディレクトリの全 YAML ファイルを取得
	files, err := filepath.Glob(filepath.Join(dirPath, "*.yml"))
	log.Printf("Loading fixtures from %v", files)
	if err != nil {
		return fmt.Errorf("failed to find fixture files: %w", err)
	}

	// YAML ファイルごとに処理
	for _, file := range files {
		log.Printf("Loading fixture: %s", file)

		// YAML ファイルを読み込む
		data, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read YAML file: %w", err)
		}

		// YAML を動的に解析
		var content map[string][]map[string]interface{}
		if err := yaml.Unmarshal(data, &content); err != nil {
			return fmt.Errorf("failed to parse YAML: %w", err)
		}

		// 各テーブルのデータを挿入
		for table, records := range content {
			// 既存データを削除
			_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", table))
			if err != nil {
				return fmt.Errorf("failed to clear table %s: %w", table, err)
			}

			// データを挿入
			for _, record := range records {
				if err := insertRecord(db, table, record); err != nil {
					return fmt.Errorf("failed to insert record into %s: %w", table, err)
				}
			}
		}
	}

	log.Println("All fixtures loaded successfully")
	return nil
}

// insertRecord は動的なカラムと値で INSERT 文を構築し、MySQL にデータを挿入する
func insertRecord(db *sql.DB, table string, record map[string]interface{}) error {
	columns := []string{}
	values := []interface{}{}
	placeholders := []string{}

	// カラム名と値を取得
	for column, value := range record {
		columns = append(columns, column)
		values = append(values, value)
		placeholders = append(placeholders, "?")
	}

	// INSERT クエリを組み立て
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		"`"+join(columns, "`, `")+"`", // カラム名をバッククォートで囲む
		join(placeholders, ", "),
	)

	// クエリ実行
	_, err := db.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("insert error: %w", err)
	}

	return nil
}

// join はスライスを指定のセパレータで結合するユーティリティ関数
func join(strs []string, sep string) string {
	result := ""
	for i, str := range strs {
		if i > 0 {
			result += sep
		}
		result += str
	}
	return result
}

