package export

import (
	"fmt"
	"os"

	"github.com/Jiang-Gianni/gosqlx/rdms"
)

func ExportFile(qr *rdms.QueryResults, insertIntoTable string, outputFile string, extension string) error {
	switch extension {
	case "sql":
		return os.WriteFile(outputFile, []byte(ExportSql(qr, insertIntoTable)), os.ModePerm)
	case "csv":
		return os.WriteFile(outputFile, []byte(ExportCsv(qr)), os.ModePerm)
	default:
		return fmt.Errorf("invalid extension")
	}
}
