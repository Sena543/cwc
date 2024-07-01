package wordcount

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
)

func ReadFile(filePath string) []byte {
	fileData, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("error reading file: ", err)
	}
	return fileData
}

func GetFileBytes(fileData []byte) int64 {

	return int64(len(fileData))
}

func GetLineCount(fileData []byte) int64 {
	return int64(len(bytes.Split(fileData, []byte("\n"))))
}

func GetCharacterCount(fileData []byte) int64 {
	return int64(len(bytes.Runes(fileData)))
}

func GetWordCount(fileData []byte) int64 {
	return int64(len(bytes.Fields(fileData)))
}

func CountEntry() {
	b, w, l, c := getCmdArgs()
	//check if input is comming from Stdin
	stdInCheck, err := os.Stdin.Stat()

	if err != nil {
		fmt.Println("Std error: ", err)
	}

	var fileDataPointer []byte
	var fileName string = ""
	if (stdInCheck.Mode() & os.ModeCharDevice) == 0 { // data being piped to standard input
		fileDataPointer, _ = io.ReadAll(os.Stdin)
	} else {
		currWorkDir, _ := os.Getwd()
		filePath := path.Join(currWorkDir, os.Args[len(os.Args)-1])
		fileDataPointer = ReadFile(filePath)
		fileName = os.Args[len(os.Args)-1]
	}

	if b {
		fmt.Println(GetFileBytes(fileDataPointer))
	}

	if w {
		fmt.Println(GetWordCount(fileDataPointer))
	}

	if l {
		fmt.Println(GetLineCount(fileDataPointer))
	}

	if c {
		fmt.Println(GetCharacterCount(fileDataPointer))
	}

	if !b && !w && !l && !c {
		fmt.Println(GetLineCount(fileDataPointer), GetWordCount(fileDataPointer),
			GetCharacterCount(fileDataPointer), fileName)
	}

}

func getCmdArgs() (bool, bool, bool, bool) {
	var byteCount, wordCount, lineCount, characterCount bool

	flag.BoolVar(&byteCount, "c", false, "Returns the number of bytes in the file passed")
	flag.BoolVar(&wordCount, "w", false, "Returns the number of words in the file passed")
	flag.BoolVar(&lineCount, "l", false, "Returns the number of lines in the file passed")
	flag.BoolVar(&characterCount, "m", false, "Returns the number of characters in the file passed")

	flag.Parse()

	return byteCount, wordCount, lineCount, characterCount
}
