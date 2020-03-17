package main

import (
	"fmt"

	"github.com/zserge/webview"
)

const JavaScript = `
	// Clear local storage
	localStorage.clear()

	let setRequestHeader = XMLHttpRequest.prototype.setRequestHeader
	let isAuth = (key, value) => {
		return key == "Authorization" && value && !value.startsWith("Bearer")
	 }

	XMLHttpRequest.prototype.setRequestHeader = function() {
		if (isAuth(arguments[0], arguments[1])) {
			receivedToken(arguments[1])
			window.location = ""
			return
		}

		setRequestHeader.apply(this, arguments)
	}
`

func main() {
	wv := webview.New(false)
	wv.SetSize(1000, 800, webview.HintNone)
	wv.SetTitle("Discord Login")
	wv.Navigate("https://discordapp.com/login")
	wv.Init(JavaScript)

	err := wv.Bind("receivedToken", func(token string) {
		fmt.Print(token)

		wv.Dispatch(func() {
			wv.Terminate()
		})
	})

	if err != nil {
		panic(err)
	}

	wv.Run()
}
