package xmlHelpers

import (
	"encoding/xml"

	"golang.org/x/exp/slices"
)

func GetAttr(elem *xml.StartElement, name string) (string, bool) {
	attrs := elem.Attr
	i := slices.IndexFunc(attrs, func(a xml.Attr) bool {
		return a.Name.Local == name
	})

	if i < 0 {
		return "", false
	}

	return attrs[i].Value, true
}
