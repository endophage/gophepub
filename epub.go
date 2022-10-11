package main

import (
	"archive/zip"
	"encoding/xml"
	"log"
)

func LoadMetadata(absPath string) *EPUBMetadata {
	zr, err := zip.OpenReader(absPath)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer zr.Close()
	for _, f := range zr.File {
		if f.Name == "META-INF/container.xml" {
			log.Println("Found container.xml")
			return loadMetadataFrom(zr, f)
		}
	}
	return nil
}

func loadMetadataFrom(zr *zip.ReadCloser, f *zip.File) *EPUBMetadata {
	r, err := f.Open()
	if err != nil {
		log.Println(err)
		return nil
	}
	defer r.Close()
	decoder := xml.NewDecoder(r)
	c := EPUBContainer{}
	err = decoder.Decode(&c)
	if err != nil {
		log.Println(err)
		return nil
	}

	log.Println("XML Version: %s", c.Version)
	log.Println("XMLNS: %s", c.XMLNS)
	log.Println("Rootfiles:")
	for _, rf := range c.RootFiles {
		log.Println("    %s", rf.Path)
	}

	if len(c.RootFiles) == 0 {
		log.Println("no root files found")
		return nil
	}
	rootFile := c.RootFiles[0]
	for _, f := range zr.File {
		if f.Name == rootFile.Path {
			r, err = f.Open()
			if err != nil {
				log.Println(err)
				return nil
			}
			decoder = xml.NewDecoder(r)
			meta := EPUBPackage{}
			err = decoder.Decode(&meta)
			if err != nil {
				log.Println(err)
				return nil
			}
			return &meta.Metadata
		}
	}

	return nil
}
