package loki

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LokiOpt struct {
	Client *LokiClient
}

type Loki struct {
	opt    *LokiOpt
	stream *Stream
}

func NewLoki(opt *LokiOpt, labels Label) *Loki {
	l := &Loki{}

	l.stream = NewStream(labels)
	l.opt = opt

	return l
}

func (l *Loki) AddMessage(message string) {
	l.stream.AddValue(message)
}

func (l *Loki) Fire() error {
	if !l.stream.HasValues() {
		return fmt.Errorf("there is no message in stream")
	}

	streams := NewStreams()
	streams.AddStream(l.stream.Labels, l.stream.Values)

	j, err := json.Marshal(streams)
	if err != nil {
		return err
	}

	_, err = http.Post(l.opt.Client.Host+"/loki/api/v1/push", "application/json", bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	l.stream.ClearValues()

	fmt.Println(string(j))
	return nil

}
