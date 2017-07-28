package udfs

import (
	"errors"
	"strings"
	
	"pfi/sensorbee/sensorbee/core"
	"pfi/sensorbee/sensorbee/data"
)

type Join struct {

}

func (j *Join) Call(ctx *core.Context, args ...data.Value) (data.Value, error) {
	empty := data.String("")
	if len(args) == 1 {
		return empty, nil
	}

	switch args[0].Type() {
	case data.TypeString:  // my_join("a", "b", "c", "sep") form
		var ss []string
		for _, v := range args {
			s, err := data.AsString(v)
			if err != nil {
				return empty, err
			}
			ss = append(ss, s)
		}
		return data.String(strings.Join(ss[:len(ss)-1], ss[len(ss)-1])), nil

		case 
	}

}
