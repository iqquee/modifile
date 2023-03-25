package modifile

import (
	"bufio"
	"os"
)

func Compare(file1, file2 string) error {
	f1, err := os.Open(file1)
	if err != nil {
		return err
	}

	f2, err := os.Open(file2)
	if err != nil {
		return err
	}

	scanner1 := bufio.NewScanner(f1)
	scanner2 := bufio.NewScanner(f2)

	var fileContent1 []string
	var fileContent2 []string

	for scanner1.Scan() {
		currentLineText := scanner1.Text()
		fileContent1 = append(fileContent1, currentLineText)
	}

	for scanner2.Scan() {
		currentLineText := scanner2.Text()
		fileContent2 = append(fileContent2, currentLineText)
	}

	// for i, v := range file1Lines1 {
	// 	fmt.Printf("This is the index: %d and value: %s\n", i, v)
	// }

	return nil

}
