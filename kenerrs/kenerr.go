package kenerrs

import "fmt"

func CodeFormatterErrors(cmd string, msg string, fmterr error) error {
	return fmt.Errorf(`[MEW-1] code formatter raise error: command=%s, err=%s, msg=%s`, cmd, msg, fmterr)
}

func FuncReceiverEmptyNameError(caller string) error {
	return fmt.Errorf(`[MEW-2] Function reciever raise error: receiver has empty name, %s`, caller)
}

func FuncReceiverEmptyTypeError(caller string) error {
	return fmt.Errorf(`[MEW-2] Function reciever raise error: receiver has empty type, %s`, caller)
}

func FuncSignatureEmptyFuncNameError(caller string) error {
	return fmt.Errorf(`[MEW-3] Function signarute raise error: func name is empty, %s`, caller)
}

func FuncParamNameEmptyError(caller string) error {
	return fmt.Errorf(`[MEW-4] Function param raise error: empty name param, %s`, caller)
}

func LastParameterTypeIsEmptyError(caller string) error {
	return fmt.Errorf(`[MEW-5] Function param raise error: last param type is empty, %s`, caller)
}

func UnnamedReturnTypeAppearsAfterNamedReturnTypeError(caller string) error {
	return fmt.Errorf(`[MEW-6] Function signature raise error: unnamed return type appears after named return type, %s`, caller)
}

func FuncSignatureIsNilError(caller string) error {
	return fmt.Errorf(`[MEW-7] Function signature raise error: signature is nil, %s`, caller)
}

func FuncInvocationParameterIsEmptyError(caller string) error {
	return fmt.Errorf(`[MEW-8] Function invocation raise error: param is empty, %s`, caller)
}

func AnonFuncSignatureIsNilError(caller string) error {
	return fmt.Errorf(`[MEW-9] Anon function signature raise error: signature is nil, %s`, caller)
}
