package schemas

import (
	"encoding/xml"
	"github.com/stretchr/testify/require"
	"testing"
)

const containerSample = `<?xml version="1.0"?>
<container version="1.0" xmlns="urn:oasis:names:tc:opendocument:xmlns:container">
   <rootfiles>
      <rootfile
          full-path="EPUB/My_Crazy_Life.opf"
          media-type="application/oebps-package+xml" />
   </rootfiles>
</container>`

func TestParseContainerXML(t *testing.T) {
	parsed := Container{}
	require.NoError(t, xml.Unmarshal([]byte(containerSample), &parsed))

	require.Len(t, parsed.RootFiles.RootFile, 1)
	rootfile := parsed.RootFiles.RootFile[0]
	require.Equal(
		t,
		"application/oebps-package+xml",
		*rootfile.MediaType,
	)
	require.Equal(
		t,
		"EPUB/My_Crazy_Life.opf",
		rootfile.FullPath,
	)
}
