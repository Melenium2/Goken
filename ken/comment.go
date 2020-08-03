package ken

import "fmt"

type Comment struct {
	comment string
}

func NewComment(comment string) *Comment {
	return &Comment{
		comment: comment,
	}
}

func NewCommentf(comment string, args ...interface{}) *Comment {
	return &Comment{
		comment: fmt.Sprintf(comment, args),
	}
}

func (c *Comment) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)

	return fmt.Sprintf("%s// %s\n", indentLvl, c.comment), nil
}
