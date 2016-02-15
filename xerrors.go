package xerrors

/*
* Work in progress (custom) error handling that aims to generalize the error types for easier error handling 
*/

import (
	"runtime"
	"strconv"
	)

var sysErrMsg = "Application Error!"
var peerErrMsg = "Peer Error!"

/*
* The system errors are FATAL, UNRECOVERABLE errors in the BUSINESS logic of the application. The system error should be considered as an ASSERT.
* The CURRENT executed job should be aborted.
* When a system error appears it means that there is a BUG in the code. 
* A general message should be displayed to the end user.
* The debugMsg is for the developer only and should contain the origin of the error (eg filename and line number).
*/

type SysErr struct {
    debugMsg string
}


func (e SysErr) GetDebugMsg() string{
    return e.debugMsg
}

func (e SysErr) Error() string {
    return  sysErrMsg
}

func NewSysErr() SysErr{
    _, fn, line, _ := runtime.Caller(1)
    return SysErr{sysErrMsg + ` file name: ` + fn + ` line: ` + strconv.Itoa(line) }
}

/*
* The peer errors are FATAL, UNRECOVERABLE errors in a REMOTE system.
* The CURRENT executed job should be aborted.
* When a Peer error apears it means that the remote system is down.
* A general message should be displayed to the end user with the name (or unique identifier) of the remote system
* debugMsg is for the developer so that the error can be handled easier.
*/

type PeerErr struct{
    debugMsg  string
}


func (e PeerErr) Error() (string) {
    return  peerErrMsg
}

func (e PeerErr) GetDebugMsg() (string){
    return e.debugMsg
}

func NewPeerErr(debugMsg string) PeerErr{
    return PeerErr{debugMsg}
}

/*
* The UI errors may be FATAL, but can also be RECOVERABLE. It is should be considered as an EXCEPTION. 
* If the current executed job is aborted depends on the error itself.
*/

type UIErr struct{
        uiMsg		string // Message that should be displayed to the user, saying what the user should do
	debugMsg	string // Message for the developer
	code		string
	IsRetryable	bool
}

func (e UIErr) Error() (string) {
        return  e.uiMsg
}

func (e UIErr) GetDebugMsg() (string, string){
    return e.debugMsg, e.code
}

func NewUIErr(uiMsg string, debugMsg string, code string, IsRetryable bool) UIErr{
        return UIErr{uiMsg, debugMsg, code, IsRetryable}
}


