package tiktokads

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"strings"
)

func convertCsv[T any](source io.Reader) ([]*T, error) {
	r := csv.NewReader(source)
	r.Comma = ','
	r.LazyQuotes = true
	r.TrimLeadingSpace = true

	headers := []string{}

	results := []*T{}

	for {
		record, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}

		if len(headers) == 0 {
			for _, h := range record {
				headers = append(headers, slugify(h))
			}
		} else {
			if len(record) != len(headers) {
				return nil, errors.New("wrong number of headers")
			}
			data := map[string]string{}
			for i, h := range record {
				data[headers[i]] = h
			}
			buf, _ := json.Marshal(data)
			var row T
			if err = json.Unmarshal(buf, &row); nil != err {
				return nil, err
			} else {
				results = append(results, &row)
			}
		}
	}

	return results, nil
}

func slugify(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))
	s = strings.ReplaceAll(s, "(", "_")
	s = strings.ReplaceAll(s, ")", "_")
	s = strings.ReplaceAll(s, " ", "_")

	return s
}
