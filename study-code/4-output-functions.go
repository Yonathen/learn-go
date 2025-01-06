package main

import "fmt"

func main() {
	var str1, str2 = "Print : Prints", "On Same lines"

	fmt.Print(str1)
	fmt.Print(str2)
	fmt.Print("\n", str1, "\n", str2, "\n")

	var str3, str4 = "Println : Prints", "Always on the new line"
	fmt.Println(str3)
	fmt.Println(str4)

	var (
		fmt1        = "i has value = %v and type %T\n"
		fmt2        = "j has value = %v and type %T\n"
		fmt3        = "hex dump of i = %x\n"
		str  string = "Hello Again"
		j    int    = 20
	)
	fmt.Printf(fmt1, str, str)
	fmt.Printf(fmt2, j, j)
	fmt.Printf(fmt3, str)

	fmt.Println("\nFloat Formatting Verbs")
	var txt = "AAAA"
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	fmt.Println("\nInteger Formatting Verbs")
	var i = 15

	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)

	fmt.Println("\nFloat Formatting Verbs")
	var f = 3.141

	fmt.Printf("%e\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%6.2f\n", f)
	fmt.Printf("%g\n", f)
}
