package main

import (
	"encoding/csv"
	"html/template"
	"os"
	"strings"
)

// index of cli arguments
const filenameIndex = 0

// template for index
const templateHTML = ".github/workflows/template.html"

// directory for GitHub Pages
const docsDir = "docs/"

// Context includes table header and body
type Context struct {
	Title  string
	Header []string
	Body   [][]string
}

func main() {
	// get csv filename from cli args
	source := getSourceFilename()
	// get cannonical filename w/o extension
	sourceID := getSourceID(source)

	csvFileReader := getSourceFileReader(sourceID)
	// close file after processing
	defer csvFileReader.Close()

	// create special csv reader
	csvReader := csv.NewReader(csvFileReader)

	// get populated context for template
	context := getContext(csvReader, sourceID)

	// get html file writer
	htmlWriter := getHTMLFileWriter(sourceID)
	// close file after processing
	defer htmlWriter.Close()

	// populate template with context
	processTemplate(htmlWriter, context)
}

func getSourceFilename() string {
	// get cli arguments without program name
	args := os.Args[1:]
	// return cli argument as .csv source filename
	return args[filenameIndex]
}

func getSourceID(source string) string {
	// remove last .csv
	str := strings.TrimSuffix(source, ".csv")

	return str
}

func getSourceFileReader(sourceID string) *os.File {
	// get full csv path
	source := getCSVFilename(sourceID)
	// save csv file reader
	csvFileReader, err := os.Open(source)
	// check reading errors
	if err != nil {
		panic(err)
	}

	return csvFileReader
}

func getCSVFilename(sourceID string) string {
	// concats directory path and csv extension
	return sourceID + ".csv"
}

func getContext(reader *csv.Reader, sourceID string) *Context {
	// create context for template
	context := &Context{Title: sourceID}

	// populate context
	addHeader(context, reader)
	addBody(context, reader)

	return context
}

func addHeader(context *Context, reader *csv.Reader) {
	// table header is the first row in csv file
	header, err := reader.Read()
	// check reading error
	if err != nil {
		panic(err)
	}

	// set header of future template
	context.Header = header
}

func addBody(context *Context, reader *csv.Reader) {
	// table body is other table rows
	body, err := reader.ReadAll()
	// check reading error
	if err != nil {
		panic(err)
	}

	// set body of future template
	context.Body = body
}

func getHTMLFileWriter(sourceID string) *os.File {
	// get full html file path
	destination := getHTMLFilename(sourceID)
	// create html file writer
	htmlWriter, err := os.Create(destination)
	// check file creation errors
	if err != nil {
		panic(err)
	}

	return htmlWriter
}

func getHTMLFilename(sourceID string) string {
	println("processing " + sourceID)
	// concats directory path and html extension
	return docsDir + sourceID + ".html"
}

func processTemplate(writer *os.File, context *Context) {
	// parse template
	tmpl, err := template.ParseFiles(templateHTML)
	// check parsing errors
	if err != nil {
		panic(err)
	}

	// create html file from saved context
	err = tmpl.Execute(writer, context)
	// check generation errors
	if err != nil {
		panic(err)
	}
}
