package xmlHelpers

import (
	"encoding/xml"
	"testing"
)

func TestGetAttr(t *testing.T) {
	cases := []struct {
		element       xml.StartElement
		query         string
		expectedValue string
		expectedFound bool
	}{
		{
			xml.StartElement{
				Name: xml.Name{Local: "element", Space: ""},
				Attr: []xml.Attr{
					{
						Name:  xml.Name{Local: "attribute1", Space: ""},
						Value: "value1",
					},
					{
						Name:  xml.Name{Local: "attribute2", Space: ""},
						Value: "value2",
					},
				},
			},
			"attribute3",
			"",
			false,
		},
		{
			xml.StartElement{
				Name: xml.Name{Local: "element", Space: ""},
				Attr: []xml.Attr{
					{
						Name:  xml.Name{Local: "attribute3", Space: ""},
						Value: "value3",
					},
					{
						Name:  xml.Name{Local: "attribute4", Space: ""},
						Value: "value4",
					},
				},
			},
			"attribute",
			"value4",
			true,
		},
	}

	for _, c := range cases {
		resultValue, resultFound := GetAttr(&c.element, c.query)

		if resultValue != c.expectedValue {
			t.Errorf("Expected value %s to equal %s", resultValue, c.expectedValue)
		}

		if resultFound != c.expectedFound {
			t.Errorf("Expected value %t to equal %t", resultFound, c.expectedFound)
		}
	}
}
