# golist

Golist is my package, which includes a List type that has the same methods as JavaScript arrays, and a few additional ones.

List[T any] []T

Methods:
+Map
+ForEach
+Filter
+Push
+Concat
+Pop
+Splice

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
  	people := golist.List[Person]{{1.85, 78}, {1.7, 56}, {1.84, 76}, {1.82, 72}}
  	fmt.Println(pessoas.Splice(1, 1)) // [{1.85, 78}]
  	fmt.Println(pessoas) // [{1.7, 56}, {1.84, 76}, {1.82, 72}]
  
   pessoas.ForEach(func(v Person, index int) {
  		fmt.Println(v.BMI()) // shows each person's BMI
  	})
  }
```
