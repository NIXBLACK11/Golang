package api

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserById)
	return http.ListenAndServe(s.listenAddr, nil)
}

