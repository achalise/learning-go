package valid_parenthesis

import (
	"testing"
	"fmt"
)

func TestIsValid(t *testing.T)  {
	input := "({})"
	v :=  isValid(input);
	fmt.Printf("%s is valid: %c", input, v)
}
