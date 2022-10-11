package main

import "encoding/xml"

type EPUBContainer struct {
	XMLName   xml.Name       `xml:"container"`
	RootFiles []EPUBRootFile `xml:"rootfiles>rootfile" json:"rootfiles"`
	XMLNS     string         `xml:"xmlns,attr"`
	Version   string         `xml:"version,attr"`
}

type EPUBRootFile struct {
	XMLName xml.Name `xml:"rootfile"`
	Path    string   `xml:"full-path,attr"`
}

type EPUBPackage struct {
	XMLName  xml.Name     `xml:"package"`
	Metadata EPUBMetadata `xml:"metadata" json:"metadata"`
}

type EPUBMetadata struct {
	XMLName xml.Name `xml:"metadata"`
	Title   string   `xml:"title" json:"title"`
	Author  string   `xml:"creator" json:"author"`
}
