package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type JapaneseCharacter struct {
	Id        int
	Character string
}

func main() {
	fmt.Println("starting")

	file, err := os.Open("./hiriganaList.txt")

	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	id := 1

	characterSlice := make([]JapaneseCharacter, 0)

	for scanner.Scan() {
		c := JapaneseCharacter{Id: id, Character: strings.TrimSpace(scanner.Text())}
		characterSlice = append(characterSlice, c)
		id++
	}

	// fmt.Println(characterSlice)
	jsonfile, _ := json.MarshalIndent(characterSlice, "", " ")
	_ = ioutil.WriteFile("hirigana.json", jsonfile, 0644)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
