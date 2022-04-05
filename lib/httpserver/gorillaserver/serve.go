package gorillaserver

func (s *gorillaServer) Serve(addr string) error {
	return s.server.ListenAndServe()
}
