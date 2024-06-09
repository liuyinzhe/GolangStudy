package main

import (
	"a1/lib"
	"a1/lib/lib2" // from a1.lib import lib2
	"fmt"
)

func main() {
	//lib.lib1Test()
	//lib2.lib2Test()
	fmt.Println("a1", "import a1.lib")
	lib.Lib1Test()
	fmt.Println("a1", "import a1.lib.lib2")
	lib2.Lib2Test()
}
