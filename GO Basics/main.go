package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strings"
)

func main() {
	// Hello World Example
	fmt.Println("It's yo boy Crimson!!!")

	// // Variable example
	// // Here you can print VARIABLES exactly same as PYTHON
	// // Declare variables exactly like JavaScript

	// var name = "Kamal"
	// const surname = "Maharana"
	// fmt.Println("Name:", name, surname)

	// // Just like 'C' language, format specifier %d, %s; %v is variable format specifier
	// fmt.Printf("Name: %v %v\n", name, surname)

	// // Data Types
	// // If we declare a variable and don't initialize it with a value,
	// // then we need to declare it's DATA TYPE
	// // When you initialize, GO smartly assigns the type based on value assigned.

	// var middlename string
	// var age int
	// var male bool
	// var hobbies [3] string

	// middlename = "Crimson"
	// age = 21
	// male = true

	// hobbies[0] = "Cricket"
	// hobbies[1] = "Programming"
	// hobbies[2] = "Listening Music"

	// // String Concatentation same as PYTHON
	// name += " " + middlename
	// fullname := name + " " + surname
	// fmt.Println("Name: ", fullname)
	// fmt.Println("Age:", age)
	// if male {
	// 	fmt.Println("Gender: Male")
	// } else {
	// 	fmt.Println("Gender: Female")
	// }
	// for index, value := range hobbies {
	// 	fmt.Printf("Hobby[%d]: %s\n", index, value)
	// }

	// // strings.Join(slice, seperator_string)
	// fmt.Println(strings.Join(hobbies[0:3], " "))

	// Go also support SLICING same as PYTHON
	// Here SLICE is actual data type unlike PYTHON where SLICE of LIST gives LIST
	// SLICE are RESIZABLE arrays

	// var skills []string = make([]string, 5, 10)
	// skills[0] = "Python"
	// skills[1] = "GO"
	// additional_skills := []string{"Woosh"}
	// skills = append(skills, additional_skills...)
	// for index, value := range skills {
	// 	fmt.Println(index, value)
	// }
	// Length of array or slice = len(arr)

	// Taking user input
	// For this we require pointers
	// We need to declare variable WITH TYPE and then take input on it's pointer
	// var name string
	// var skills []string
	// var skill string
	// fmt.Print("Enter you name: ")
	// fmt.Scan(&name) // & sign represent pointer, & attached to any variable will give it's address
	// fmt.Print("Your name: ", name)
	// println()
	// for i := 0; i < 5; i++ {
	// 	fmt.Print("Enter your skill: ")
	// 	fmt.Scanln(&skill)
	// 	skills = append(skills, skill)
	// }
	// // fmt.Println()
	// println(skills)

	// Different ways to take input other than fmt.Scan()
	// var skills []string
	// 1. Scanner (like JAVA)
	// scanner := bufio.NewScanner(os.Stdin)
	// Infinite FOR LOOP
	// for {
	// 	fmt.Println("Enter your skill:")
	// 	scanner.Scan()
	// 	text := scanner.Text()
	// 	// var text string
	// 	// fmt.Scan(&text)
	// 	if len(text) != 0 {
	// 		skills = append(skills, text)
	// 	} else {
	// 		break
	// 	}
	// }
	// fmt.Println("My Skills are:", skills, len(skills))

	// name := "Kamal Maharana"
	// // string.Fields(s) will return a SLICE of strings space seperated
	// strings.Fields(name)
	server()
}
