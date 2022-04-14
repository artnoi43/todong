package ginserver

func (g *ginServer) Serve(addr string) error {
	return g.engine.Run(addr)
}
