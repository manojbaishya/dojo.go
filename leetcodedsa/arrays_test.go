package leetcodedsa

import (
	"fmt"
	assertions "github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKeyToValuesStore(t *testing.T) {
	assert := assertions.New(t)

	category := "Plants"
	species := []string{"Hibiscus", "Palm", "Rose", "Gerbera"}

	expected := &KeyToValuesStore{key: category, values: species}
	actual := NewKeyToValuesStore(category, species...)

	assert.Equal(expected, actual)
}

func TestKey(t *testing.T) {
	assert := assertions.New(t)

	category := "Plants"
	kv := NewKeyToValuesStore(category, "Hibiscus", "Palm", "Rose", "Gerbera")

	expected := category
	actual := kv.Key()

	assert.Equal(expected, actual)
}

func TestValues(t *testing.T) {
	assert := assertions.New(t)

	category := "Plants"
	species := []string{"Hibiscus", "Palm", "Rose", "Gerbera"}

	kv := &KeyToValuesStore{key: category, values: species}
	expected := species
	actual := kv.Values()

	fmt.Println(actual)

	assert.Equal(expected, actual)
}
