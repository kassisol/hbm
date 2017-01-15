package image

import (
	"fmt"
	"strings"
)

type Image struct {
	ID       string
	Registry string
	Name     string
	Tag      string
	Official bool
}

func NewImage(img string) Image {
	image := Image{}

	name, tag := GetNameTag(img)

	result := strings.Split(name, "/")
	count := len(result)

	if count >= 3 {
		image = Image{
			Registry: result[0],
			Name:     strings.Join(result[1:count], "/"),
			Tag:      tag,
			Official: false,
		}
	} else if count == 2 {
		image = Image{
			Name:     name,
			Tag:      tag,
			Official: false,
		}
	} else if count == 1 {
		image = Image{
			Name:     result[0],
			Tag:      tag,
			Official: true,
		}
	}

	return image
}

func (img *Image) String() string {
	if img.Registry != "" && img.Name != "" {
		return fmt.Sprintf("%s/%s", img.Registry, img.Name)
	} else if img.Name != "" {
		return fmt.Sprintf("%s", img.Name)
	}

	return fmt.Sprintf("%s", img.Name)
}

func GetNameTag(name string) (string, string) {
	nt := strings.SplitN(name, ":", 2)
	count := len(nt)

	if count == 2 {
		return nt[0], nt[1]
	} else if count == 1 {
		return nt[0], ""
	}

	return "", ""
}
