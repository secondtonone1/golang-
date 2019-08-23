package serviceconfig

import "errors"

const (
	Authaddress = "localhost:50051"

	Loginaddress = "localhost:50052"

	Registeraddress = "localhost:50053"

	DBaddress = "localhost:50054"

	SAVEGOROUTINE_NUM = 2
	SAVECHAN_SIZE     = 1024
)

const (
	RSP_SUCCESS       = 1
	RSP_ACTNOTREG     = 2
	RSP_ACTHASREG     = 3
	RSP_UNKOWNERR     = 4
	RSP_SAVEMSGERR    = 5
	RSP_PROTOMARSHERR = 6
)

var (
	ErrAccountNotReg        = errors.New("Account wasn't register")
	ErrAccountHasReg        = errors.New("Account has been registered")
	ErrUnknowErr            = errors.New("Unknown Error!")
	ErrAuthServerInit       = errors.New("AuthServer init failed")
	ErrDBMgrInit            = errors.New("DB Manager init failed")
	ErrDBGetValue           = errors.New("DBGetValue failed")
	ErrDBPutValue           = errors.New("ErrDBPutValue Failed")
	ErrAllSaveRoutinesClose = errors.New("All save routines are closed")
	ErrDBServerInit         = errors.New("DBServer init failed")
)
