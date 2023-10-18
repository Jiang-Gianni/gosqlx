{% import "github.com/Jiang-Gianni/gosqlx/rdms" %}
{% import "strings" %}

{% func ExportSql(qr *rdms.QueryResults, insertIntoTable string) %}INSERT INTO {%s insertIntoTable %}({%s strings.Join(qr.Columns, ", ") %}) VALUES {% for i, row := range qr.Rows %}
('{%s= strings.Join(row, "', '") %}'){% if len(qr.Rows) > 0 && i != len(qr.Rows) - 1 %},{% else %};{% endif %}{% endfor %}
{% endfunc %}