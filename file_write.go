package main

import (
	"log"
	"os"
)

func main() {
	// ファイルを生成
	file, err := os.Create("./file.txt")
	if err != nil {
		// エラー処理
		log.Fatal(err)
	}

	//書き込むデータを用意する

	//Write()を使って書き込む
	_, err = file.WriteString("Hello, world!\n")
	if err != nil {
		log.Fatal(err)
	}
}
