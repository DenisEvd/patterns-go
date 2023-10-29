package composite

import "errors"

var ErrCanNotAddComponentToLeave = errors.New("can not add component to the leave")

type Component interface {
	Add(child Component) error
	Child() []Component

	Path(name string) string
}

type Directory struct {
	name  string
	child []Component
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) Add(child Component) error {
	d.child = append(d.child, child)

	return nil
}

func (d *Directory) Child() []Component {
	return d.child
}

func (d *Directory) Path(name string) string {
	curPart := d.name + "/"
	if d.name == name {
		return curPart
	}

	var path string
	for _, e := range d.child {
		path = e.Path(name)
	}

	if path == "" {
		return ""
	}

	return curPart + path
}

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{name: name}
}

func (f *File) Add(child Component) error {
	return ErrCanNotAddComponentToLeave
}

func (f *File) Child() []Component {
	return []Component{}
}

func (f *File) Path(name string) string {
	if f.name == name {
		return name
	}

	return ""
}
