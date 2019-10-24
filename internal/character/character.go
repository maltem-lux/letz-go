package character

import (
	"fmt"
)

type Character struct {
	Id int32
	Age int32
	Name string
	Gender string
}

func (c *Character) ToString() string {
	return fmt.Sprintf("Id : %v, Age : %v, Name : %v, Gender : %v.",
		c.Id, c.Age, c.Name, c.Gender)
}
