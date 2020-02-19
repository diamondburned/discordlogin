package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zserge/webview"
)

const JavaScript = `
	let setRequestHeader = XMLHttpRequest.prototype.setRequestHeader
	let isAuth = (key, value) => {
		return key == "Authorization" && value && !value.startsWith("Bearer")
	 }

	XMLHttpRequest.prototype.setRequestHeader = function() {
		if (isAuth(arguments[0], arguments[1])) {
			window.external.invoke(arguments[1])
			window.location = ""
			return
		}

		setRequestHeader.apply(this, arguments)
	}
`

func main() {
	wv := webview.New(webview.Settings{
		Title:     "Discord Login",
		URL:       "https://discordapp.com/login",
		Width:     1000,
		Height:    800,
		Resizable: true,

		ExternalInvokeCallback: receivedToken,
	})

	wv.Dispatch(func() {
		if err := wv.Eval(JavaScript); err != nil {
			log.Fatalln("Failed to evaluate JavaScript:", err)
		}
	})

	wv.Run()
}

func receivedToken(wv webview.WebView, token string) {
	wv.Dispatch(func() {
		wv.Terminate()
	})
	wv.Exit()

	fmt.Println(token)
	os.Exit(0)
}
