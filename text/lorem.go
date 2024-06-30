package text

import (
	"fmt"

	"github.com/google/uuid"

	lorem "github.com/derektata/lorem/ipsum"
)

func PrintUuids() {
	for i := 0; i < 5; i++ {
		fmt.Println(uuid.New().String())
	}
}

func LoremGenerator() {
	// Create a new generator
	g := lorem.NewGenerator()
	g.WordsPerSentence = 10     // Customize how many words per sentence
	g.SentencesPerParagraph = 5 // Customize how many sentences per paragraph
	g.CommaAddChance = 3        // Customize the chance of a comma being added to a sentence

	// Generate words
	words := g.Generate(100)
	fmt.Println(words)

	// Generate paragraphs
	paragraphs := g.GenerateParagraphs(3)
	fmt.Println(paragraphs)
}
