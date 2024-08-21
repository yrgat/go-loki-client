package loki

type LokiClient struct {
	Host string
}

func NewLokiClient(host string) *LokiClient {
	return &LokiClient{}
}

func (lc *LokiClient) SetHost(host string) {
	lc.Host = host
}

func (lc *LokiClient) SendStreams() {

}
