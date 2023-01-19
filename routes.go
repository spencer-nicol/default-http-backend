package main

func (s *server) routes() {
	s.router.HandleFunc("/", handleDefault())
	s.router.HandleFunc("/healthz", handleHealthCheck())
}
