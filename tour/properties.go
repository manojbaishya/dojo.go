package tour

import (
	"fmt"
	"hash/fnv"
)

type GetterDemo struct {
	Name        string
	Description string
}

func (getterDemo *GetterDemo) Hash() uint64 {
	return computeHash(getterDemo.Name, getterDemo.Description)
}

func computeHash(name string, desc string) uint64 {
	hsh := fnv.New64a()
	hsh.Write([]byte(fmt.Sprintf("%s-%s", name, desc)))

	hash := hsh.Sum64()
	return hash
}

func NewGetterDemo(name string, desc string) GetterDemo {
	return GetterDemo{
		Name:        name,
		Description: desc,
	}
}
