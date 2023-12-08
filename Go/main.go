/**
 * package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func writeToFile(filePath, content string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File written successfully.")
}

func main() {
	filePath := "checkPrime.go"

	content := "package main\n\n"
	content += "import (\n"
	content += "\t\"fmt\"\n"
	content += "\t\"math\"\n"
	content += "\t\"os\"\n"
	content += "\t\"strconv\"\n"
	content += ")\n\n"
	content += "func isPrime(num int) bool {\n"
	content += "\tif num < 2 {\n"
	content += "\t\treturn false\n"
	content += "\t}\n\n"
	content += "\treturn all(num %% (2:math.Sqrt(float64(num))) != 0)\n"
	content += "}\n\n"
	content += "func all(vals []bool) bool {\n"
	content += "\tfor _, v := range vals {\n"
	content += "\t\tif !v {\n"
	content += "\t\t\treturn false\n"
	content += "\t\t}\n"
	content += "\t}\n"
	content += "\treturn true\n"
	content += "}\n\n"
	content += "func checkPrime(num int) bool {\n"
	content += "\tfor i := 1; i <= 1000; i++ {\n"
	content += "\t\tif isPrime(i) && num == i {\n"
	content += "\t\t\treturn true\n"
	content += "\t\t} else if !isPrime(i) && num == i {\n"
	content += "\t\t\treturn false\n"
	content += "\t\t}\n"
	content += "\t}\n"
	content += "\treturn false\n"
	content += "}\n\n"
	content += "func main() {\n"
	content += "\tnumStr := \"\"\n"
	content += "\tfmt.Print(\"Enter a number: \")\n"
	content += "\tfmt.Scanln(&numStr)\n\n"
	content += "\tnum, err := strconv.Atoi(numStr)\n"
	content += "\tif err != nil {\n"
	content += "\t\tfmt.Println(\"Invalid input. Please enter a valid integer.\")\n"
	content += "\t\treturn\n"
	content += "\t}\n\n"
	content += "\tif checkPrime(num) {\n"
	content += "\t\tfmt.Printf(\"%d is a prime number.\\n\", num)\n"
	content += "\t} else {\n"
	content += "\t\tfmt.Printf(\"%d is not a prime number.\\n\", num)\n"
	content += "\t}\n"
	content += "}\n"

	writeToFile(filePath, content)
}

DO THIS IN GO:
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func writeToFile(filePath, content string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File written successfully.")
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	filePath := "checkPrime.go"

	content := "package main\n\n"
	content += "import (\n"
	content += "\t\"fmt\"\n"
	content += "\t\"strconv\"\n"
	content += ")\n\n"
	content += "func isPrime(num int) bool {\n"
	content += "\tif num < 2 {\n"
	content += "\t\treturn false\n"
	content += "\t}\n\n"
	content += "\tfor i := 2; i <= num/2; i++ {\n"
	content += "\t\tif num%i == 0 {\n"
	content += "\t\t\treturn false\n"
	content += "\t\t}\n"
	content += "\t}\n\n"
	content += "\treturn true\n"
	content += "}\n\n"
	content += "func all(vals []bool) bool {\n"
	content += "\tfor _, v := range vals {\n"
	content += "\t\tif !v {\n"
	content += "\t\t\treturn false\n"
	content += "\t\t}\n"
	content += "\t}\n"
	content += "\treturn true\n"
	content += "}\n\n"
	content += "func main() {\n"
	content += "\tnumStr := \"\"\n"
	content += "\tfmt.Print(\"Enter a number: \")\n"
	content += "\tfmt.Scanln(&numStr)\n\n"
	content += "\tnum, err := strconv.Atoi(numStr)\n"
	content += "\tif err != nil {\n"
	content += "\t\tfmt.Println(\"Invalid input. Please enter a valid integer.\")\n"
	content += "\t\treturn\n"
	content += "\t}\n\n"
	for i := 1; i <= 1000; i++ {
		if isPrime(i) {
			content += "\tif num == " + strconv.Itoa(i) + " {\n"
			content += "\t\tfmt.Printf(\"%d is a prime number.\\n\", num)\n"
			content += "\t}\n"
		} else {
			content += "\tif num == " + strconv.Itoa(i) + " {\n"
			content += "\t\tfmt.Printf(\"%d is not a prime number.\\n\", num)\n"
			content += "\t}\n"
		}
	}
	content += "}\n"

	writeToFile(filePath, content)
}
