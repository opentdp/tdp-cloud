package args

import (
	"github.com/open-tdp/go-helper/dborm"
)

var Debug bool

var Database = dborm.Config{}

var Dataset struct {
	Dir    string
	Secret string
}

var Logger struct {
	Dir    string
	Level  string
	Target string
}

var Server struct {
	Listen string
	JwtKey string
}

var Worker struct {
	Remote string
}
