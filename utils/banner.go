package utils

import (
	"fmt"
	"github.com/mbndr/figlet4go"
)

func Banner(message string) {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorYellow,
		figlet4go.ColorYellow,
	}

	renderStr, _ := ascii.RenderOpts(message, options)
	fmt.Print(renderStr)
}

func BannerColor(message string, colors ...figlet4go.AnsiColor) {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	var list []figlet4go.Color
	for _, c := range colors {
		list = append(list, c)
	}
	options.FontColor = list
	renderStr, _ := ascii.RenderOpts(message, options)
	fmt.Print(renderStr)
}
