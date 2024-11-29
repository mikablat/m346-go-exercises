package main

import "fmt"

func main() {
	modules:= map[int]string{
		104: "Modul 104 Text Beispiel",
		117: "Modul 117 Text Beispiel",
		346: "Modul 346 Text Beispiel",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 117)
	// TODO: add one
	modules[100] = "Neues Modul 100"
	// TODO: replace one
	modules[104] = "Modul zum ersetzen von 104"
	fmt.Println(modules)
}
