package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Options struct {
	OrganizeType *string
	LetterNum    *int
	Pattern      *string
	IsDryrun     *bool
	IsVerbose    *bool
}

func main() {
	options := parseArgs()

	targetDirectories := flag.Args()

	printf(*options.IsVerbose, "整頓対象ディレクトリ: %s\n", strings.Join(targetDirectories, ", "))
	printf(*options.IsVerbose, "整頓タイプ: %s\n", *options.OrganizeType)

	for _, targetDir := range targetDirectories {

		// 絶対パスに変換
		targetDir, err := filepath.Abs(targetDir)
		printf(*options.IsVerbose, "ディレクトリの整頓を開始: %s\n", targetDir)

		files, err := os.ReadDir(targetDir)
		if err != nil {
			message := fmt.Sprintf("ディレクトリの読み込みに失敗しました: %s", err.Error())
			panic(message)
		}

		// ディレクトリ内のファイルをひとつずつ走査
		// 1. ファイル名の先頭 `letter_num` を抜き出して、ディレクトリを作成
		//    - ディレクトリが既に存在する場合は何もしない
		// 2. 対象ファイルを「1.」で作成した(または、作成する予定だった)ディレクトリへ移動
		for _, file := range files {

			// 絶対パスへ変換
			fileName := file.Name()
			fileAbsPath := filepath.Join(targetDir, fileName)

			// ディレクトリは無視
			if file.IsDir() {
				continue
			}

			// 生成ディレクトリ取得
			createDirPath := getCreateDir(targetDir, fileName, options)

			if !isExists(createDirPath) {
				printf(*options.IsVerbose, "ディレクトリを作成します: %s\n", createDirPath)
				createDir(createDirPath, *options.IsDryrun)
			}

			// 作成した(または、する予定だった)ディレクトリへファイルを移動
			moveTo := filepath.Join(createDirPath, fileName)
			printf(*options.IsVerbose, "ファイルを移動します: %s to %s\n", fileAbsPath, moveTo)
			moveFile(fileAbsPath, moveTo, *options.IsDryrun)
		}

	}

}

func parseArgs() Options {
	options := Options{
		OrganizeType: flag.String("type", "pattern", "整頓タイプ(letter_num or pattern)."),
		LetterNum:    flag.Int("letter_num", 8, "文字数。整頓タイプ「letter_num」で使用。ここで指定した長さの先頭文字列が整頓先ディレクトリになります"),
		Pattern:      flag.String("pattern", "^(.*?)_", "正規表現パターン。整頓タイプ「pattern」で使用。正規表現マッチの $1 にマッチする文字列が整頓先ディレクトリになります"),
		IsDryrun:     flag.Bool("dryrun", false, "dry run flag"),
		IsVerbose:    flag.Bool("verbose", false, "verbose mode flag"),
	}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options...] TARGET_DIRS...\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	return options
}

func isExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func getCreateDir(targetDir string, fileName string, options Options) string {

	printf(*options.IsVerbose, "対象ファイル: %s\n", fileName)

	createDirName := ""

	switch *options.OrganizeType {
	case "letter_num":
		// ファイル名の先頭 `letter_num` 文字を取得
		prefixLetters := string([]rune(fileName)[:*options.LetterNum])
		createDirName = prefixLetters

	case "pattern":
		// 正規表現マッチの $1 にマッチする文字列を取得s
		r := regexp.MustCompile(*options.Pattern + "(.*)")
		createDirName = r.ReplaceAllString(fileName, "$1")

	default:
		message := fmt.Sprintf("整頓タイプが不正です: %s,letter_num か pattern を指定してください。", *options.OrganizeType)
		panic(message)
	}

	// 整頓対象ディレクトリと join して、作成するディレクトリのパスを生成
	createDirPath := filepath.Join(targetDir, createDirName)
	printf(*options.IsVerbose, "整頓先ディレクトリ: %s\n", createDirPath)

	return createDirPath

}

func createDir(createDirPath string, isDryrun bool) {
	if !isDryrun {
		err := os.Mkdir(createDirPath, os.ModePerm)
		if err != nil {
			message := fmt.Sprintf("ディレクトリの作成に失敗しました: %s", err.Error())
			panic(message)
		}
	}
}

func moveFile(targetFile string, moveTo string, isDryrun bool) {
	if !isDryrun {
		err := os.Rename(targetFile, moveTo)
		if err != nil {
			message := fmt.Sprintf("ファイルの移動に失敗しました: %s", err.Error())
			panic(message)
		}
	}
}

func printf(isVerbose bool, format string, argv ...any) {
	if isVerbose {
		fmt.Printf(format, argv...)
	}
}
