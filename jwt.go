package jwt

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Token represents a JWT token.
// See: https://tools.ietf.org/html/rfc7519
//
type Token struct {
	raw       []byte
	dot1      int
	dot2      int
	header    Header
	claims    []byte
	signature []byte
}

func (t *Token) String() string {
	return string(t.raw)
}

// Bytes representation of the token.
func (t *Token) Bytes() []byte {
	return t.raw
}

// DecodeClaims into a given container.
func (t *Token) DecodeClaims(into interface{}) error {
	err := json.Unmarshal(t.claims, into)
	if err != nil {
		return fmt.Errorf("jwt: claims are not valid: %w", err)
	}
	return nil
}

// Header of the token.
func (t *Token) Header() Header {
	return t.header
}

// Claims of the token.
func (t *Token) Claims() []byte {
	return t.claims
}

// Signature of the token.
func (t *Token) Signature() []byte {
	return t.signature
}

// HeaderPart of the token (base64 encoded).
func (t *Token) HeaderPart() []byte {
	return t.raw[:t.dot1]
}

// ClaimsPart of the token (base64 encoded).
func (t *Token) ClaimsPart() []byte {
	return t.raw[t.dot1+1 : t.dot2]
}

// PayloadPart returns token's payload (header and claims in base64).
func (t *Token) PayloadPart() []byte {
	return t.raw[:t.dot2]
}

// SignaturePart of the token (base64 encoded).
func (t *Token) SignaturePart() []byte {
	return t.raw[t.dot2+1:]
}

// Header is a JWT header.
// See: https://tools.ietf.org/html/rfc7519#section-5
//
type Header struct {
	Algorithm   Algorithm `json:"alg"`
	Type        string    `json:"typ,omitempty"` // only "JWT" can be here
	ContentType string    `json:"cty,omitempty"`
}

func (h *Header) String() string {
	b, _ := h.MarshalJSON()
	return string(b)
}

// MarshalJSON implements the json.Marshaler interface.
func (h *Header) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString(`{"alg":"`)
	buf.WriteString(string(h.Algorithm))

	if h.Type != "" {
		buf.WriteString(`","typ":"`)
		buf.WriteString(h.Type)
	}
	if h.ContentType != "" {
		buf.WriteString(`","cty":"`)
		buf.WriteString(h.ContentType)
	}
	buf.WriteString(`"}`)

	return buf.Bytes(), nil
}
