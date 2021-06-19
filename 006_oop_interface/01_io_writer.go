package main

import (
	"fmt"
	"os"
)

func main() {
	// เราใช้ fmt.Println() ประจำเลย
	// ถ้าไปดูไส้ในมันคือ fmt.Fprintln(os.Stdout, a...)
	// os.Stdout implement io.Writer
	_, _ = fmt.Fprintln(os.Stdout, "Hello World")
}
