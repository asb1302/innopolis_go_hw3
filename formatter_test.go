package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddFormatter(t *testing.T) {
	ast := assert.New(t)
	chainFormatter := NewChainFormatter()

	chainFormatter.AddFormatter(BoldFormatter{})

	ast.Contains(chainFormatter.formatters, BoldFormatter{})
}

func TestFormat(t *testing.T) {
	ast := assert.New(t)
	chainFormatter := NewChainFormatter()
	chainFormatter.AddFormatter(PlainTextFormatter{})
	chainFormatter.AddFormatter(BoldFormatter{})
	chainFormatter.AddFormatter(CodeFormatter{})
	chainFormatter.AddFormatter(ItalicFormatter{})
	text := "Пример текста для форматирования"

	formattedText := chainFormatter.Format(text)

	ast.Equal("_`**Пример текста для форматирования**`_", formattedText)
}
