package builder

import (
	"errors"
	"strings"
)

var ErrEmptyArgument = errors.New("string argument is empty")

type Builder interface {
	WithHeader(header string) error
	WithText(text string)
	WithFooter(footer string) error

	Build() Product
}

type ProductBuilder struct {
	product *Product
}

func NewProductBuilder() ProductBuilder {
	return ProductBuilder{product: &Product{}}
}

func (p *ProductBuilder) WithHeader(header string) error {
	if header == "" {
		return ErrEmptyArgument
	}

	p.product.Header = header

	return nil
}

func (p *ProductBuilder) WithText(text string) {
	p.product.Text += text
}

func (p *ProductBuilder) WithFooter(footer string) error {
	if footer == "" {
		return ErrEmptyArgument
	}

	p.product.Footer = footer

	return nil
}

func (p *ProductBuilder) Build() *Product {
	return p.product
}

type Product struct {
	Header string
	Text   string
	Footer string
}

func (p *Product) ToString() string {
	strBuilder := strings.Builder{}

	strBuilder.WriteString(p.Header + "\n")
	strBuilder.WriteString(p.Text + "\n")
	strBuilder.WriteString(p.Footer + "\n")

	return strBuilder.String()
}
