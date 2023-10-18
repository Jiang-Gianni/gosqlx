package rdms

import (
	"time"
)

type QueryResults struct {
	Columns  []string
	Rows     [][]string
	ExecTime time.Duration
	Query    string
}

func (d *Database) ExecQuery(query string) (*QueryResults, error) {
	qr := &QueryResults{
		Query: query,
	}

	defer func(start time.Time) {
		qr.ExecTime = time.Duration(time.Since(start))
	}(time.Now())

	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	qr.Columns = cols

	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols))

	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		result := make([]string, len(cols))
		if err := rows.Scan(dest...); err != nil {
			return nil, err
		}
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = ""
			} else {
				result[i] = string(raw)
			}
		}
		qr.Rows = append(qr.Rows, result)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return qr, nil
}
