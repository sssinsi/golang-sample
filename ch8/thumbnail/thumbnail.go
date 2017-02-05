package thumbnail

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//ImageFile はinfileから画像を読み込み、同じディレクトリにその画像のサムネイルサイズの画像を書き出す。
//ImageFile は生成されたファイル名を返します。例："foo.thumb.jpg"
func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile) //e.g: ".jpg", ".JPEG".拡張子を取得
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}

// Image returns a thmbnail-size version of src.
func Image(src image.Image) image.Image {
	//Compute thumbnail size, preserving aspect ratio.
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect) //portrait
	} else {
		height = int(128 / aspect) //landscape
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)

}
func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}
	return out.Close()
}
