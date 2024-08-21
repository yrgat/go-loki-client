package loki

type Stream struct {
	Labels Label  `json:"stream"`
	Values Values `json:"values"`
}

type Streams struct {
	Streams []Stream `json:"streams"`
}

func NewStream(labels Label) *Stream {
	return &Stream{Labels: labels}
}

func (s *Stream) AddValue(message string) {
	s.Values.AddValueWithCurrentTime(message)
}

func (s *Stream) ClearValues() {
	s.Values = NewValues()
}

func (s *Stream) HasValues() bool {
	return len(s.Values) != 0
}

func NewStreams() *Streams {
	return &Streams{}
}

func (s *Streams) AddStream(labels Label, values Values) {
	s.Streams = append(s.Streams, Stream{Labels: labels, Values: values})
}
