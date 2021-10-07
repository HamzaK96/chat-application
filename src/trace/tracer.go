package trace

import (
	"fmt"
	"io"
)

type nilTracer struct {
}

func (t *nilTracer) Trace(a ...interface{}) {

}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

type Tracer interface {
	Trace(...interface{})
}

func Off() Tracer {
	return &nilTracer{}
}

func New(w io.Writer) Tracer { return &tracer{out: w} }
