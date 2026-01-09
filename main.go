package main

import (
	_ "embed"

	webview "github.com/webview/webview_go"
)

//go:embed index.html
var html string

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("SVG 预览器")
	w.SetSize(1200, 800, webview.HintNone)
	w.SetHtml(html)
	w.Run()
}
