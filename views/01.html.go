// Code generated by qtc from "01.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/01.html:1
package views

//line views/01.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/01.html:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/01.html:1
func StreamFirst(qw422016 *qt422016.Writer) {
//line views/01.html:1
	qw422016.N().S(`
<div>This is first</div>
`)
//line views/01.html:3
}

//line views/01.html:3
func WriteFirst(qq422016 qtio422016.Writer) {
//line views/01.html:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/01.html:3
	StreamFirst(qw422016)
//line views/01.html:3
	qt422016.ReleaseWriter(qw422016)
//line views/01.html:3
}

//line views/01.html:3
func First() string {
//line views/01.html:3
	qb422016 := qt422016.AcquireByteBuffer()
//line views/01.html:3
	WriteFirst(qb422016)
//line views/01.html:3
	qs422016 := string(qb422016.B)
//line views/01.html:3
	qt422016.ReleaseByteBuffer(qb422016)
//line views/01.html:3
	return qs422016
//line views/01.html:3
}
