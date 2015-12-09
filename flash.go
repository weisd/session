package session

// import (
// 	"github.com/labstack/echo"
// 	"net/url"
// )

// ___________.____       _____    _________ ___ ___
// \_   _____/|    |     /  _  \  /   _____//   |   \
//  |    __)  |    |    /  /_\  \ \_____  \/    ~    \
//  |     \   |    |___/    |    \/        \    Y    /
//  \___  /   |_______ \____|__  /_______  /\___|_  /
//      \/            \/       \/        \/       \/

// type Flash struct {
// 	ctx *echo.Context
// 	url.Values
// 	ErrorMsg, WarningMsg, InfoMsg, SuccessMsg string
// }

// func (f *Flash) set(name, msg string, current ...bool) {

// 	f.Set(name, msg)

// }

// func (f *Flash) Error(msg string, current ...bool) {
// 	f.ErrorMsg = msg
// 	f.set("error", msg, current...)
// }

// func (f *Flash) Warning(msg string, current ...bool) {
// 	f.WarningMsg = msg
// 	f.set("warning", msg, current...)
// }

// func (f *Flash) Info(msg string, current ...bool) {
// 	f.InfoMsg = msg
// 	f.set("info", msg, current...)
// }

// func (f *Flash) Success(msg string, current ...bool) {
// 	f.SuccessMsg = msg
// 	f.set("success", msg, current...)
// }
