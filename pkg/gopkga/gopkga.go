package gopkga

import "fmt"

// Hello returns a greeting for the named person.
func Hello1(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!... this is from gopkga", name)
    return message
}
