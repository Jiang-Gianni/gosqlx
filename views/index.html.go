// Code generated by qtc from "index.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/index.html:1
package views

//line views/index.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/index.html:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/index.html:2
type Page struct {
	Title       string
	Description string
}

//line views/index.html:8
func StreamCommonHeaders(qw422016 *qt422016.Writer) {
//line views/index.html:8
	qw422016.N().S(`
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="./assets/pico.css">
<!-- <link rel="stylesheet" href="./assets/water.css"> -->
<script src="./assets/htmx.js" async defer></script>

<!-- TODO remove before pushing to production -->
<!-- <script>
    var wsuri = "ws://localhost:1234/ws";
    window.onload = function () {
        sock = new WebSocket(wsuri);
        sock.onopen = function () {
            console.log("connected to " + wsuri);
        }
        sock.onclose = function (e) {
            console.log("connection closed (" + e.code + ")");
        }
        sock.onmessage = function (e) {
            window.location.href = window.location.href
        }
    };
</script> -->

`)
//line views/index.html:32
}

//line views/index.html:32
func WriteCommonHeaders(qq422016 qtio422016.Writer) {
//line views/index.html:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/index.html:32
	StreamCommonHeaders(qw422016)
//line views/index.html:32
	qt422016.ReleaseWriter(qw422016)
//line views/index.html:32
}

//line views/index.html:32
func CommonHeaders() string {
//line views/index.html:32
	qb422016 := qt422016.AcquireByteBuffer()
//line views/index.html:32
	WriteCommonHeaders(qb422016)
//line views/index.html:32
	qs422016 := string(qb422016.B)
//line views/index.html:32
	qt422016.ReleaseByteBuffer(qb422016)
//line views/index.html:32
	return qs422016
//line views/index.html:32
}

//line views/index.html:34
func StreamIndexPage(qw422016 *qt422016.Writer, p *Page, body string) {
//line views/index.html:34
	qw422016.N().S(`
<!DOCTYPE html>
<html lang="en" data-theme="dark">

<head>
    <title>`)
//line views/index.html:39
	qw422016.E().S(p.Title)
//line views/index.html:39
	qw422016.N().S(`</title>
    <meta name="description" content="`)
//line views/index.html:40
	qw422016.E().S(p.Description)
//line views/index.html:40
	qw422016.N().S(`">
    `)
//line views/index.html:41
	qw422016.N().S(CommonHeaders())
//line views/index.html:41
	qw422016.N().S(`

</head>

<body>
    `)
//line views/index.html:46
	qw422016.N().S(body)
//line views/index.html:46
	qw422016.N().S(`
</body>

</html>
`)
//line views/index.html:50
}

//line views/index.html:50
func WriteIndexPage(qq422016 qtio422016.Writer, p *Page, body string) {
//line views/index.html:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/index.html:50
	StreamIndexPage(qw422016, p, body)
//line views/index.html:50
	qt422016.ReleaseWriter(qw422016)
//line views/index.html:50
}

//line views/index.html:50
func IndexPage(p *Page, body string) string {
//line views/index.html:50
	qb422016 := qt422016.AcquireByteBuffer()
//line views/index.html:50
	WriteIndexPage(qb422016, p, body)
//line views/index.html:50
	qs422016 := string(qb422016.B)
//line views/index.html:50
	qt422016.ReleaseByteBuffer(qb422016)
//line views/index.html:50
	return qs422016
//line views/index.html:50
}

//line views/index.html:54
func StreamIndexBody(qw422016 *qt422016.Writer) {
//line views/index.html:54
	qw422016.N().S(`
<h1> Title </h1>
<button hx-get="first">
    <a href="first"></a>
    First
</button>
<code>Let's go there</code>
`)
//line views/index.html:61
}

//line views/index.html:61
func WriteIndexBody(qq422016 qtio422016.Writer) {
//line views/index.html:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/index.html:61
	StreamIndexBody(qw422016)
//line views/index.html:61
	qt422016.ReleaseWriter(qw422016)
//line views/index.html:61
}

//line views/index.html:61
func IndexBody() string {
//line views/index.html:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/index.html:61
	WriteIndexBody(qb422016)
//line views/index.html:61
	qs422016 := string(qb422016.B)
//line views/index.html:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/index.html:61
	return qs422016
//line views/index.html:61
}
