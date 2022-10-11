package schemas

import "encoding/xml"

type Encryption struct {
	XMLName       xml.Name        `xml:"encryption"`
	XMLNS         string          `xml:"xmlns,attr"`
	XMLNSEnc      string          `xml:"xmlns enc,attr"`
	XMLNSDs       string          `xml:"xmlns ds,attr"`
	EncryptedKeys []EncryptedKey  `xml:"EncryptedKey"`
	EncryptedData []EncryptedData `xml:"EncryptedData"`
}

type EncryptedKey struct {
	XMLName          xml.Name         `xml:"EncryptedKey"`
	ID               string           `xml:"Id,attr"`
	EncryptionMethod EncryptionMethod `xml:"EncryptionMethod"`
	KeyInfo          KeyInfo          `xml:"KeyInfo"`
	CipherData       CipherData       `xml:"CipherData"`
}

type EncryptionMethod struct {
	XMLName   xml.Name `xml:"EncryptionMethod""`
	Algorithm string   `xml:"Algorithm,attr"`
}

type KeyInfo struct {
	XMLName         xml.Name        `xml:"KeyInfo"`
	KeyName         KeyName         `xml:"KeyName"`
	RetrievalMethod RetrievalMethod `xml:"RetrievalMethod"`
}

type KeyName struct {
	XMLName xml.Name `xml:"KeyName"`
	Value   string   `xml:",chardata"`
}

type RetrievalMethod struct {
	XMLName xml.Name `xml:"RetrievalMethod"`
	URI     string   `xml:"URI,attr"`
	Type    string   `xml:"Type,attr"`
}

type CipherData struct {
	XMLName         xml.Name        `xml:"CipherData"`
	CipherValue     CipherValue     `xml:"CipherValue"`
	CipherReference CipherReference `xml:"CipherReference"`
}

type CipherValue struct {
	XMLName xml.Name `xml:"CipherValue"`
	Value   string   `xml:",chardata"`
}

type CipherReference struct {
	XMLName xml.Name `xml:"CipherReference"`
	URI     string   `xml:"URI,attr"`
}

type EncryptedData struct {
	XMLName          xml.Name         `xml:"EncryptedData"`
	ID               string           `xml:"Id,attr"`
	EncryptionMethod EncryptionMethod `xml:"EncryptionMethod"`
	KeyInfo          KeyInfo          `xml:"KeyInfo"`
	CipherData       CipherData       `xml:"CipherData"`
}

// Go has a terrible issue with xml namespaces so we need a completely separate
// set of structs for Marshalling...
type EncryptionMarshaler struct {
	XMLName       xml.Name        `xml:"encryption"`
	XMLNS         string          `xml:"xmlns,attr"`
	XMLNSEnc      string          `xml:"xmlns:enc,attr"`
	XMLNSDs       string          `xml:"xmlns:ds,attr"`
	EncryptedKeys []EncryptedKey  `xml:"EncryptedKey"`
	EncryptedData []EncryptedData `xml:"EncryptedData"`
}
