package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	cli "gopkg.in/urfave/cli.v2"

	"github.com/geniusrabbit/f2go"
)

var (
	commit     = ""
	appVersion = "develop"
)

const (
	filePermissions = 0666
)

func main() {
	app := cli.App{
		Name:      "f2go",
		Version:   appVersion,
		Usage:     "converts any file to golang",
		UsageText: "",
		Metadata: map[string]interface{}{
			"commig": commit,
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "bytes",
				Usage: "convert to the bytes array",
			},
			&cli.StringFlag{
				Name:  "template",
				Value: "data.go",
			},
			&cli.StringFlag{
				Name:  "package",
				Usage: "Custom name of the package",
			},
			&cli.StringFlag{
				Name:  "varname",
				Usage: "Custom name of the variable",
			},
			&cli.BoolFlag{
				Name:  "f",
				Usage: "Save result to the file",
			},
		},
		ArgsUsage: "source_file [destination file]",
		Action:    convertFile,
	}

	if err := app.Run(os.Args); err != nil {
		log.Printf("[ERROR] %v", err)
	}
}

func convertFile(ctx *cli.Context) (err error) {
	srcFile := ctx.Args().First()
	dstFile := ctx.Args().Get(1)

	if srcFile == "" {
		return fmt.Errorf("undefined souce file")
	}

	// Prepare source absolute path
	if srcFile, err = abspath(srcFile, ""); err != nil {
		return err
	}

	// Open source file
	source, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer source.Close()

	var destination io.Writer
	if ctx.Bool("f") || dstFile != "" {
		if dstFile == "" {
			dstFile = srcFile + ".go"
		}

		// Prepare destination absolute path
		if dstFile, err = abspath(dstFile, filepath.Dir(srcFile)); err != nil {
			return err
		}

		// Open destination file
		openFlags := os.O_RDWR | os.O_CREATE | os.O_TRUNC
		destination, err = os.OpenFile(dstFile, openFlags, filePermissions)
		if err != nil {
			return err
		}
		defer destination.(io.Closer).Close()
	} else {
		destination = os.Stdout
	}

	var converter *f2go.Converter
	if ctx.Bool("bytes") {
		converter = f2go.NewConverter(source, destination, f2go.ByteEncoder)
	} else {
		converter = f2go.NewConverter(source, destination, f2go.StringEncoder)
	}

	// Create context options
	renderCtx := &f2go.RenderContext{
		PackageName: ctx.String("package"),
		DataName:    ctx.String("varname"),
	}

	if renderCtx.PackageName == "" {
		path := dstFile
		if path == "" {
			path = srcFile
		}
		renderCtx.PackageName = filepath.Base(filepath.Dir(path))
	}
	if renderCtx.DataName == "" {
		renderCtx.DataName = f2go.VariableNamePrepare(filepath.Base(srcFile))
	}
	return converter.Render(ctx.String("template"), renderCtx)
}

func abspath(fileame, wd string) (_ string, err error) {
	if !filepath.IsAbs(fileame) {
		if wd == "" {
			if wd, err = os.Getwd(); err != nil {
				return "", err
			}
		}
		return filepath.Join(wd, fileame), nil
	}
	return fileame, nil
}
