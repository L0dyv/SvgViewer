package main

import (
	_ "embed"

	"golang.design/x/clipboard"

	webview "github.com/webview/webview_go"
)

//go:embed index.html
var html string

func main() {
	// 初始化剪贴板
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("SVG 预览器")
	w.SetSize(1200, 800, webview.HintNone)

	// 绑定 Go 函数给 JS 调用，绕过 WebView 的剪贴板安全限制
	w.Bind("getClipboardText", func() string {
		return string(clipboard.Read(clipboard.FmtText))
	})

	w.SetHtml(html)
	w.Run()
}
