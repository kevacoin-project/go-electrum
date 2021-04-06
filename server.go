package electrum

// Ping send a ping to the target server to ensure it is responding and
// keeping the session alive.
// https://electrumx.readthedocs.io/en/latest/protocol-methods.html#server-ping
func (s *Server) Ping() error {
	err := s.request("server.ping", []interface{}{}, nil)

	return err
}

// ServerAddPeer adds your new server into the remote server own peers list.
// This should not be used if you are a client.
// https://electrumx.readthedocs.io/en/latest/protocol-methods.html#server-add-peer
func (s *Server) ServerAddPeer(features *ServerFeaturesResult) error {
	var resp basicResp

	err := s.request("server.add_peer", []interface{}{features}, &resp)

	return err
}

// ServerBanner returns the banner for this remote server.
// https://electrumx.readthedocs.io/en/latest/protocol-methods.html#server-banner
func (s *Server) ServerBanner() (string, error) {
	var resp basicResp

	err := s.request("server.banner", []interface{}{}, &resp)

	return resp.Result, err
}

// ServerDonation returns the donation address for this remote server
// https://electrumx.readthedocs.io/en/latest/protocol-methods.html#server-donation-address
func (s *Server) ServerDonation() (string, error) {
	var resp basicResp

	err := s.request("server.donation_address", []interface{}{}, &resp)

	return resp.Result, err
}

type host struct {
	TCPPort uint16 `json:"tcp_port,omitempty"`
	SSLPort uint16 `json:"ssl_port,omitempty"`
}

// ServerFeaturesResp represent the response to GetFeatures().
type ServerFeaturesResp struct {
	Result *ServerFeaturesResult `json:"result"`
}

// ServerFeaturesResult represent the data sent or receive in RPC call "server.features" and
// "server.add_peer".
type ServerFeaturesResult struct {
	GenesisHash   string          `json:"genesis_hash"`
	Hosts         map[string]host `json:"hosts"`
	ProtocolMax   string          `json:"protocol_max"`
	ProtocolMin   string          `json:"protocol_min"`
	Pruning       bool            `json:"pruning,omitempty"`
	ServerVersion string          `json:"server_version"`
	HashFunction  string          `json:"hash_function"`
}

// ServerFeatures returns a list of features and services supported by the remote server.
// https://electrumx.readthedocs.io/en/latest/protocol-methods.html#server-features
func (s *Server) ServerFeatures() (*ServerFeaturesResult, error) {
	var resp ServerFeaturesResp

	err := s.request("server.features", []interface{}{}, &resp)

	return resp.Result, err
}

// ServerPeers returns a list of peers this remote server is aware of.
// https://electrumx.readthedocs.io/en/latest/protocol-methods.html#server-peers-subscribe
func (s *Server) ServerPeers() (interface{}, error) {
	resp := &struct {
		Result [][]interface{} `json:"result"`
	}{}
	err := s.request("server.peers.subscribe", []interface{}{}, &resp)

	return resp.Result, err
}

// ServerVersionResp represent the response to ServerVersion().
type ServerVersionResp struct {
	Result [2]string `json:"result"`
}

// ServerVersion identify the client to the server, and negotiate the protocol version.
// This call must be sent first, or the server will default to an older protocol version.
// https://electrumx.readthedocs.io/en/latest/protocol-methods.html#server-version
func (s *Server) ServerVersion() (serverVer, protocolVer string, err error) {
	var resp ServerVersionResp

	err = s.request("server.version", []interface{}{ClientVersion, ProtocolVersion}, &resp)
	if err != nil {
		serverVer = ""
		protocolVer = ""
	} else {
		serverVer = resp.Result[0]
		protocolVer = resp.Result[1]
	}

	return
}
