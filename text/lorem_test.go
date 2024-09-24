package text

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
	assertions "github.com/stretchr/testify/assert"
)

func TestGenerateUuid(t *testing.T) {
	assert := assertions.New(t)
	uuidString := GenerateUuid()
	assert.NotNil(uuidString, "GenerateUuid must always return a value!")
}

func TestGenerateFakeStrings(t *testing.T) {
	assert := assertions.New(t)

	const numOfElements int = 200
	data := GenerateFakeStrings(numOfElements)

	collection := lo.Map(data, func(item string, idx int) string {
		return fmt.Sprintf("%d. %s", idx, item)
	})

	assert.NotNil(collection, "Collection cannot be nil!")

	for _, val := range collection {
		fmt.Println(val)
	}
}
