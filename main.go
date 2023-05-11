package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var (
	revert = flag.Bool("r", false, "Revert conversion")
	binary = flag.Bool("b", false, "Convert to binary")
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func convertFromBin(src string) string {
	var out []byte
	for start := 0; start < len(src); start += 8 {
		b, err := strconv.ParseInt(src[start:start+8], 2, 64)
		CheckErr(err)
		out = append(out, byte(b))
	}
	return string(out)
}

func convertToBin(src string) string {
	var dst string
	for _, c := range src {
		dst = fmt.Sprintf("%s%.8b", dst, c)
	}
	return dst
}

func main() {
	flag.Parse()

	args := flag.Args()
	input := strings.Join(args, " ")

	if *revert {
		if *binary {

			fmt.Println(convertFromBin(input))

		} else {

			out, err := hex.DecodeString(input)
			CheckErr(err)
			fmt.Println(string(out))

		}
	} else {
		if *binary {

			fmt.Println(convertToBin(input))

		} else {

			out := hex.EncodeToString([]byte(input))
			fmt.Println(strings.ToUpper(out))

		}
	}
}
