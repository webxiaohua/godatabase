package pkg1

import "fmt"

const (
	PKG1_ID int64 = 1
)

var Pkg1TypeInstance = &Pkg1Type{Id:1000}

type Pkg1Type struct {
	Id int64
}

func init()  {
	fmt.Println("pkg1 init...")
}