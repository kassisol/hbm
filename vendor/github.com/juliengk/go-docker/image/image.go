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
	name, tag := GetNameTag(img)

	image := Image{
		Name: name,
		Tag:  tag,
	}

	result := strings.Split(name, "/")
	count := len(result)

	if count >= 3 {
		image.Registry = result[0]
		image.Name = strings.Join(result[1:count], "/")
	} else if count == 2 {
		if validateRegistry(result[0]) {
			image.Registry = result[0]
			image.Name = result[1]
		}
	} else if count == 1 {
		image.Official = true
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
