package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/JuanPerdomo00/glain/utils"
)

func main() {
	gifsDirectory := flag.String("gifs", "~/.config/glain/gifs", "GIF Directory")
	indexPath := flag.String("index", "~/.config/glain/.gif-index", "GIF Index Path")
	marginPath := flag.String("margin", "~/.config/glain/margin.txt", "Margin file path")
	flag.Parse()

	home, _ := os.UserHomeDir()
	*gifsDirectory = strings.Replace(*gifsDirectory, "~", home, 1)
	*indexPath = strings.Replace(*indexPath, "~", home, 1)
	*marginPath = strings.Replace(*marginPath, "~", home, 1)

	gifs := utils.LoadGifs(*indexPath)

	done := make(chan string)
	go func() {
		tmpf, err := os.CreateTemp("", "glain-*")
		if err != nil {
			log.Fatal(err)
		}
		cmd := exec.Command("fastfetch", "--logo", *marginPath)
		cmd.Stdout = tmpf
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		tmpf.Close()
		done <- tmpf.Name()
	}()

	gif := utils.PickGif(gifs)
	gifPath := *gifsDirectory + "/" + gif.Name

	tmpPath := <-done

	fmt.Print("\033[2J\033[H")

	rtmp, err := os.ReadFile(tmpPath)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Print(string(rtmp))
	os.Remove(tmpPath)
	fmt.Print("\033[H\033[2B")
	var proto string = utils.DetectProtocol()
	utils.RunTimg(gifPath, proto, gif.SetGeometry())

}
