// Copyright 2017 The go-fairblock Authors
// This file is part of the go-fairblock library.
//
// The go-fairblock library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-fairblock library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-fairblock library. If not, see <http://www.gnu.org/licenses/>.

package fbcash

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"runtime"
	"sync"

	"github.com/fairblock/go-fairblock/common"
	"github.com/fairblock/go-fairblock/consensus"
	"github.com/fairblock/go-fairblock/core/types"
	"github.com/fairblock/go-fairblock/log"
)

// Seal implements consensus.Engine, attempting to find a nonce that satisfies
// the block's difficulty requirements.
func (fbcash *Fbcash) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	// If we're running a fake PoW, simply return a 0 nonce immediately
	if fbcash.fakeMode {
		header := block.Header()
		header.Nonce, header.MixDigest = types.BlockNonce{}, common.Hash{}
		return block.WithSeal(header), nil
	}
	// If we're running a shared PoW, delegate sealing to it
	if fbcash.shared != nil {
		return fbcash.shared.Seal(chain, block, stop)
	}
	// Create a runner and the multiple search threads it directs
	abort := make(chan struct{})
	found := make(chan *types.Block)

	fbcash.lock.Lock()
	threads := fbcash.threads
	if fbcash.rand == nil {
		seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			fbcash.lock.Unlock()
			return nil, err
		}
		fbcash.rand = rand.New(rand.NewSource(seed.Int64()))
	}
	fbcash.lock.Unlock()
	if threads == 0 {
		threads = runtime.NumCPU()
	}
	if threads < 0 {
		threads = 0 // Allows disabling local mining without extra logic around local/remote
	}
	var pend sync.WaitGroup
	for i := 0; i < threads; i++ {
		pend.Add(1)
		go func(id int, nonce uint64) {
			defer pend.Done()
			fbcash.mine(block, id, nonce, abort, found)
		}(i, uint64(fbcash.rand.Int63()))
	}
	// Wait until sealing is terminated or a nonce is found
	var result *types.Block
	select {
	case <-stop:
		// Outside abort, stop all miner threads
		close(abort)
	case result = <-found:
		// One of the threads found a block, abort all others
		close(abort)
	case <-fbcash.update:
		// Thread count was changed on user request, restart
		close(abort)
		pend.Wait()
		return fbcash.Seal(chain, block, stop)
	}
	// Wait for all miners to terminate and return the block
	pend.Wait()
	return result, nil
}

// mine is the actual proof-of-work miner that searches for a nonce starting from
// seed that results in correct final block difficulty.
func (fbcash *Fbcash) mine(block *types.Block, id int, seed uint64, abort chan struct{}, found chan *types.Block) {
	// Extract some data from the header
	var (
		header = block.Header()
		hash   = header.HashNoNonce().Bytes()
		target = new(big.Int).Div(maxUint256, header.Difficulty)

		number  = header.Number.Uint64()
		dataset = fbcash.dataset(number)
	)
	// Start generating random nonces until we abort or find a good one
	var (
		attempts = int64(0)
		nonce    = seed
	)
	logger := log.New("miner", id)
	logger.Trace("Started fbcash search for new nonces", "seed", seed)
	for {
		select {
		case <-abort:
			// Mining terminated, update stats and abort
			logger.Trace("Fbcash nonce search aborted", "attempts", nonce-seed)
			fbcash.hashrate.Mark(attempts)
			return

		default:
			// We don't have to update hash rate on every nonce, so update after after 2^X nonces
			attempts++
			if (attempts % (1 << 15)) == 0 {
				fbcash.hashrate.Mark(attempts)
				attempts = 0
			}
			// Compute the PoW value of this nonce
			digest, result := hashimotoFull(dataset, hash, nonce)
			if new(big.Int).SetBytes(result).Cmp(target) <= 0 {
				// Correct nonce found, create a new header with it
				header = types.CopyHeader(header)
				header.Nonce = types.EncodeNonce(nonce)
				header.MixDigest = common.BytesToHash(digest)

				// Seal and return a block (if still needed)
				select {
				case found <- block.WithSeal(header):
					logger.Trace("Fbcash nonce found and reported", "attempts", nonce-seed, "nonce", nonce)
				case <-abort:
					logger.Trace("Fbcash nonce found but discarded", "attempts", nonce-seed, "nonce", nonce)
				}
				return
			}
			nonce++
		}
	}
}
