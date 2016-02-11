package xerrors

/*
* Work in progress (custom) error handling that aims to generalize the error types for easier error handling 
*/

import (
	"runtime"
	"strconv"
	)

var sysErrMsg = "Application Error!"
/*
* The system errors are FATAL, UNRECOVERABLE errors in the BUSINESS logic of the application. The system error should be considered as an ASSERT.
* The CURRENT executed job should be aborted.
* When a system error appears it means that there is a BUG in the code. 
* A general message should be displayed to the end user.
* The debug_msg is for the developer only and should contain the origin of the error (eg filename and line number).
*/

type SysErr struct {
    debugMsg string
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
*/
type PeerErr struct{
    source  string
}

func (e PeerErr) Error() string {
    return  e.source + `: Peer Error!`
}

func NewPeerErr(source string) PeerErr{
    return PeerErr{source}
}

/*
* The UI errors may be FATAL, but can also be RECOVERABLE. It is should be considered as an EXCEPTION. 
* If the current executed job is aborted depends on the error itself.
* The UI error's ui_msg should be shown to the user. In some cases the message of the actual error may not be user friendly(eg "No disk on device!")
* and the developer may want to mask it for example to "Not enough system resources. Please try again later!". In this case the debug_msg should contain the
* REAL error and the ui_msg should contain the USER FRIENDLY error.
*/
type UiErr struct{
        uiMsg	    string
	debugMsg   string
}

func (e UiErr) Error() string {
        return  e.uiMsg
}

func NewUiErr(uiMsg string, debugMsg string) UiErr{
        return UiErr{uiMsg, debugMsg}
}

