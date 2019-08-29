package terraform

import (
	"fmt"
	"os"
)

// Log =>
func (t *Terraform) Log(message ...interface{}) {
	fmt.Println(message...)
}

// Quit =>
func (t *Terraform) Quit(message ...interface{}) {
	t.Log(message...)
	os.Exit(1)
}
