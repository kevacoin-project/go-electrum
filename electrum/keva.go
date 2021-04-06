package electrum

type GetKeyValueResp struct {
	Result GetKeyValueResult `json:"result"`
}

type KeyValue struct {
	Hash   string `json:"tx_hash"`
	Height int32  `json:"height"`
	Key    string `json:"key"`
	Value  string `json:"value,omitempty"`
	Type   string `json:"type"`
	Time   uint32 `json:"time"`
}

type GetKeyValueResult struct {
	KeyValueList []*KeyValue `json:"keyvalues"`
	MinTxNum     int32       `json:"min_tx_num"`
}

func (s *Server) GetKeyValues(namespaceScripthash string, minTxNum int32) (GetKeyValueResult, error) {
	var resp GetKeyValueResp

	err := s.request("blockchain.keva.get_keyvalues", []interface{}{namespaceScripthash, minTxNum}, &resp)
	if err != nil {
		return GetKeyValueResult{}, err
	}

	return resp.Result, err
}
