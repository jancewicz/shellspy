package main

import "log"

var version = "0.0.1"

// Your CLI goes here!
func main() {
	// file, err := os.Create("shellspy.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// fmt.Println("Recording session to 'shellspy.txt'")
	// readInput := shellspy.ReadUserInput

	// if err := shellspy.RunShell(readInput, file); err != nil {
	// 	log.Fatal(err)
	// }

	cfg := config{
		addr: ":8080",
		env:  "DEV",
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	err := app.run(mux)
	if err != nil {
		log.Fatal(err)
	}
}
