package main

import "os"
import "fmt"

func main() {
	app, err := NewMainLoop();
	if err != nil {
		fmt.Println(err);
		os.Exit(-1);
	}

	app.run();
	app.exit();
}