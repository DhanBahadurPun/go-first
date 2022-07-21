package packagechallenge

import "fmt"

var PackageVar = "Hi, I am from the package challenge!"

func PrintMe(s1 string, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}
