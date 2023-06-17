// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package state

import (
	"io"

	"golang.org/x/crypto/blake2b"

	"github.com/iotaledger/hive.go/kvstore/mapdb"
	"github.com/iotaledger/wasp/packages/isc/coreutil"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/buffered"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/trie"
	"github.com/iotaledger/wasp/packages/util/rwutil"
)

type block struct {
	trieRoot             trie.Hash
	mutations            *buffered.Mutations
	previousL1Commitment *L1Commitment
}

var _ Block = &block{}

func BlockFromBytes(data []byte) (Block, error) {
	return rwutil.ReaderFromBytes(data, new(block))
}

func (b *block) Bytes() []byte {
	return rwutil.WriterToBytes(b)
}

func (b *block) essenceBytes() []byte {
	ww := rwutil.NewBytesWriter()
	ww.WriteFromFunc(b.writeEssence)
	return ww.Bytes()
}

func (b *block) Equals(other Block) bool {
	return b.Hash().Equals(other.Hash())
}

func (b *block) Hash() (ret BlockHash) {
	hash := blake2b.Sum256(b.essenceBytes())
	copy(ret[:], hash[:])
	return ret
}

func (b *block) L1Commitment() *L1Commitment {
	return newL1Commitment(b.TrieRoot(), b.Hash())
}

func (b *block) Mutations() *buffered.Mutations {
	return b.mutations
}

func (b *block) MutationsReader() kv.KVStoreReader {
	return buffered.NewBufferedKVStoreForMutations(
		kv.NewHiveKVStoreReader(mapdb.NewMapDB()),
		b.mutations,
	)
}

func (b *block) PreviousL1Commitment() *L1Commitment {
	return b.previousL1Commitment
}

func (b *block) StateIndex() uint32 {
	return codec.MustDecodeUint32(b.MutationsReader().Get(kv.Key(coreutil.StatePrefixBlockIndex)))
}

func (b *block) TrieRoot() trie.Hash {
	return b.trieRoot
}

func (b *block) Read(r io.Reader) error {
	rr := rwutil.NewReader(r)
	rr.ReadN(b.trieRoot[:])
	rr.ReadFromFunc(b.readEssence)
	return rr.Err
}

func (b *block) Write(w io.Writer) error {
	ww := rwutil.NewWriter(w)
	ww.WriteN(b.trieRoot[:])
	ww.WriteFromFunc(b.writeEssence)
	return ww.Err
}

func (b *block) readEssence(r io.Reader) (int, error) {
	rr := rwutil.NewReader(r)
	counter := rwutil.NewReadCounter(rr)
	b.mutations = buffered.NewMutations()
	rr.Read(b.mutations)
	hasPrevL1Commitment := rr.ReadBool()
	if hasPrevL1Commitment {
		b.previousL1Commitment = new(L1Commitment)
		rr.Read(b.previousL1Commitment)
	}
	return counter.Count(), rr.Err
}

func (b *block) writeEssence(w io.Writer) (int, error) {
	ww := rwutil.NewWriter(w)
	ww.WriteN(b.Mutations().Bytes())
	ww.WriteBool(b.PreviousL1Commitment() != nil)
	if b.PreviousL1Commitment() != nil {
		ww.WriteN(b.PreviousL1Commitment().Bytes())
	}
	return len(ww.Bytes()), ww.Err
}
