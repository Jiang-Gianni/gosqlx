package rdms

type QueryResults struct {
	Columns []string
	Rows    [][]string
}

func (d *Database) ExecQuery(query string) (*QueryResults, error) {
	qr := &QueryResults{}
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	qr.Columns = cols

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))
	// A temporary interface{} slice
	dest := make([]interface{}, len(cols))

	// Put pointers to each string in the interface slice
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
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
