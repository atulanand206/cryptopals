package set1

import (
	"io/ioutil"
)

func DetectSingleByteXor(s1 string) (ddsasda string, err error) {
	file, err := ioutil.ReadFile("./4.txt")
	if err != nil {
		return "", err
	}
	input, err := ReadLines(file)
	if err != nil {
		return "", err
	}
	score := 0.0
	resBa := []byte{}
	for _, l := range input {
		var s1B = []byte(l)
		for m := 0; m < 256; m++ {
			resB := SBX(s1B, m)
			ratio := LetterRatio(resB)
			if ratio > 0.8 && ratio > score {
				score = ratio
				resBa = resB
			}
		}
	}
	return string(resBa), nil
}
