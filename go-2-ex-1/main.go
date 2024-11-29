 package main

import "fmt"

type FullName struct {
	Firstname string
	Lastname string
}


type Birthdate struct {
	Day byte
	Month byte
	Year int16
}

type Profile struct {
	name FullName
	dayborn Birthdate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		
		name: FullName{Firstname:"Mika",Lastname:"Blättler"},
		dayborn: Birthdate{Day:3, Month:7, Year:2008},
		NumberOfSiblings: 1,   
		ZodiacSign:'\u264b', 
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
