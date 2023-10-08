package analyser

import (
	"encoding/xml"
	"io"

	xmlHelpers "github.com/jm-szlendak/go-analyze/helpers"
)

type PostAnalysisResults struct {
	TotalPosts        int `json:"total_posts"`
	Questions         int `json:"questions"`
	Answers           int `json:"answers"`
	AnsweredQuestions int `json:"answered_questions"`
}

func AnalysePosts(xmlDecoder *xml.Decoder) (PostAnalysisResults, error) {
	results := PostAnalysisResults{}

	for {
		token, err := xmlDecoder.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			return PostAnalysisResults{}, err
		}

		elem, is_start_element := token.(xml.StartElement)

		if !is_start_element {
			continue
		}

		if elem.Name.Local != "row" {
			continue
		}

		postType, found := xmlHelpers.GetAttr(&elem, "PostTypeId")
		if found {
			switch postType {
			case "1":
				results.Questions++
				_, hasAnswer := xmlHelpers.GetAttr(&elem, "AcceptedAnswerId")
				if hasAnswer {
					results.AnsweredQuestions++
				}
			case "2":
				results.Answers++
			}
		}

		results.TotalPosts++

	}

	return results, nil
}
