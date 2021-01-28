package takeaction

import (
	"fmt"
	"io"
)

type Helloer interface {
	Hello(string)
}

type Greeting string

func (g Greeting) Hello(name string) {
	fmt.Printf("%s %s\n", g, name)
}

type Invitation struct {
	event string
}
func (inv *Invitation) Hello(name string) {
	fmt.Printf("Welcome to my %s, %s! Please on come in\n", inv.event, name)
}

func InterfaceDemo() {
	var h Helloer
	//h = Greeting("G'day")
	h = &Invitation { event: "hackathon" }
	h.Hello("Benjamin")
}



// TerminalWriter writes to a terminal, breaking the lines
// at the given width.
type TerminalWriter struct {
	width int
}

// Write writes slice p to stdout, in chunks of tw.width bytes,
// separated by newline.
// It returns the number of successfully written bytes, and
// any error that occurred.
// If the complete slice is written, Write returns error io.EOF
//
//It must meet the following requirements:
//
//It writes to os.Stdout. (You can use the standard fmt.Print/f/ln family here, or os.Stdout.Write().)
//After tw.width bytes, start a new line.
//When the function encounters any error during writing, return an appropriate error, and the number of bytes successfully written.

func (tw *TerminalWriter) Write(p []byte) (n int, err error) {

	for _, e := range p {
		n++
		if n % tw.width == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("%s", string(e))
		}
	}

	return n, io.EOF
}

// NewTerminalWriter creates a new TerminalWriter. width is
// the terminal's width.
func NewTerminalWriter(width int) *TerminalWriter {
	return &TerminalWriter{width: width}
}

func RunTerminalWriter() {
	s := []byte("This is a long string converted into a byte slice for testing the TerminalWriter.")
	tw := NewTerminalWriter(20)
	n, err := tw.Write(s)
	fmt.Println(n, "bytes written. Error:", err)

}