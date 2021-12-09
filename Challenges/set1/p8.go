package set1

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func DetectAESECB() (decoded int, err error) {
	file, err := ioutil.ReadFile("./8.txt")
	if err != nil {
		return -1, err
	}
	input, err := ReadLines(file)
	if err != nil {
		return -1, err
	}
	res := Detect(input)
	fmt.Println(res)
	return res, nil
}

func ReadLines(f []byte) ([][]byte, error) {
	var input [][]byte
	scanner := bufio.NewScanner(bytes.NewReader(f))
	for scanner.Scan() {
		hdl := make([]byte, hex.DecodedLen(len(scanner.Bytes())))
		_, err := hex.Decode(hdl, scanner.Bytes())
		if err != nil {
			return nil, err
		}
		input = append(input, hdl)
	}
	return input, nil
}

func Detect(input [][]byte) int {
	for i, ln := range input {
		chunks := make([][]byte, 0)
		for j := 0; j < len(ln); j += 16 {
			batch := ln[j:min(j+15, len(ln))]
			for _, c := range chunks {
				if bytes.Equal(c, batch) {
					return i + 1
				}
			}
			chunks = append(chunks, batch)
		}
	}
	return 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
