package ken

type NewLine struct {

}

func NewNewLine() *NewLine {
	return &NewLine{}
}

func (l *NewLine) Generate(indent int) (string, error) {
	return "\n", nil
}
