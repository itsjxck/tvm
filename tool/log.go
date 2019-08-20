package tool

import (
	"fmt"
	"os"
)

// Log =>
func (t *Tool) Log(message ...interface{}) {
	m := []string{fmt.Sprintf("%s", t.Name)}
	var i []interface{}
	i = append(i, m)
	message = append(i, message...)
	fmt.Println(message...)
}

// Quit =>
func (t *Tool) Quit(message ...interface{}) {
	t.Log(message...)
	os.Exit(1)
}
