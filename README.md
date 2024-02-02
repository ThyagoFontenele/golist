# golist

Golist is my package, which includes a List type that has some good methods for arrays

> List[T any] []T

** Methods:
- Where(fn func(v T, index int) bool) []T
- ForEach(fn func(v T, index int))
- Filter(fn func(v T, index int) bool) []T
- Push(value T)
- Concat(value []T)
- Pop()
- Splice(elementIndex int, n int)

#Example
```
  import (
  	"fmt"
   	"github.com/ThyagoFontenele/golist"
  )
  
  type Person struct {
  	height, weight float32
  }
  
  func (p Person) BMI() float32 {
  	return p.weight / (p.height * p.height)
  }
  
  func main() {
    people := := golist.List[Person] {
  		{"Thyago", 1.85, 78},
  		{"Julia", 1.7, 56},
  		{"Fernando", 1.84, 76},
  		{"Maria", 1.82, 72},
	  }           

    people.Splice(0, 1) // Removing the firts person
    people.Pop()        // Removing the last person
    fmt.Println(people) // {{"Julia", 1.7, 56}, {"Fernando", 1.84, 76}}
    
    pessoas.ForEach(func(v Person, index int) {
      fmt.Println(v.BMI())  // shows each person's BMI
    })
  }
```
