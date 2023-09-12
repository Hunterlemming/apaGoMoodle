package main

// TODO: Replace this with terminal input
const inputName = "./files/input/association-test.xml"

// TODO: Replace this with terminal input
const outputName = "./files/output/out.xml"

func main() {
	// f, err := os.Open(inputName)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()

	// decoder := xml.NewDecoder(f)
	// questions, err := ParseQuestions(decoder)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(questions)

	writeFile(outputName)
}
