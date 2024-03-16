package main

import "fmt"

type OptFunc func(*Opts)

type Server struct {
	Opts Opts
}

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
	}
}

func withTls(opts *Opts) {
	opts.tls = true
}

func withMaxConn(n int) OptFunc {
	return func(o *Opts) {
		o.maxConn = n
	}
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, optFn := range opts {
		optFn(&o)
	}
	return &Server{
		Opts: o,
	}

}

func main() {
	s := newServer(withTls, withMaxConn(20))
	fmt.Printf("%+v\n", s)
}
