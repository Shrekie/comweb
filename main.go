package main

import "github.com/shrekie/comweb/menu"

func main() {
	bufferInput := new(menu.BufferInput)
	menu.Select(bufferInput)
}
