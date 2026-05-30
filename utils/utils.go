package utils

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/JuanPerdomo00/glain/model"
)

func LoadGifs(indexPath string) []model.Gif {
	file, err := os.Open(indexPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var gifs []model.Gif

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		if len(line) < 3 {
			continue
		}
		width, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalf("width: %v", err)
		}

		height, err := strconv.Atoi(line[2])

		if err != nil {
			log.Fatalf("height: %v", err)
		}

		gifs = append(gifs, model.Gif{Name: line[0], Width: width, Height: height})
	}

	return gifs

}

func PickGif(gifs []model.Gif) model. Gif {
	return gifs[rand.Intn(len(gifs))]
}



func DetectProtocol() string {
	if os.Getenv("KITTY_WINDOW_ID") != "" {
		return "kitty"
	}
	term := os.Getenv("TERM")
	if term == "xterm-kitty" {
		return "kitty"
	}
	if term == "foot" || term == "foot-extra" {
		return "sixel"
	}
	if os.Getenv("TERM_PROGRAM") == "ghostty" {
		return "quarter"
	}
	return "sixel"
}

func RunTimg(gifPath, proto, geometry string) {
	timgPath, err := exec.LookPath("timg")
	if err != nil {
		log.Fatal(err)
	}

	errSyscall := syscall.Exec(
		timgPath,
		[]string{"timg", "-p", proto, "--loops", "0", "-g", geometry, gifPath},
		os.Environ(),
	)

	if errSyscall != nil {
		log.Fatal(errSyscall)
	}
}
