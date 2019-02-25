package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"io/ioutil"
	log "github.com/c2nc/protoc-go-inject-tag/logger"
	"os"
	"path"
	"regexp"
)

const (
	AppName = "protoc-go-inject-field"
	AppVer = "1.0.1"
	AppDescr = "Protobuf custom fields"
)

var rComment = regexp.MustCompile(`//\s*@inject_field:\s+(\S+)\s+(\S+)$`)

// customField is the information of the custom customField to inject to struct.
type customField struct {
	fieldName string
	fieldType string
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
		log.Fatal("input file is mandatory")
	}
	
	areas, err := parseFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	if err = writeFile(inputFile, areas); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.SetLogLevel("debug")
}

func main() {
	app := cli.App(AppName, AppDescr)
	app.Version("V version", fmt.Sprintf("%s v%s", AppName, AppVer))
	
	var (
		input = app.StringOpt("I input", ".", "Input directory")
	)
	
	app.Action = func() {
		content, err := ioutil.ReadDir(*input)
		if err != nil { log.Fatalf("Read file error %v", err) }
		
		for _, f := range content {
			fpath := path.Join(*input, f.Name())
			if path.Ext(fpath) == ".go" {
				processFile(fpath)
			}
		}
	}
	
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("%s error: %v", AppName, err)
	}
}