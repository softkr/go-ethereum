// Copyright 2018 The aqchain-blockchain Authors
// This file is part of the aqchain-blockchain library.
//
// The aqchain-blockchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The aqchain-blockchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the aqchain-blockchain library. If not, see <http://www.gnu.org/licenses/>.

// Package poc implements the delegated-proof-of-stake consensus engine.

package poc

import (
	"gitee.com/aqchain/go-ethereum/common"
	"gitee.com/aqchain/go-ethereum/consensus"
	"gitee.com/aqchain/go-ethereum/core/types"
	"gitee.com/aqchain/go-ethereum/rpc"
)

// API is a user facing RPC API to allow controlling the signer and voting
// mechanisms of the delegated-proof-of-stake scheme.
type API struct {
	chain consensus.ChainReader
	poc   *PoC
}

// GetSnapshot retrieves the state snapshot at a given block.
func (api *API) GetSnapshot(number *rpc.BlockNumber) (*Snapshot, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return its snapshot
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.poc.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil,  defaultLoopCntRecalculateSigners)

}

// GetSnapshotAtHash retrieves the state snapshot at a given block.
func (api *API) GetSnapshotAtHash(hash common.Hash) (*Snapshot, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.poc.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil,  defaultLoopCntRecalculateSigners)
}

// GetSnapshotAtNumber retrieves the state snapshot at a given block.
func (api *API) GetSnapshotAtNumber(number uint64) (*Snapshot, error) {
	header := api.chain.GetHeaderByNumber(number)
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.poc.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil,  defaultLoopCntRecalculateSigners)
}
