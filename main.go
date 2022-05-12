package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"

	ico "github.com/Kodeworks/golang-image-ico"
)

type icon struct {
	size int
	name string
	ext  string
}

var webappFilesSet = []icon{
	{ext: "ico", size: 16, name: "favicon"},
	{ext: "png", size: 256, name: "icon"},
	{ext: "png", size: 16, name: "favicon-16x16"},
	{ext: "png", size: 32, name: "favicon-32x32"},
	{ext: "png", size: 192, name: "android-chrome-192x192"},
	{ext: "png", size: 512, name: "android-chrome-512x512"},
	{ext: "png", size: 180, name: "apple-touch-icon"},
	{ext: "png", size: 60, name: "apple-touch-icon-60x60"},
	{ext: "png", size: 76, name: "apple-touch-icon-76x76"},
	{ext: "png", size: 120, name: "apple-touch-icon-120x120"},
	{ext: "png", size: 152, name: "apple-touch-icon-152x152"},
	{ext: "png", size: 180, name: "apple-touch-icon-180x180"},
	{ext: "png", size: 144, name: "msapplication-icon-144x144"},
	{ext: "png", size: 150, name: "mstile-150x150"},
}

func main() {
	fmt.Println("svg-to-pwa-icons v1.0.0 - https://github.com/sattellite/svg-to-pwa-icons")

	svg, err := tryOpenSVG()
	if err != nil {
		fmt.Printf("failed read file: %v\n", err)
		return
	}
	defer svg.Close()

	icon, err := oksvg.ReadIconStream(svg)
	if err != nil {
		fmt.Printf("failed parse svg: %v\n", err)
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed get cwd: %v\n", err)
		return
	}

	outdir := path.Join(cwd, "out")
	err = os.Mkdir(outdir, os.FileMode(0755))
	if err != nil {
		fmt.Printf("failed create output dir: %v\n", err)
		return
	}

	for _, f := range webappFilesSet {
		fmt.Printf("processing %s\n", f.name+"."+f.ext)
		err := createIcon(outdir, f, icon)
		if err != nil {
			fmt.Printf("failed draw icon: %v\n", err)
		}
	}

}

func tryOpenSVG() (*os.File, error) {
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("usage: %s <svg file>", os.Args[0])
	}
	filename := os.Args[1]
	if filename == "" {
		return nil, fmt.Errorf("no file specified")
	}
	return os.Open(filename)
}

func createIcon(outdir string, icon icon, svg *oksvg.SvgIcon) error {
	svg.SetTarget(0, 0, float64(icon.size), float64(icon.size))
	rgba := image.NewRGBA(image.Rect(0, 0, icon.size, icon.size))
	svg.Draw(rasterx.NewDasher(icon.size, icon.size, rasterx.NewScannerGV(icon.size, icon.size, rgba, rgba.Bounds())), 1)

	return writeToFile(outdir, icon, rgba)
}

func writeToFile(outdir string, icon icon, img image.Image) error {
	out, err := os.Create(path.Join(outdir, icon.name+"."+icon.ext))
	if err != nil {
		return err
	}
	defer out.Close()

	if icon.ext == "png" {
		err = png.Encode(out, img)
	} else if icon.ext == "ico" {
		err = ico.Encode(out, img)
	} else {
		err = fmt.Errorf("unknown file extension: %s", icon.ext)
	}
	if err != nil {
		return err
	}
	return nil
}
