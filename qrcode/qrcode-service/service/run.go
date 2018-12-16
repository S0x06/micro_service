package service

/*
type Config struct {
	HTTPAddr  string
	DebugAddr string
	GRPCAddr  string
}

func NewEndpoints() request.Endpoints {

}
*/
func Run(cfg Config) {

	endpoints := NewEndpoints()

	// Mechanical domain.
	errc := make(chan error)

	// Interrupt handler.
	go handlers.InterruptHandler(errc)


	// Debug listener.
	go func() {
		log.Println("transport", "debug", "addr", cfg.DebugAddr)

		m := http.NewServeMux()
		m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

		errc <- http.ListenAndServe(cfg.DebugAddr, m)
	}()


	// Run!
	// 

	
	log.Println("exit", <-errc)


}
