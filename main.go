package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/mbndr/figlet4go"
	"gopkg.in/yaml.v3"
)

func main() {
	fmt.Print(showBanner())
	fmt.Printf("v1.0.0\n\n")
	fmt.Println("+--------------------------------------------------------------+\n")
	fmt.Println(Channa(readConfigFile()))
}

func showBanner() string {
	ascii := figlet4go.NewAsciiRender()

	options := figlet4go.NewRenderOptions()
	options.FontName = "larry3d"
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorRed,
		figlet4go.ColorBlue,
		figlet4go.ColorCyan,
		figlet4go.ColorMagenta,
	}

	renderStr, _ := ascii.RenderOpts("Channa", options)

	return renderStr
}

func readConfigFile() map[string]interface{} {
	filename := flag.String("f", "", "must fill filename want to convert to")
	flag.Parse()

	if filename == nil {
		return nil
	}

	fi, err := os.ReadFile("./resources/" + *filename)
	if err != nil {
		panic(err)
	}

	src := make(map[string]interface{})

	if err := json.Unmarshal(fi, &src); err != nil {
		if err = yaml.Unmarshal(fi, &src); err != nil {
			return nil
		}
	}

	return src
}
