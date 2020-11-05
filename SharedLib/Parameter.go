package SharedLib

import "flag"

type ParameterSet struct {
	Host *string
	Port *string
}

func ParseParameter() ParameterSet {
	var params ParameterSet
	params.Host = flag.String("h", "127.0.0.1", "Host")
	params.Port = flag.String("p", "8888", "Port")
	return params
}
