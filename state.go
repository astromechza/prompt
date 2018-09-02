package main

import (
	"encoding/json"
	"os"
	"path"
	"time"
)

type BeforeState struct {
	Time time.Time `json:"time"`
}

func StatePath(uid string) string {
	return path.Join(os.TempDir(), ".prompt."+uid)
}

func PutState(state BeforeState, uid string) error {
	f, err := os.OpenFile(StatePath(uid), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	if err = enc.Encode(&state); err != nil {
		return err
	}
	return nil
}

func TryPopState(uid string) (BeforeState, error) {
	o := BeforeState{}
	f, err := os.Open(StatePath(uid))
	if err != nil {
		if os.IsNotExist(err) {
			return o, nil
		}
		return o, err
	}
	defer f.Close()
	defer os.Remove(StatePath(uid))
	dec := json.NewDecoder(f)
	if err = dec.Decode(&o); err != nil {
		return o, err
	}
	return o, nil
}
