package template_method

type QuotesInterface interface {
	Open() string
	Close() string
}

type Quotes struct {
	QuotesInterface
}

func (q *Quotes) Quotes(str string) string {
	return q.Open() + str + q.Close()
}

func NewQuotes(qt QuotesInterface) *Quotes {
	return &Quotes{QuotesInterface: qt}
}

type FrenchQuotes struct{}

func (f *FrenchQuotes) Open() string {
	return "«"
}

func (f *FrenchQuotes) Close() string {
	return "»"
}

type GermanQuotes struct{}

func (g *GermanQuotes) Open() string {
	return "„"
}

func (g *GermanQuotes) Close() string {
	return "“"
}
