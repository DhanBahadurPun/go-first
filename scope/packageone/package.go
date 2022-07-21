package packageone

var privateVar = "I am Private"

var PublicVar = "I am public (or exported)"

func notExported() {
	//  available in this packageone
}

func Exported() {
	// available in all other package
}
