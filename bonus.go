package greetings
import "fmt"
var message string
var n [100] int
type Animals struct{
    name string
    age int
    color string

}
type Persons struct{
     name string
     Lastname string
     id int
     age int
     language string
}
func Array() int{
    for i =0; i<100; i++{
    n[i] = i
   }

   for j= 0; j<100; j++{
        return n[j]
   }
}
func Message(name string) string{
    message := fmt.Sprintf("hi , %v Welcome!", name)
    return message
}
func main(){
    var animal Animals
    var person Persons

    animal.name = "monkey"
    animal.age = 1
    animal.color = "black"

    person.name = "Zangar"
    person.Lastname = "Tasbolat"
    person.id ="3424234324"
    person.age ="24"
    person.language = "kazakh"
}
func (a *Animals) change(){
    a.name = "newName"
}