package dot_go

import "fmt"

func SprintPtr[T any](input *T) string {
	if input == nil {
		return ""
	}
	return fmt.Sprintf("%v", *input)
}
