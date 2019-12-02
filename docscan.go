package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"io/ioutil"
)

// The path to the text file to be scanned
const PATH string = "examples/paragraph.txt"

// Types of line
const BLANK string = "B"
const PARAGRAPH string = "P"
const HEADER string = "H"
const LIST_ITEM string = "L"

type Block struct {
	Type string
	Body string
}

// Calculates type of line
func lineType(line string) string {
	if len(line) > 0 {
		return PARAGRAPH
	}

	return BLANK
}

// Iterate through lines and merge paragraphs
func mergeParagraphs(paras []Block) []Block {
	blocks := make([]Block, 0)

	for i := 0; i < len(paras); i++ {
		if i > 0 && paras[i-1].Type == PARAGRAPH {
			lastIndex := len(blocks) - 1
			blocks[lastIndex].Body += " " + paras[i].Body
			continue
		}

		blocks = append(blocks, Block{ Type: lineType(paras[i].Type), Body: paras[i].Body })
	}

	return blocks
}

func main() {
	// Open the file
	file, err := os.Open(PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make([]Block, 0)

	// Loop through each line of the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		thisBlock := Block{Type: lineType(text),
			Body: text}

		lines = append(lines, thisBlock)
	}

	// Run processing functions
	paragraphs := mergeParagraphs(lines)

	// Output JSON file
	outFile, _ := json.MarshalIndent(paragraphs, "", " ")
	_ = ioutil.WriteFile("test.json", outFile, 0644)

}
