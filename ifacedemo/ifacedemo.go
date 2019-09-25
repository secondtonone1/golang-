package main

import "fmt"

type EmpInter interface {
	GetNum() int
}

type EmpStruct struct {
	num int
}

func (es *EmpStruct) GetNum() int {
	return es.num
}

func main() {

	emps := EmpStruct{num: 1}
	var empi EmpInter
	empi = &emps
	fmt.Println(empi)
	fmt.Println(emps)

}

//go build -gcflags "-l" -o ifacedemo  ifacedemo.go
//go tool objdump -s "main\.main" ifacedemo
/*
ifacedemo.go:17       0x48ffb0                65488b0c2528000000      MOVQ GS:0x28, CX
ifacedemo.go:17       0x48ffb9                488b8900000000          MOVQ 0(CX), CX
ifacedemo.go:17       0x48ffc0                483b6110                CMPQ 0x10(CX), SP
ifacedemo.go:17       0x48ffc4                0f86c1000000            JBE 0x49008b
ifacedemo.go:17       0x48ffca                4883ec50                SUBQ $0x50, SP
ifacedemo.go:17       0x48ffce                48896c2448              MOVQ BP, 0x48(SP)
ifacedemo.go:17       0x48ffd3                488d6c2448              LEAQ 0x48(SP), BP
ifacedemo.go:19       0x48ffd8                488d0521c30100          LEAQ runtime.types+111360(SB), AX
ifacedemo.go:19       0x48ffdf                48890424                MOVQ AX, 0(SP)
ifacedemo.go:19       0x48ffe3                e868b3f7ff              CALL runtime.newobject(SB)
ifacedemo.go:19       0x48ffe8                488b442408              MOVQ 0x8(SP), AX
ifacedemo.go:19       0x48ffed                4889442430              MOVQ AX, 0x30(SP)
ifacedemo.go:19       0x48fff2                48c70001000000          MOVQ $0x1, 0(AX)
ifacedemo.go:22       0x48fff9                488b0d48ce0400          MOVQ go.itab.*main.EmpStruct,main.EmpInter+8(SB), CX
ifacedemo.go:22       0x490000                0f57c0                  XORPS X0, X0
ifacedemo.go:22       0x490003                0f11442438              MOVUPS X0, 0x38(SP)
ifacedemo.go:22       0x490008                48894c2438              MOVQ CX, 0x38(SP)
ifacedemo.go:22       0x49000d                4889442440              MOVQ AX, 0x40(SP)
ifacedemo.go:22       0x490012                488d4c2438              LEAQ 0x38(SP), CX
ifacedemo.go:22       0x490017                48890c24                MOVQ CX, 0(SP)
ifacedemo.go:22       0x49001b                48c744240801000000      MOVQ $0x1, 0x8(SP)
ifacedemo.go:22       0x490024                48c744241001000000      MOVQ $0x1, 0x10(SP)
ifacedemo.go:22       0x49002d                e8de98ffff              CALL fmt.Println(SB)
ifacedemo.go:23       0x490032                488b442430              MOVQ 0x30(SP), AX
ifacedemo.go:23       0x490037                488b00                  MOVQ 0(AX), AX
ifacedemo.go:23       0x49003a                48890424                MOVQ AX, 0(SP)
ifacedemo.go:23       0x49003e                e88d89f7ff              CALL runtime.convT64(SB)
ifacedemo.go:23       0x490043                488b442408              MOVQ 0x8(SP), AX
ifacedemo.go:23       0x490048                0f57c0                  XORPS X0, X0
ifacedemo.go:23       0x49004b                0f11442438              MOVUPS X0, 0x38(SP)
ifacedemo.go:23       0x490050                488d0da9c20100          LEAQ runtime.types+111360(SB), CX
ifacedemo.go:23       0x490057                48894c2438              MOVQ CX, 0x38(SP)
ifacedemo.go:23       0x49005c                4889442440              MOVQ AX, 0x40(SP)
ifacedemo.go:23       0x490061                488d442438              LEAQ 0x38(SP), AX
ifacedemo.go:23       0x490066                48890424                MOVQ AX, 0(SP)
ifacedemo.go:23       0x49006a                48c744240801000000      MOVQ $0x1, 0x8(SP)
ifacedemo.go:23       0x490073                48c744241001000000      MOVQ $0x1, 0x10(SP)
ifacedemo.go:23       0x49007c                e88f98ffff              CALL fmt.Println(SB)
ifacedemo.go:25       0x490081                488b6c2448              MOVQ 0x48(SP), BP
ifacedemo.go:25       0x490086                4883c450                ADDQ $0x50, SP
ifacedemo.go:25       0x49008a                c3                      RET
ifacedemo.go:17       0x49008b                e860f2fbff              CALL runtime.morestack_noctxt(SB)
ifacedemo.go:17       0x490090                e91bffffff              JMP main.main(SB)
:-1                   0x490095                cc                      INT $0x3
:-1                   0x490096                cc                      INT $0x3
:-1                   0x490097                cc                      INT $0x3
:-1                   0x490098                cc                      INT $0x3
:-1                   0x490099                cc                      INT $0x3
:-1                   0x49009a                cc                      INT $0x3
:-1                   0x49009b                cc                      INT $0x3
:-1                   0x49009c                cc                      INT $0x3
:-1                   0x49009d                cc                      INT $0x3
:-1                   0x49009e                cc                      INT $0x3
:-1                   0x49009f                cc                      INT $0x3
*/
