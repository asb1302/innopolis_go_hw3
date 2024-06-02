package main

type Formatter interface {
	Format(string) string
}

type PlainTextFormatter struct{}

func (p PlainTextFormatter) Format(text string) string {
	return text
}

type BoldFormatter struct{}

func (b BoldFormatter) Format(text string) string {
	return "**" + text + "**"
}

type CodeFormatter struct{}

func (c CodeFormatter) Format(text string) string {
	return "`" + text + "`"
}

type ItalicFormatter struct{}

func (i ItalicFormatter) Format(text string) string {
	return "_" + text + "_"
}

type ChainFormatter struct {
	formatters []Formatter
}

func NewChainFormatter() *ChainFormatter {
	return &ChainFormatter{}
}

func (cf *ChainFormatter) AddFormatter(formatter Formatter) {
	cf.formatters = append(cf.formatters, formatter)
}

func (cf *ChainFormatter) Format(text string) string {
	formattedText := text
	for _, formatter := range cf.formatters {
		formattedText = formatter.Format(formattedText)
	}
	return formattedText
}
