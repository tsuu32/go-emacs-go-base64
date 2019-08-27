package main

// int plugin_is_GPL_compatible;
import "C"

import (
	"encoding/base64"
	"github.com/sigma/go-emacs"
)

func init() {
	emacs.Register(initModule)
}

func initModule(env emacs.Environment) {
	env.RegisterFunction("go-emacs-go-base64-encode", B64Encode, 1, "Return base64 encoded string.", nil)
	env.RegisterFunction("go-emacs-go-base64-decode", B64Decode, 1, "Return base64 decoded string.", nil)
	env.ProvideFeature("go-emacs-go-base64")
}

func B64Encode(ctx emacs.FunctionCallContext) (emacs.Value, error) {
	env := ctx.Environment()
	stdlib := ctx.Environment().StdLib()

	s, err := ctx.GoStringArg(0)
	if err != nil {
		return stdlib.Nil(), err
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(s))

	return env.String(encoded), nil
}

func B64Decode(ctx emacs.FunctionCallContext) (emacs.Value, error) {
	env := ctx.Environment()
	stdlib := ctx.Environment().StdLib()

	s, err := ctx.GoStringArg(0)
	if err != nil {
		return stdlib.Nil(), err
	}

	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return stdlib.Nil(), err
	}

	return env.String(string(decoded)), nil
}

func main() {}
