package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3" // _ にしないとコンパイルエラーとなる（私用しないため _ ）
)

// DB Path (生徒の情報を記載しているデータベース)
// const dbPath = "/Users/admin/go/homework/line-bot"

// コネクションプールを作成
// var DbConnection *sql.DB

// 生徒のデータ格納用のユーザ定義型
type Student struct {
	Id   int64  // 生徒の ID
	Name string // 生徒の名前
	Time int    // 対象の生徒の学習制限時間
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	DbConnection, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return fmt.Errorf("データベースの Open : %w", err)
	}

	if err := createTable(DbConnection); err != nil {
		return fmt.Errorf("データベースの 作成 : %w", err)
	}

	students := []*Student{{Name: "郷 花子", Time: 5}, {Name: "郷 太郎", Time: 10}, {Name: "鈴宮 太郎", Time: 15}}

	if err := insertTable(DbConnection, students); err != nil {
		return fmt.Errorf("データベースへの情報追加 : %w", err)
	}

	scanTable(DbConnection)

	// defer DbConnection.Close()
	return nil
}

// Student テーブルの作成
func createTable(db *sql.DB) error {
	const sql = `CREATE TABLE IF NOT EXISTS student(
		Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL,
		Time INTEGER NOT NULL
	);`

	_, err := db.Exec(sql)
	if err != nil {
		return fmt.Errorf("テーブル作成: %w", err)
	}

	return nil
}

// Student テーブルに生徒情報を追加する
func insertTable(db *sql.DB, students []*Student) error {
	for i := range students {
		const sql = "INSERT INTO student(name, limit) VALUES (?, ?)"
		r, err := db.Exec(sql, students[i].Name, students[i].Time)
		if err != nil {
			return err
		}
		id, err := r.LastInsertId()
		if err != nil {
			return err
		}
		students[i].Id = id
	}
	return nil
}

// Student テーブルから情報をスキャンする
func scanTable(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM student WHERE Name = ?", os.Args)
	if err != nil {
		return fmt.Errorf("テーブルから名前を取得: %w", err)
	}

	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.Id, &s.Name, &s.Time); err != nil {
			return fmt.Errorf("データ追加: %w", err)
		}
		fmt.Println(s)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("エラー発生: %w", err)
	}

	return nil
}
