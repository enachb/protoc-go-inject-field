package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"

	cli "github.com/jawher/mow.cli"
)

const (
	AppName  = "protoc-go-inject-field"
	AppVer   = "1.0.1"
	AppDescr = "Protobuf custom fields"
)

var rComment = regexp.MustCompile(`//\s*@inject_field:\s+(\S+)\s+(\S+)\s*(\S*)$`)

// customField is the information of the custom customField to inject to struct.
type customField struct {
	fieldName string
	fieldType string
	fieldTag  string
}

// textArea is information of the struct to inject custom fields to.
type textArea struct {
	name  string
	start int
	end   int
	// where to insert custom fields, should be after end of last field.
	// We have to record it because between the closing brace and the end
	// of the struct can contains comment, white spaces...
	insertPos int
	fields    []*customField
}

func processFile(inputFile string) {
	if len(inputFile) == 0 {
		ErrorLog.Fatal("input file is mandatory")
	}

	areas, err := parseFile(inputFile)
	if err != nil {
		ErrorLog.Fatal(err)
	}
	if err = writeFile(inputFile, areas); err != nil {
		ErrorLog.Fatal(err)
	}
}

func main() {
	app := cli.App(AppName, AppDescr)
	app.Version("V version", fmt.Sprintf("%s v%s", AppName, AppVer))

	var (
		input = app.StringOpt("I input", ".", "Input directory")
	)

	app.Action = func() {
		content, err := ioutil.ReadDir(*input)
		if err != nil {
			ErrorLog.Fatalf("Read file error %v", err)
		}

		for _, f := range content {
			fpath := path.Join(*input, f.Name())

			if len(f.Name()) > 7 {
				ext := string(f.Name()[len(f.Name())-5:])
				if ext == "pb.go" {
					processFile(fpath)
				}
			}
		}
	}

	if err := app.Run(os.Args); err != nil {
		ErrorLog.Fatalf("%s error: %v", AppName, err)
	}
}
