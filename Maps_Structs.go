package main

import (
	"fmt"
	"reflect"
)

// Doctor : We capitalize the name of the struct and also the fields in order to export them
// Lowercase field names would not allow anyone to access the fields of the struct
type Doctor struct {
	Number     int `required max:"100"` // This is called a tag. Extra annotation that can be used by some library for any purpose like validation
	ActorName  string
	Companions []string
}

// Surgeon has an embedded struct. It simply inherits the fields of the struct inside itself
type Surgeon struct {
	Doctor
	Speciality string
}

func main() {

	// The value datatype can be anything
	// The key datatype has to be something that can be tested for equality.
	// We can't use slices or maps as key types. We can use bool, numeric, float, string, arrays
	var medianAge map[string]int = map[string]int{
		"baroda":    25,
		"mumbai":    22,
		"bangalore": 26,
	}

	fmt.Println(medianAge)

	m := make(map[string]int)
	m["nisarg"] = 5
	m["hello"] = 1
	delete(m, "hello")
	val, ok := m["hello"]
	fmt.Println(val, ok) // ok to check if the element is present
	val, ok = m["nisarg"]
	fmt.Println(val, ok)
	fmt.Println(len(m))
	n := m
	n["nisarg"] = 7 // Reference(pointer) based like slice
	fmt.Println(n)

	// Structs
	// Structs are not references(pointers) like slices or maps
	// Structs copy the entire struct not pass a reference
	var aDoctor Doctor = Doctor{
		Number:    5,
		ActorName: "Nisarg",
		Companions: []string{
			"Hello",
			"World",
			"Welcome",
		},
	}

	// Anonymous struct
	bDoctor := struct {
		name string
		age  int
	}{
		name: "Nisarg",
		age:  21,
	}

	fmt.Println(aDoctor)
	fmt.Println(bDoctor)

	sDoc := Surgeon{}
	sDoc.Number = 23
	sDoc.Speciality = "Heart"

	t := reflect.TypeOf(Doctor{})
	field, _ := t.FieldByName("Number")
	fmt.Println(t)
	fmt.Println(field.Tag)
	fmt.Println(sDoc)
}
