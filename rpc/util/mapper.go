// Copyright (C) 2018  MediBloc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package rpc

import (
	"github.com/medibloc/go-medibloc/consensus/dpos/pb"
	"github.com/medibloc/go-medibloc/core"
	"github.com/medibloc/go-medibloc/rpc/pb"
	"github.com/medibloc/go-medibloc/util/byteutils"
	"github.com/medibloc/go-nebulas/util"
)

func coreAccount2rpcAccount(account *core.Account, address string) *rpcpb.GetAccountResponse {
	if account == nil {
		return &rpcpb.GetAccountResponse{
			Address:       address,
			Balance:       "0",
			Nonce:         0,
			Vesting:       "0",
			Voted:         "",
			Records:       []string{},
			CertsIssued:   []string{},
			CertsReceived: []string{},
			TxsFrom:       []string{},
			TxsTo:         []string{},
		}
	}
	return &rpcpb.GetAccountResponse{
		Address:       address,
		Balance:       account.Balance().String(),
		Nonce:         account.Nonce(),
		Vesting:       account.Vesting().String(),
		Voted:         byteutils.Bytes2Hex(acc.Voted()),
		Records:       nil, // TODO @ggomma
		CertsIssued:   nil, // TODO @ggomma
		CertsReceived: nil, // TODO @ggomma
		TxsFrom:       byteutils.BytesSlice2HexSlice(account.TxsFrom()),
		TxsTo:         byteutils.BytesSlice2HexSlice(account.TxsTo()),
	}
}

func coreBlock2rpcBlock(block *corepb.Block) (*rpcpb.GetBlockResponse, error) {
	var rpcTxs []*rpcpb.GetTransactionResponse
	for _, tx := range block.GetTransactions() {
		rpcTx, err := coreTx2rpcTx(tx, true)
		if err != nil {
			return nil, err
		}
		rpcTxs = append(rpcTxs, rpcTx)
	}

	return &rpcpb.BlockResponse{
		Height:            block.Height,
		Hash:              byteutils.Bytes2Hex(block.Header.Hash),
		ParentHash:        byteutils.Bytes2Hex(block.Header.ParentHash),
		Coinbase:          byteutils.Bytes2Hex(block.Header.Coinbase),
		Reward:            nil, // todo
		Supply:            nil, //todo
		Timestamp:         block.Header.Timestamp,
		ChainId:           block.Header.ChainId,
		Alg:               block.Header.Alg,
		Sign:              byteutils.Bytes2Hex(block.Header.Sign),
		AccsRoot:          byteutils.Bytes2Hex(block.Header.AccsRoot),
		TxsRoot:           byteutils.Bytes2Hex(block.Header.TxsRoot),
		UsageRoot:         byteutils.Bytes2Hex(block.Header.UsageRoot),
		RecordsRoot:       byteutils.Bytes2Hex(block.Header.RecordsRoot),
		CertificationRoot: byteutils.Bytes2Hex(block.Header.DposRoot),
		DposRoot:          nil, //todo
		Transactions:      rpcPbTxs,
	}, nil
}

func dposCandidate2rpcCandidate(candidate *dpospb.Candidate) (*rpcpb.Candidate, error) {
	collateral, err := util.NewUint128FromFixedSizeByteSlice(candidate.Collateral)
	if err != nil {
		return nil, err
	}

	votePower, err := util.NewUint128FromFixedSizeByteSlice(candidate.VotePower)
	if err != nil {
		return nil, err
	}

	return &rpcpb.Candidate{
		Address:   byteutils.Bytes2Hex(candidate.Address),
		Collatral: collatral.String(),
		VotePower: votePower.String(),
	}, nil
}

func coreTx2rpcTx(tx *corepb.Transaction, executed bool) (*rpcpb.GetTransactionResponse, error) {
	value, err := util.NewUint128FromFixedSizeByteSlice(tx.Value)
	if err != nil {
		return nil, err
	}

	return &rpcpb.GetTransactionResponse{
		Hash:      byteutils.Bytes2Hex(tx.Hash),
		From:      byteutils.Bytes2Hex(tx.From),
		To:        byteutils.Bytes2Hex(tx.To),
		Value:     value.String(),
		Timestamp: tx.Timestamp,
		Data: &rpcpb.TransactionData{
			Type:    tx.Data.Type,
			Payload: string(tx.Data.Payload),
		},
		Nonce:     tx.Nonce,
		ChainId:   tx.ChainId,
		Alg:       tx.Alg,
		Sign:      byteutils.Bytes2Hex(tx.Sign),
		PayerSign: byteutils.Bytes2Hex(tx.PayerSign),
		Executed:  executed,
	}, nil
}
