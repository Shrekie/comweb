package menu

import (
	"bufio"
	"fmt"
	"github.com/shrekie/comweb/commands"
	"os"
)

// InputReader implements a text display and line reader
type InputReader interface {
	Prompt(t string) (string, error)
}

var bufferReader = bufio.NewReader(os.Stdin)

// BufferInput is a buffer stream line reader
type BufferInput struct{}

// Prompt for BufferInput
func (ir *BufferInput) Prompt(t string) (string, error) {
	fmt.Print(t)
	text, err := bufferReader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return text[:len(text)-1], nil
}

// @todo: send name, arg, site all in one go and
// parse the segmented strings
var menuStages = [3]string{"Name: ", "Argument: ", "Site: "}

// Select stages the three inputs for a command
func Select(ir InputReader) (string, error) {
	for index, element := range menuStages {
		menuStages[index], _ = ir.Prompt(element)
	}
	query := &commands.Query{Name: menuStages[0],
		Arg: menuStages[1], Site: menuStages[2]}

	result, _ := query.Execute()
	fmt.Printf("%s \n", result)
	return result, nil
}
