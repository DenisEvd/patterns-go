package flyweight

import "fmt"

type Image interface {
	Draw(width int, height int) string
}

type ImageFlyweightFactory struct {
	pool map[string]Image
}

func NewImageFlyweightFactory() *ImageFlyweightFactory {
	return &ImageFlyweightFactory{
		pool: make(map[string]Image),
	}
}

func (i *ImageFlyweightFactory) GetImage(filename string) Image {
	if _, ok := i.pool[filename]; !ok {
		i.pool[filename] = &ConcreteImage{filename: filename}
	}

	return i.pool[filename]
}

type ConcreteImage struct {
	filename string
}

func (c *ConcreteImage) Draw(width int, height int) string {
	return fmt.Sprintf("Draw image: %s\nWidth: %d Height: %d", c.filename, width, height)
}
