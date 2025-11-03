package main

import "strconv"

type SlotData struct {
	AttestationsCount          uint64  `json:"attestations_count"`
	AttesterSlashingsCount     uint64  `json:"attester_slashings_count"`
	BlockRoot                  string  `json:"block_root"`
	DepositsCount              uint64  `json:"deposits_count"`
	Epoch                      uint64  `json:"epoch"`
	ExecBaseFeePerGas          uint64  `json:"exec_base_fee_per_gas"`
	ExecBlockHash              string  `json:"exec_block_hash"`
	ExecBlockNumber            uint64  `json:"exec_block_number"`
	ExecExtraData              string  `json:"exec_extra_data"`
	ExecFeeRecipient           string  `json:"exec_fee_recipient"`
	ExecGasLimit               uint64  `json:"exec_gas_limit"`
	ExecGasUsed                uint64  `json:"exec_gas_used"`
	ExecLogsBloom              string  `json:"exec_logs_bloom"`
	ExecParentHash             string  `json:"exec_parent_hash"`
	ExecRandom                 string  `json:"exec_random"`
	ExecReceiptsRoot           string  `json:"exec_receipts_root"`
	ExecStateRoot              string  `json:"exec_state_root"`
	ExecTimestamp              uint64  `json:"exec_timestamp"`
	ExecTransactionsCount      uint64  `json:"exec_transactions_count"`
	Graffiti                   string  `json:"graffiti"`
	GraffitiText               string  `json:"graffiti_text"`
	ParentRoot                 string  `json:"parent_root"`
	Proposer                   uint64  `json:"proposer"`
	ProposerSlashingsCount     uint64  `json:"proposer_slashings_count"`
	Slot                       uint64  `json:"slot"`
	StateRoot                  string  `json:"state_root"`
	Status                     string  `json:"status"`
	SyncAggregateParticipation float64 `json:"sync_aggregate_participation"`
	VoluntaryExitsCount        uint64  `json:"voluntary_exits_count"`
	WithdrawalCount            uint64  `json:"withdrawal_count"`
	BlobCount                  uint64  `json:"blob_count"`
	Eth1dataBlockHash          string  `json:"eth1data_blockhash"`
	Eth1dataDepositCount       uint64  `json:"eth1data_depositcount"`
	Eth1dataDepositRoot        string  `json:"eth1data_depositroot"`
}

func buildSlotDataFromMap(m map[string]interface{}) SlotData {
	var s SlotData
	// note: we only read existing upstream keys; output uses unified snake_case via tags
	s.AttestationsCount = asUint(m["attestationscount"]) // upstream uses compact naming
	s.AttesterSlashingsCount = asUint(m["attesterslashingscount"])
	s.BlockRoot = asString(m["blockroot"])
	s.DepositsCount = asUint(m["depositscount"])
	s.Epoch = asUint(m["epoch"])
	s.ExecBaseFeePerGas = asUint(m["exec_base_fee_per_gas"])
	s.ExecBlockHash = asString(m["exec_block_hash"])
	s.ExecBlockNumber = asUint(m["exec_block_number"])
	s.ExecExtraData = asString(m["exec_extra_data"])
	s.ExecFeeRecipient = asString(m["exec_fee_recipient"])
	s.ExecGasLimit = asUint(m["exec_gas_limit"])
	s.ExecGasUsed = asUint(m["exec_gas_used"])
	s.ExecLogsBloom = asString(m["exec_logs_bloom"])
	s.ExecParentHash = asString(m["exec_parent_hash"])
	s.ExecRandom = asString(m["exec_random"])
	s.ExecReceiptsRoot = asString(m["exec_receipts_root"])
	s.ExecStateRoot = asString(m["exec_state_root"])
	s.ExecTimestamp = asUint(m["exec_timestamp"])
	s.ExecTransactionsCount = asUint(m["exec_transactions_count"])
	s.Graffiti = asString(m["graffiti"])
	s.GraffitiText = asString(m["graffiti_text"])
	s.ParentRoot = asString(m["parentroot"])
	s.Proposer = asUint(m["proposer"])
	s.ProposerSlashingsCount = asUint(m["proposerslashingscount"])
	s.Slot = asUint(m["slot"])
	s.StateRoot = asString(m["stateroot"])
	s.Status = asString(m["status"])
	s.SyncAggregateParticipation = asFloat(m["syncaggregate_participation"])
	s.VoluntaryExitsCount = asUint(m["voluntaryexitscount"])
	s.WithdrawalCount = asUint(m["withdrawalcount"])
	s.BlobCount = asUint(m["blob_count"])
	s.Eth1dataBlockHash = asString(m["eth1data_blockhash"])
	s.Eth1dataDepositCount = asUint(m["eth1data_depositcount"])
	s.Eth1dataDepositRoot = asString(m["eth1data_depositroot"])
	return s
}

func asUint(v interface{}) uint64 {
	switch t := v.(type) {
	case float64:
		return uint64(t)
	case string:
		if t == "" {
			return 0
		}
		n, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			return 0
		}
		return n
	default:
		return 0
	}
}

func asFloat(v interface{}) float64 {
	switch t := v.(type) {
	case float64:
		return t
	case string:
		if t == "" {
			return 0
		}
		f, err := strconv.ParseFloat(t, 64)
		if err != nil {
			return 0
		}
		return f
	default:
		return 0
	}
}

func asString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
