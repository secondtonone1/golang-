package main

import "fmt"

type EmpInter interface {
}

type EmpStruct struct {
	num int
}

func main() {

	emps := EmpStruct{num: 1}
	var empi EmpInter
	empi = emps
	fmt.Println(empi)
	fmt.Println(emps)

}

//go build -gcflags "-l" -o efacedemo  efacedemo.go
//go tool objdump -s "main\.main" efacedemo

/*
  efacedemo.go:12       0x48ffa0                65488b0c2528000000      MOVQ GS:0x28, CX
  efacedemo.go:12       0x48ffa9                488b8900000000          MOVQ 0(CX), CX
  efacedemo.go:12       0x48ffb0                483b6110                CMPQ 0x10(CX), SP
  efacedemo.go:12       0x48ffb4                0f86ae000000            JBE 0x490068
  efacedemo.go:12       0x48ffba                4883ec48                SUBQ $0x48, SP
  efacedemo.go:12       0x48ffbe                48896c2440              MOVQ BP, 0x40(SP)
  efacedemo.go:12       0x48ffc3                488d6c2440              LEAQ 0x40(SP), BP
  efacedemo.go:16       0x48ffc8                48c7042401000000        MOVQ $0x1, 0(SP)
  efacedemo.go:16       0x48ffd0                e8fb89f7ff              CALL runtime.convT64(SB)  //注意这里调用runtime包的convT64函数
  efacedemo.go:16       0x48ffd5                488b442408              MOVQ 0x8(SP), AX
  efacedemo.go:17       0x48ffda                0f57c0                  XORPS X0, X0
  efacedemo.go:17       0x48ffdd                0f11442430              MOVUPS X0, 0x30(SP)
  efacedemo.go:17       0x48ffe2                488d0d17c20100          LEAQ runtime.types+111104(SB
  efacedemo.go:17       0x48ffe9                48894c2430              MOVQ CX, 0x30(SP)
  efacedemo.go:17       0x48ffee                4889442438              MOVQ AX, 0x38(SP)
  efacedemo.go:17       0x48fff3                488d442430              LEAQ 0x30(SP), AX
  efacedemo.go:17       0x48fff8                48890424                MOVQ AX, 0(SP)
  efacedemo.go:17       0x48fffc                48c744240801000000      MOVQ $0x1, 0x8(SP)
  efacedemo.go:17       0x490005                48c744241001000000      MOVQ $0x1, 0x10(SP)
  efacedemo.go:17       0x49000e                e8fd98ffff              CALL fmt.Println(SB)
  efacedemo.go:18       0x490013                48c7042401000000        MOVQ $0x1, 0(SP)
  efacedemo.go:18       0x49001b                e8b089f7ff              CALL runtime.convT64(SB)   //注意这里调用runtime包的convT64函数
  efacedemo.go:18       0x490020                488b442408              MOVQ 0x8(SP), AX
  efacedemo.go:18       0x490025                0f57c0                  XORPS X0, X0
  efacedemo.go:18       0x490028                0f11442430              MOVUPS X0, 0x30(SP)
  efacedemo.go:18       0x49002d                488d0dccc10100          LEAQ runtime.types+111104(SB
  efacedemo.go:18       0x490034                48894c2430              MOVQ CX, 0x30(SP)
  efacedemo.go:18       0x490039                4889442438              MOVQ AX, 0x38(SP)
  efacedemo.go:18       0x49003e                488d442430              LEAQ 0x30(SP), AX
  efacedemo.go:18       0x490043                48890424                MOVQ AX, 0(SP)
  efacedemo.go:18       0x490047                48c744240801000000      MOVQ $0x1, 0x8(SP)
  efacedemo.go:18       0x490050                48c744241001000000      MOVQ $0x1, 0x10(SP)
  efacedemo.go:18       0x490059                e8b298ffff              CALL fmt.Println(SB)
  efacedemo.go:20       0x49005e                488b6c2440              MOVQ 0x40(SP), BP
  efacedemo.go:20       0x490063                4883c448                ADDQ $0x48, SP
  efacedemo.go:20       0x490067                c3                      RET
  efacedemo.go:12       0x490068                e883f2fbff              CALL runtime.morestack_noctx
  efacedemo.go:12       0x49006d                e92effffff              JMP main.main(SB)
*/
