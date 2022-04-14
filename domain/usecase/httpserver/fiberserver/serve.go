package fiberserver

func (f *fiberServer) Serve(addr string) error {
	return f.app.Listen(addr)
}
