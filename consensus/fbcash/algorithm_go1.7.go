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

// +build !go1.8

package fbcash

// cacheSize calculates and returns the size of the fbcash verification cache that
// belongs to a certain block number. The cache size grows linearly, however, we
// always take the highest prime below the linearly growing threshold in order to
// reduce the risk of accidental regularities leading to cyclic behavior.
func cacheSize(block uint64) uint64 {
	// If we have a pre-generated value, use that
	epoch := int(block / epochLength)
	if epoch < len(cacheSizes) {
		return cacheSizes[epoch]
	}
	// We don't have a way to verify primes fast before Go 1.8
	panic("fast prime testing unsupported in Go < 1.8")
}

// datasetSize calculates and returns the size of the fbcash mining dataset that
// belongs to a certain block number. The dataset size grows linearly, however, we
// always take the highest prime below the linearly growing threshold in order to
// reduce the risk of accidental regularities leading to cyclic behavior.
func datasetSize(block uint64) uint64 {
	// If we have a pre-generated value, use that
	epoch := int(block / epochLength)
	if epoch < len(datasetSizes) {
		return datasetSizes[epoch]
	}
	// We don't have a way to verify primes fast before Go 1.8
	panic("fast prime testing unsupported in Go < 1.8")
}
