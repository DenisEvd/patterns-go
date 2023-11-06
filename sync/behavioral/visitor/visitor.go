package visitor

import "strings"

type Visitor interface {
	VisitPark(park *Park) string
	VisitMarket(market *Market) string
	VisitZoo(zoo *Zoo) string
}

type People struct{}

func (p *People) VisitPark(park *Park) string {
	return park.MakePhoto()
}

func (p *People) VisitMarket(market *Market) string {
	return market.BuyFruits()
}

func (p *People) VisitZoo(zoo *Zoo) string {
	return zoo.FeedAnimals()
}

type Place interface {
	Accept(v Visitor) string
}

type City struct {
	places []Place
}

func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

func (c *City) Accept(v Visitor) string {
	sb := strings.Builder{}

	for _, e := range c.places {
		sb.WriteString(e.Accept(v) + "\n")
	}

	return sb.String()
}

type Park struct{}

func (p *Park) Accept(v Visitor) string {
	return v.VisitPark(p)
}

func (p *Park) MakePhoto() string {
	return "Make photo in the park"
}

type Market struct{}

func (m *Market) Accept(v Visitor) string {
	return v.VisitMarket(m)
}

func (m *Market) BuyFruits() string {
	return "Buy fruits"
}

type Zoo struct{}

func (z *Zoo) Accept(v Visitor) string {
	return v.VisitZoo(z)
}

func (z *Zoo) FeedAnimals() string {
	return "Feed lion"
}
