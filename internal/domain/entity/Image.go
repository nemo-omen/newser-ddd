package entity

import (
	"encoding/json"
	"newser/internal/domain/value"
)

type Image struct {
	ID    value.ImageId
	Title value.Title
	URL   value.Link
}

type ImageProps struct {
	Title string
	URL   string
}

func NewImage(props ImageProps) *Image {
	validTitle, err := value.NewTitle(props.Title)
	if err != nil {
		return nil
	}

	validUrl, err := value.NewLink(props.URL)
	if err != nil {
		return nil
	}

	return &Image{
		ID:    value.ImageId{ID: value.NewId()},
		Title: validTitle,
		URL:   validUrl,
	}
}

func (i *Image) JSON() []byte {
	type jsonImage struct {
		ID    value.ImageId `json:"id"`
		Title value.Title   `json:"title"`
		URL   value.Link    `json:"url"`
	}
	ji := jsonImage{
		ID:    i.ID,
		Title: i.Title,
		URL:   i.URL,
	}
	jsi, _ := json.MarshalIndent(ji, "", "  ")
	return jsi
}

func (i *Image) String() string {
	return string(i.JSON())
}
