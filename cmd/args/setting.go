package args

type Server struct {
	Listen string
	Dsn    string
}

type Worker struct {
	Remote string
}

type Setting struct {
	Server Server
	Worker Worker
}
