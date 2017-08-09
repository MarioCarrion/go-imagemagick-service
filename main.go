// Port of http://www.imagemagick.org/MagickWand/resize.htm
//   from http://www.imagemagick.org/MagickWand/
package main

import (
	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	mw.ReadImage("imagemagick.jpg")

	width := mw.GetImageWidth() / 2
	height := mw.GetImageHeight() / 2

	mw.ResizeImage(width, height, imagick.FILTER_LANCZOS)
	mw.SetImageCompressionQuality(95)
	mw.WriteImage("output.jpg")
}
