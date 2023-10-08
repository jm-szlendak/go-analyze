package analyser

import (
	"encoding/xml"
	"errors"
)

func Analyse(decoder *xml.Decoder) (interface{}, error) {

	for {
		token, err := decoder.Token()

		if err != nil {
			return nil, err
		}

		element, isStartElement := token.(xml.StartElement)

		if !isStartElement {
			continue
		}

		switch element.Name.Local {
		case "posts":
			return AnalysePosts(decoder)
		}

		break
	}

	return nil, errors.New("cannot determine error type")
}
