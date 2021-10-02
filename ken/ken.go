package ken

import (
	"bytes"
	"io"
	"os/exec"
	"strings"

	"github.com/Melenium2/goken/kenerrs"
)

type State interface {
	Generate(indentlvl int) (string, error)
}

func Indent(level int) string {
	var indent string
	{
		for i := 0; i < level; i++ {
			indent += "\t"
		}
	}

	return indent
}

type Root struct {
	states         []State
	gofmt          bool
	gofmtOptions   []string
	goimports      bool
	syntaxChecking bool
}

func NewRoot(states ...State) *Root {
	return &Root{
		states: states,
	}
}

func (r *Root) AddState(state State) *Root {
	return &Root{
		states:         append(r.states, state),
		gofmt:          r.gofmt,
		gofmtOptions:   r.gofmtOptions,
		goimports:      r.goimports,
		syntaxChecking: r.syntaxChecking,
	}
}

func (r *Root) States(states ...State) *Root {
	return &Root{
		states:         states,
		gofmt:          r.gofmt,
		gofmtOptions:   r.gofmtOptions,
		goimports:      r.goimports,
		syntaxChecking: r.syntaxChecking,
	}
}

func (r *Root) Gofmt(gofmtOptions ...string) *Root {
	return &Root{
		states:         r.states,
		gofmt:          true,
		gofmtOptions:   gofmtOptions,
		goimports:      r.goimports,
		syntaxChecking: r.syntaxChecking,
	}
}

func (r *Root) Goimports() *Root {
	return &Root{
		states:         r.states,
		gofmt:          r.gofmt,
		gofmtOptions:   r.gofmtOptions,
		goimports:      true,
		syntaxChecking: r.syntaxChecking,
	}
}

func (r *Root) EnableSyntaxChecking() *Root {
	return &Root{
		states:         r.states,
		gofmt:          r.gofmt,
		gofmtOptions:   r.gofmtOptions,
		goimports:      r.goimports,
		syntaxChecking: true,
	}
}

func (r *Root) Generate(indentlvl int) (string, error) {
	var code string
	{
		for _, state := range r.states {
			gen, err := state.Generate(indentlvl)
			if err != nil {
				return "", err
			}
			code += gen
		}

		if r.syntaxChecking {
			_, err := r.applyGofmt(code, "-e")
			if err != nil {
				return "", err
			}
		}

		if r.gofmt {
			var err error
			{
				code, err = r.applyGofmt(code, r.gofmtOptions...)
				if err != nil {
					return "", err
				}
			}
		}

		if r.goimports {
			var err error
			{
				code, err = r.applyGoimports(code)
				if err != nil {
					return "", err
				}
			}
		}
	}

	return code, nil
}

func (r *Root) applyGofmt(generatedCode string, gofmtOptions ...string) (string, error) {
	return applyCodeFormatter(generatedCode, "gofmt", gofmtOptions...)
}

func (r *Root) applyGoimports(generatedCode string) (string, error) {
	return applyCodeFormatter(generatedCode, "goimports")
}

func applyCodeFormatter(generatedCode string, formatterCmdName string, options ...string) (string, error) {
	var (
		formattedCmd = exec.Command(formatterCmdName, options...)
		stdinPipe, _ = formattedCmd.StdinPipe()
		out, errout  bytes.Buffer
	)

	formattedCmd.Stderr = &errout
	formattedCmd.Stdout = &out

	if err := formattedCmd.Start(); err != nil {
		cmds := []string{formatterCmdName}
		return "", kenerrs.CodeFormatterErrors(strings.Join(append(cmds, options...), " "), errout.String(), err)
	}

	_, err := io.WriteString(stdinPipe, generatedCode)
	if err != nil {
		return "", err
	}

	if err = stdinPipe.Close(); err != nil {
		return "", err
	}

	if err = formattedCmd.Wait(); err != nil {
		cmds := []string{formatterCmdName}
		return "", kenerrs.CodeFormatterErrors(strings.Join(append(cmds, options...), " "), errout.String(), err)
	}
	return out.String(), nil
}
