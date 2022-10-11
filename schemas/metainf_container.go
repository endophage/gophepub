package schemas

import "encoding/xml"

type Container struct {
	XMLName   xml.Name  `xml:"container"`
	RootFiles RootFiles `xml:"rootfiles" json:"rootfiles"`
	Links     []Link    `xml:"links>link" json:"links"`
	XMLNS     string    `xml:"xmlns,attr"`
	Version   string    `xml:"version,attr"`
}

type RootFiles struct {
	XMLName  xml.Name   `xml:"rootfiles"`
	RootFile []RootFile `xml:"rootfile"`
}

type RootFile struct {
	XMLName   xml.Name `xml:"rootfile"`
	FullPath  string   `xml:"full-path,attr"`
	MediaType *string  `xml:"media-type,attr"`
}

type Link struct {
	XMLName   xml.Name `xml:"link"`
	HREF      string   `xml:"href,attr"`
	MediaType *string  `xml:"media-type,attr"`
	Rel       string   `xml:"rel,attr"`
}
