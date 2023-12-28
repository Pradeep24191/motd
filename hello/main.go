package main

import (
	"fmt"
	"src/motd/message"
	"os"
	"bufio"
	"strings"
	"flag"
)

/*func main() {

	file, err := os.OpenFile("notes.txt", os.O_WRONLY|os.O_CREATE, 0644)

	defer file.Close()

	if err != nil {
		fmt.Println("Unable to open file : ", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Your Greeting : ")
	pharse, _  := reader.ReadString('\n')
	pharse = strings.TrimSpace(pharse)

	fmt.Println("Your Name : ")
	name, _  := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	message := message.Greeting(name, pharse)
	// fmt.Println(message)
	// err := ioutil.WriteFile("src/motd", []byte(message), 0644)

	_, err = file.Write([]byte(message))

	if err != nil {
		fmt.Println("Unable to write file : ", err)
		os.Exit(1)
	}
}
*/

// basic func method
// func greeting(name string, message string) string {
// 	return fmt.Sprintf("%s, %s", message, name)
// }


func main () {
	// Define flags
	var name string
	var greeting string
	var prompt bool
	var preview bool

	// Parse flags
	flag.StringVar(&name, "name", "", "name to use in greeting message")
	flag.StringVar(&greeting, "greeting", "", "Greeting message")
	flag.BoolVar(&preview, "preview", false, "use to preview input message")
	flag.BoolVar(&prompt, "prompt", false, "use prompt to input name and message")

	flag.Parse()

	// Show usage if flags are invalid
	if prompt == false && (name == "" || greeting == "") {
		flag.Usage()
		os.Exit(1)
	}

	// Optionally print flags and exist based on DEBUG env variables
	if os.Getenv("DEBUG") != "" {
		fmt.Println("Name : ", name)
		fmt.Println("Greeting : ", greeting)
		fmt.Println("Prompt : ", preview)
		fmt.Println("Preview : ", prompt)
		os.Exit(0)
	}

	// Conditionally read from Stdin
	if prompt {
		name, greeting = renderPrompt()
	}

	// Generate message
	message := message.Greeting(name, greeting)

	// Either preview the message or write to the file	
	if preview {
		fmt.Println(message)
	} else {
		//write content
		file, err := os.OpenFile("welcome.txt", os.O_WRONLY|os.O_CREATE, 0644)
		defer file.Close()

		handleError(err)

		_, err = file.Write([]byte(message))

		handleError(err)
	}
}


func renderPrompt() (name, greeting string) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Your Greeting : ")
	greeting, _  = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)

	fmt.Println("Your Name : ")
	name, _  = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
}
