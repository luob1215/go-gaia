// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package fbc

import (
	"math/big"

	"github.com/fairblock/go-fairblock/common"
	"github.com/fairblock/go-fairblock/common/hexutil"
	"github.com/fairblock/go-fairblock/core"
	"github.com/fairblock/go-fairblock/fbc/downloader"
	"github.com/fairblock/go-fairblock/fbc/gasprice"
)

func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               uint64
		SyncMode                downloader.SyncMode
		LightServ               int  `toml:",omitempty"`
		LightPeers              int  `toml:",omitempty"`
		MaxPeers                int  `toml:"-"`
		SkipBcVersionCheck      bool `toml:"-"`
		DatabaseHandles         int  `toml:"-"`
		DatabaseCache           int
		Fairblockbase               common.Address `toml:",omitempty"`
		MinerThreads            int            `toml:",omitempty"`
		ExtraData               hexutil.Bytes  `toml:",omitempty"`
		GasPrice                *big.Int
		FbcashCacheDir          string
		FbcashCachesInMem       int
		FbcashCachesOnDisk      int
		FbcashDatasetDir        string
		FbcashDatasetsInMem     int
		FbcashDatasetsOnDisk    int
		TxPool                  core.TxPoolConfig
		GPO                     gasprice.Config
		EnablePreimageRecording bool
		DocRoot                 string `toml:"-"`
		PowFake                 bool   `toml:"-"`
		PowTest                 bool   `toml:"-"`
		PowShared               bool   `toml:"-"`
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.LightServ = c.LightServ
	enc.LightPeers = c.LightPeers
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.Fairblockbase = c.Fairblockbase
	enc.MinerThreads = c.MinerThreads
	enc.ExtraData = c.ExtraData
	enc.GasPrice = c.GasPrice
	enc.FbcashCacheDir = c.FbcashCacheDir
	enc.FbcashCachesInMem = c.FbcashCachesInMem
	enc.FbcashCachesOnDisk = c.FbcashCachesOnDisk
	enc.FbcashDatasetDir = c.FbcashDatasetDir
	enc.FbcashDatasetsInMem = c.FbcashDatasetsInMem
	enc.FbcashDatasetsOnDisk = c.FbcashDatasetsOnDisk
	enc.TxPool = c.TxPool
	enc.GPO = c.GPO
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.DocRoot = c.DocRoot
	enc.PowFake = c.PowFake
	enc.PowTest = c.PowTest
	enc.PowShared = c.PowShared
	return &enc, nil
}

func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               *uint64
		SyncMode                *downloader.SyncMode
		LightServ               *int  `toml:",omitempty"`
		LightPeers              *int  `toml:",omitempty"`
		MaxPeers                *int  `toml:"-"`
		SkipBcVersionCheck      *bool `toml:"-"`
		DatabaseHandles         *int  `toml:"-"`
		DatabaseCache           *int
		Fairblockbase               *common.Address `toml:",omitempty"`
		MinerThreads            *int            `toml:",omitempty"`
		ExtraData               hexutil.Bytes   `toml:",omitempty"`
		GasPrice                *big.Int
		FbcashCacheDir          *string
		FbcashCachesInMem       *int
		FbcashCachesOnDisk      *int
		FbcashDatasetDir        *string
		FbcashDatasetsInMem     *int
		FbcashDatasetsOnDisk    *int
		TxPool                  *core.TxPoolConfig
		GPO                     *gasprice.Config
		EnablePreimageRecording *bool
		DocRoot                 *string `toml:"-"`
		PowFake                 *bool   `toml:"-"`
		PowTest                 *bool   `toml:"-"`
		PowShared               *bool   `toml:"-"`
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.LightServ != nil {
		c.LightServ = *dec.LightServ
	}
	if dec.LightPeers != nil {
		c.LightPeers = *dec.LightPeers
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.Fairblockbase != nil {
		c.Fairblockbase = *dec.Fairblockbase
	}
	if dec.MinerThreads != nil {
		c.MinerThreads = *dec.MinerThreads
	}
	if dec.ExtraData != nil {
		c.ExtraData = dec.ExtraData
	}
	if dec.GasPrice != nil {
		c.GasPrice = dec.GasPrice
	}
	if dec.FbcashCacheDir != nil {
		c.FbcashCacheDir = *dec.FbcashCacheDir
	}
	if dec.FbcashCachesInMem != nil {
		c.FbcashCachesInMem = *dec.FbcashCachesInMem
	}
	if dec.FbcashCachesOnDisk != nil {
		c.FbcashCachesOnDisk = *dec.FbcashCachesOnDisk
	}
	if dec.FbcashDatasetDir != nil {
		c.FbcashDatasetDir = *dec.FbcashDatasetDir
	}
	if dec.FbcashDatasetsInMem != nil {
		c.FbcashDatasetsInMem = *dec.FbcashDatasetsInMem
	}
	if dec.FbcashDatasetsOnDisk != nil {
		c.FbcashDatasetsOnDisk = *dec.FbcashDatasetsOnDisk
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	if dec.PowFake != nil {
		c.PowFake = *dec.PowFake
	}
	if dec.PowTest != nil {
		c.PowTest = *dec.PowTest
	}
	if dec.PowShared != nil {
		c.PowShared = *dec.PowShared
	}
	return nil
}
