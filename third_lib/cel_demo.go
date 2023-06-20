package main

/*
语法规则：https://github.com/google/cel-spec/blob/master/doc/langdef.md
go get github.com/google/cel-go@v0.10.0
*/
import (
	"fmt"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func main() {
	env, err := cel.NewEnv(cel.Declarations(
		decls.NewVar("name", decls.String),
		decls.NewVar("group", decls.String),
	))
	if err != nil {
		fmt.Println("cel.NewEnv() error1:", err)
		return
	}
	ast, issues := env.Compile(`name.startsWith("/groups/"+group)`)
	if issues != nil && issues.Err() != nil {
		fmt.Println("error2:", issues.Err())
		return
	}
	prg, err := env.Program(ast)
	if err != nil {
		fmt.Println("error3:", err)
		return
	}
	out, _, err := prg.Eval(map[string]interface{}{
		"name":  "/groups/bilibili/123",
		"group": "bilibili",
	})
	if err != nil {
		fmt.Println("error4:", err)
		return
	}
	res := out.Value().(bool)
	fmt.Println("res:", res)
}
