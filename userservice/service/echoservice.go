package service

import "log"

type Echo int

func (t *Echo) Hi(args map[string]interface{}, reply *map[string]interface{}) error {
	log.Println(args)
	*reply = map[string]interface{}{
		"a" : "aa",
	}
	return nil
}

