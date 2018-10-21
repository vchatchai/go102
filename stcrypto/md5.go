package stcrypto

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

var content = "This is content to check"

func MD5Test() {

	checkSum := MD5(content)
	checkSum2 := MD5File("data.dat")

	fmt.Printf("Checksum 1: %s\n", checkSum)
	fmt.Printf("Checksum 2: %s\n", checkSum2)
	if checkSum == checkSum2 {
		fmt.Println("Content matches!!!")
	}

}

func MD5(content string) string {
	h := md5.Sum([]byte(content))
	return fmt.Sprintf("%x", h)
}

func MD5File(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	h := md5.New()

	_, err = io.Copy(h, file)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))

}
