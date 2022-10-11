package schemas

import (
	"encoding/xml"
)

type Package struct {
	XMLName          xml.Name   `xml:"package"`
	XMLNS            string     `xml:"xmlns,attr"`
	Version          string     `xml:"version,attr"`
	UniqueIdentifier string     `xml:"unique-identifier,attr"`
	Direction        string     `xml:"dir,attr"`
	Language         string     `xml:"xml lang,attr"`
	Metadata         Metadata   `xml:"metadata" json:"metadata"`
	Manifest         Manifest   `xml:"manifest"`
	Spine            Spine      `xml:"spine"`
	Guide            Guide      `xml:"Guide"`
	Bindings         Bindings   `xml:"bindings"`
	Collection       Collection `xml:"collection"`
}

type Metadata struct {
	XMLName     xml.Name      `xml:"metadata"`
	Title       string        `xml:"dc title" json:"title"`
	Author      []string      `xml:"dc creator" json:"author"`
	Contributor []string      `xml:"dc contributor"`
	Coverage    *string       `xml:"dc coverage"`
	ID          string        `xml:"dc identifier"`
	Language    string        `xml:"dc language"`
	Date        *string       `xml:"dc date"`
	Description *string       `xml:"dc description"`
	Format      *string       `xml:"dc format"`
	Publisher   *string       `xml:"dc publisher"`
	Relation    *string       `xml:"dc relation"`
	Rights      *string       `xml:"dc rights"`
	Source      *string       `xml:"dc source"`
	Subject     *string       `xml:"dc subject"`
	Type        *string       `xml:"dc type"`
	Meta        []Meta        `xml:"meta"`
	Links       []ContentLink `xml:"link"`
}

type Meta struct {
	XMLName  xml.Name `xml:"meta"`
	Property string   `xml:"property,attr"`
	Refines  *string  `xml:"refines,attr"`
	Name     *string  `xml:"name,attr"`
	Content  *string  `xml:"content,attr"`
	Scheme   *string  `xml:"scheme,attr"`
}

type Manifest struct {
	XMLName xml.Name `xml:"manifest"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	XMLName   xml.Name `xml:"item"`
	ID        string   `xml:"id,attr"`
	HREF      string   `xml:"href,attr"`
	MediaType string   `xml:"media-type,attr"`
	Fallback  string   `xml:"fallback,attr"`
}

type Spine struct {
}

type Guide struct{}

type Bindings struct{}

type Collection struct{}

type ContentLink struct {
	XMLName    xml.Name `xml:"link"`
	HREF       string   `xml:"href,attr"`
	Rel        string   `xml:"rel,attr"`
	HREFLang   *string  `xml:"hreflang,attr"`
	ID         *string  `xml:"id,attr"`
	MediaType  *string  `xml:"media-type,attr"`
	Properties *string  `xml:"properties,attr"`
	Refines    *string  `xml:"refines,attr"`
}
