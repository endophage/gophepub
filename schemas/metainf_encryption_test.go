package schemas

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const encryptionSample = `<encryption
    xmlns ="urn:oasis:names:tc:opendocument:xmlns:container"
    xmlns:enc="http://www.w3.org/2001/04/xmlenc#"
    xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
   <enc:EncryptedKey Id="EK">
      <enc:EncryptionMethod
          Algorithm="http://www.w3.org/2001/04/xmlenc#rsa-1_5"/>
      <ds:KeyInfo>
         <ds:KeyName>
            John Smith
         </ds:KeyName>
      </ds:KeyInfo>
      <enc:CipherData>
         <enc:CipherValue>
            xyzabc
         </enc:CipherValue>
      </enc:CipherData>
   </enc:EncryptedKey>
   <enc:EncryptedData Id="ED1">
      <enc:EncryptionMethod
          Algorithm="http://www.w3.org/2001/04/xmlenc#kw-aes128"/>
      <ds:KeyInfo>
         <ds:RetrievalMethod
             URI="#EK"
             Type="http://www.w3.org/2001/04/xmlenc#EncryptedKey"/>
      </ds:KeyInfo>
      <enc:CipherData>
         <enc:CipherReference
             URI="image.jpeg"/>
      </enc:CipherData>
   </enc:EncryptedData>
</encryption>`

func TestParseEncryptionXML(t *testing.T) {
	parsed := Encryption{}
	require.NoError(t, xml.Unmarshal([]byte(encryptionSample), &parsed))

	require.Equal(
		t,
		"urn:oasis:names:tc:opendocument:xmlns:container",
		parsed.XMLNS,
	)
	require.Equal(
		t,
		"http://www.w3.org/2001/04/xmlenc#",
		parsed.XMLNSEnc,
	)
	require.Equal(
		t,
		"http://www.w3.org/2000/09/xmldsig#",
		parsed.XMLNSDs,
	)

	// Check we have one key and one data segment
	require.Len(t, parsed.EncryptedKeys, 1)
	require.Len(t, parsed.EncryptedData, 1)
	key := parsed.EncryptedKeys[0]
	data := parsed.EncryptedData[0]

	// Check key segment
	require.Equal(t, "EK", strings.TrimSpace(key.ID))
	require.Equal(
		t,
		"http://www.w3.org/2001/04/xmlenc#rsa-1_5",
		strings.TrimSpace(key.EncryptionMethod.Algorithm),
	)
	require.Equal(
		t,
		"John Smith",
		strings.TrimSpace(key.KeyInfo.KeyName.Value),
	)
	require.Equal(
		t,
		"xyzabc",
		strings.TrimSpace(key.CipherData.CipherValue.Value),
	)

	// Check data segment
	require.Equal(t, "ED1", strings.TrimSpace(data.ID))
	require.Equal(
		t,
		"http://www.w3.org/2001/04/xmlenc#kw-aes128",
		strings.TrimSpace(data.EncryptionMethod.Algorithm),
	)
	require.Equal(
		t,
		"http://www.w3.org/2001/04/xmlenc#EncryptedKey",
		strings.TrimSpace(data.KeyInfo.RetrievalMethod.Type),
	)
	require.Equal(
		t,
		"#EK",
		strings.TrimSpace(data.KeyInfo.RetrievalMethod.URI),
	)
	require.Equal(
		t,
		"image.jpeg",
		strings.TrimSpace(data.CipherData.CipherReference.URI),
	)
}
