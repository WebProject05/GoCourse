package main

// XML - Extensible Markup Language

import (
	"encoding/xml"
	"fmt"
	"log"
)

// XML is a way for a document to be read easily by machine and human readble

// Similar to json tags we can also use omitempty and "-"
type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	City    string   `xml:"city"`
	Email   string   `xml:"email"`
	Address Address  `xml:"address"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {
	person := Person{
		Name:  "Santosh",
		Age:   20,
		City:  "Hyderabad",
		Email: "San@gmail.com",
		Address: Address{
			City:  "Hyderabad",
			State: "Telangana",
		},
	}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println("The xml data is:", string(xmlData)) // Byte slice to string

	xmlData1, err1 := xml.MarshalIndent(person, "", " ") // This will make sure that the data is indentent
	if err1 != nil {
		log.Fatalln("Error:", err1)
	}
	fmt.Println("The indentent data is: \n", string(xmlData1))

	xmlRaw := `<person><name>San</name><age>20</age></person>`

	var xmlPerson Person
	err3 := xml.Unmarshal([]byte(xmlRaw), &xmlPerson)
	if err3 != nil {
		log.Fatalln("Error unmarshalling xml:", err3)
	}
	fmt.Println(xmlPerson)

	// if we want the tag to look like <person name="santosh"> then we have to use attr in the xml tagging
	
}
