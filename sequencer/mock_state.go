// Code generated by mockery v2.22.1. DO NOT EDIT.

package sequencer

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	executor "github.com/0xPolygonHermez/zkevm-node/state/runtime/executor"

	metrics "github.com/0xPolygonHermez/zkevm-node/state/metrics"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	state "github.com/0xPolygonHermez/zkevm-node/state"

	time "time"

	types "github.com/ethereum/go-ethereum/core/types"
)

// StateMock is an autogenerated mock type for the stateInterface type
type StateMock struct {
	mock.Mock
}

// Begin provides a mock function with given fields: ctx
func (_m *StateMock) Begin(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeginStateTransaction provides a mock function with given fields: ctx
func (_m *StateMock) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseBatch provides a mock function with given fields: ctx, receipt, dbTx
func (_m *StateMock) CloseBatch(ctx context.Context, receipt state.ProcessingReceipt, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, receipt, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingReceipt, pgx.Tx) error); ok {
		r0 = rf(ctx, receipt, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountReorgs provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) CountReorgs(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteBatch provides a mock function with given fields: ctx, batch, updateMerkleTree, dbTx
func (_m *StateMock) ExecuteBatch(ctx context.Context, batch state.Batch, updateMerkleTree bool, dbTx pgx.Tx) (*executor.ProcessBatchResponse, error) {
	ret := _m.Called(ctx, batch, updateMerkleTree, dbTx)

	var r0 *executor.ProcessBatchResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, state.Batch, bool, pgx.Tx) (*executor.ProcessBatchResponse, error)); ok {
		return rf(ctx, batch, updateMerkleTree, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, state.Batch, bool, pgx.Tx) *executor.ProcessBatchResponse); ok {
		r0 = rf(ctx, batch, updateMerkleTree, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*executor.ProcessBatchResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, state.Batch, bool, pgx.Tx) error); ok {
		r1 = rf(ctx, batch, updateMerkleTree, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FlushMerkleTree provides a mock function with given fields: ctx
func (_m *StateMock) FlushMerkleTree(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBalanceByStateRoot provides a mock function with given fields: ctx, address, root
func (_m *StateMock) GetBalanceByStateRoot(ctx context.Context, address common.Address, root common.Hash) (*big.Int, error) {
	ret := _m.Called(ctx, address, root)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) (*big.Int, error)); ok {
		return rf(ctx, address, root)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) *big.Int); ok {
		r0 = rf(ctx, address, root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, common.Hash) error); ok {
		r1 = rf(ctx, address, root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchByNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 *state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.Batch, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Batch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForcedBatch provides a mock function with given fields: ctx, forcedBatchNumber, dbTx
func (_m *StateMock) GetForcedBatch(ctx context.Context, forcedBatchNumber uint64, dbTx pgx.Tx) (*state.ForcedBatch, error) {
	ret := _m.Called(ctx, forcedBatchNumber, dbTx)

	var r0 *state.ForcedBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.ForcedBatch, error)); ok {
		return rf(ctx, forcedBatchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.ForcedBatch); ok {
		r0 = rf(ctx, forcedBatchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.ForcedBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, forcedBatchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForcedBatchesSince provides a mock function with given fields: ctx, forcedBatchNumber, maxBlockNumber, dbTx
func (_m *StateMock) GetForcedBatchesSince(ctx context.Context, forcedBatchNumber uint64, maxBlockNumber uint64, dbTx pgx.Tx) ([]*state.ForcedBatch, error) {
	ret := _m.Called(ctx, forcedBatchNumber, maxBlockNumber, dbTx)

	var r0 []*state.ForcedBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) ([]*state.ForcedBatch, error)); ok {
		return rf(ctx, forcedBatchNumber, maxBlockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) []*state.ForcedBatch); ok {
		r0 = rf(ctx, forcedBatchNumber, maxBlockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*state.ForcedBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, forcedBatchNumber, maxBlockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForkIDByBatchNumber provides a mock function with given fields: batchNumber
func (_m *StateMock) GetForkIDByBatchNumber(batchNumber uint64) uint64 {
	ret := _m.Called(batchNumber)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint64) uint64); ok {
		r0 = rf(batchNumber)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetLastBatch provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastBatch(ctx context.Context, dbTx pgx.Tx) (*state.Batch, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.Batch, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.Batch); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBatchNumber provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBlock provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastBlock(ctx context.Context, dbTx pgx.Tx) (*state.Block, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.Block, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.Block); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastClosedBatch provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastClosedBatch(ctx context.Context, dbTx pgx.Tx) (*state.Batch, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.Batch, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.Batch); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastL2Block provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastL2Block(ctx context.Context, dbTx pgx.Tx) (*types.Block, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*types.Block, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *types.Block); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastL2BlockHeader provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastL2BlockHeader(ctx context.Context, dbTx pgx.Tx) (*types.Header, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*types.Header, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *types.Header); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastNBatches provides a mock function with given fields: ctx, numBatches, dbTx
func (_m *StateMock) GetLastNBatches(ctx context.Context, numBatches uint, dbTx pgx.Tx) ([]*state.Batch, error) {
	ret := _m.Called(ctx, numBatches, dbTx)

	var r0 []*state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, pgx.Tx) ([]*state.Batch, error)); ok {
		return rf(ctx, numBatches, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint, pgx.Tx) []*state.Batch); ok {
		r0 = rf(ctx, numBatches, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint, pgx.Tx) error); ok {
		r1 = rf(ctx, numBatches, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastStateRoot provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastStateRoot(ctx context.Context, dbTx pgx.Tx) (common.Hash, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (common.Hash, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) common.Hash); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastTrustedForcedBatchNumber provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastTrustedForcedBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastVirtualBatchNum provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastVirtualBatchNum(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestGer provides a mock function with given fields: ctx, maxBlockNumber
func (_m *StateMock) GetLatestGer(ctx context.Context, maxBlockNumber uint64) (state.GlobalExitRoot, time.Time, error) {
	ret := _m.Called(ctx, maxBlockNumber)

	var r0 state.GlobalExitRoot
	var r1 time.Time
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (state.GlobalExitRoot, time.Time, error)); ok {
		return rf(ctx, maxBlockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) state.GlobalExitRoot); ok {
		r0 = rf(ctx, maxBlockNumber)
	} else {
		r0 = ret.Get(0).(state.GlobalExitRoot)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) time.Time); ok {
		r1 = rf(ctx, maxBlockNumber)
	} else {
		r1 = ret.Get(1).(time.Time)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64) error); ok {
		r2 = rf(ctx, maxBlockNumber)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLatestGlobalExitRoot provides a mock function with given fields: ctx, maxBlockNumber, dbTx
func (_m *StateMock) GetLatestGlobalExitRoot(ctx context.Context, maxBlockNumber uint64, dbTx pgx.Tx) (state.GlobalExitRoot, time.Time, error) {
	ret := _m.Called(ctx, maxBlockNumber, dbTx)

	var r0 state.GlobalExitRoot
	var r1 time.Time
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (state.GlobalExitRoot, time.Time, error)); ok {
		return rf(ctx, maxBlockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) state.GlobalExitRoot); ok {
		r0 = rf(ctx, maxBlockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(state.GlobalExitRoot)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) time.Time); ok {
		r1 = rf(ctx, maxBlockNumber, dbTx)
	} else {
		r1 = ret.Get(1).(time.Time)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64, pgx.Tx) error); ok {
		r2 = rf(ctx, maxBlockNumber, dbTx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLatestVirtualBatchTimestamp provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLatestVirtualBatchTimestamp(ctx context.Context, dbTx pgx.Tx) (time.Time, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (time.Time, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) time.Time); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonceByStateRoot provides a mock function with given fields: ctx, address, root
func (_m *StateMock) GetNonceByStateRoot(ctx context.Context, address common.Address, root common.Hash) (*big.Int, error) {
	ret := _m.Called(ctx, address, root)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) (*big.Int, error)); ok {
		return rf(ctx, address, root)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) *big.Int); ok {
		r0 = rf(ctx, address, root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, common.Hash) error); ok {
		r1 = rf(ctx, address, root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStoredFlushID provides a mock function with given fields: ctx
func (_m *StateMock) GetStoredFlushID(ctx context.Context) (uint64, string, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) string); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTimeForLatestBatchVirtualization provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetTimeForLatestBatchVirtualization(ctx context.Context, dbTx pgx.Tx) (time.Time, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (time.Time, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) time.Time); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) GetTransactionsByBatchNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) ([]types.Transaction, []uint8, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 []types.Transaction
	var r1 []uint8
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) ([]types.Transaction, []uint8, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) []types.Transaction); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) []uint8); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]uint8)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64, pgx.Tx) error); ok {
		r2 = rf(ctx, batchNumber, dbTx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTxsOlderThanNL1Blocks provides a mock function with given fields: ctx, nL1Blocks, dbTx
func (_m *StateMock) GetTxsOlderThanNL1Blocks(ctx context.Context, nL1Blocks uint64, dbTx pgx.Tx) ([]common.Hash, error) {
	ret := _m.Called(ctx, nL1Blocks, dbTx)

	var r0 []common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) ([]common.Hash, error)); ok {
		return rf(ctx, nL1Blocks, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) []common.Hash); ok {
		r0 = rf(ctx, nL1Blocks, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, nL1Blocks, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsBatchClosed provides a mock function with given fields: ctx, batchNum, dbTx
func (_m *StateMock) IsBatchClosed(ctx context.Context, batchNum uint64, dbTx pgx.Tx) (bool, error) {
	ret := _m.Called(ctx, batchNum, dbTx)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (bool, error)); ok {
		return rf(ctx, batchNum, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) bool); ok {
		r0 = rf(ctx, batchNum, dbTx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNum, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenBatch provides a mock function with given fields: ctx, processingContext, dbTx
func (_m *StateMock) OpenBatch(ctx context.Context, processingContext state.ProcessingContext, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, processingContext, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingContext, pgx.Tx) error); ok {
		r0 = rf(ctx, processingContext, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessBatch provides a mock function with given fields: ctx, request, updateMerkleTree
func (_m *StateMock) ProcessBatch(ctx context.Context, request state.ProcessRequest, updateMerkleTree bool) (*state.ProcessBatchResponse, error) {
	ret := _m.Called(ctx, request, updateMerkleTree)

	var r0 *state.ProcessBatchResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessRequest, bool) (*state.ProcessBatchResponse, error)); ok {
		return rf(ctx, request, updateMerkleTree)
	}
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessRequest, bool) *state.ProcessBatchResponse); ok {
		r0 = rf(ctx, request, updateMerkleTree)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.ProcessBatchResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, state.ProcessRequest, bool) error); ok {
		r1 = rf(ctx, request, updateMerkleTree)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessSequencerBatch provides a mock function with given fields: ctx, batchNumber, batchL2Data, caller, dbTx
func (_m *StateMock) ProcessSequencerBatch(ctx context.Context, batchNumber uint64, batchL2Data []byte, caller metrics.CallerLabel, dbTx pgx.Tx) (*state.ProcessBatchResponse, error) {
	ret := _m.Called(ctx, batchNumber, batchL2Data, caller, dbTx)

	var r0 *state.ProcessBatchResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []byte, metrics.CallerLabel, pgx.Tx) (*state.ProcessBatchResponse, error)); ok {
		return rf(ctx, batchNumber, batchL2Data, caller, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []byte, metrics.CallerLabel, pgx.Tx) *state.ProcessBatchResponse); ok {
		r0 = rf(ctx, batchNumber, batchL2Data, caller, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.ProcessBatchResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, []byte, metrics.CallerLabel, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, batchL2Data, caller, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreTransaction provides a mock function with given fields: ctx, batchNumber, processedTx, coinbase, timestamp, dbTx
func (_m *StateMock) StoreTransaction(ctx context.Context, batchNumber uint64, processedTx *state.ProcessTransactionResponse, coinbase common.Address, timestamp uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, processedTx, coinbase, timestamp, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *state.ProcessTransactionResponse, common.Address, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, processedTx, coinbase, timestamp, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBatchL2DataAndLER provides a mock function with given fields: ctx, batchNumber, batchL2Data, localExitRoot, dbTx
func (_m *StateMock) UpdateBatchL2DataAndLER(ctx context.Context, batchNumber uint64, batchL2Data []byte, localExitRoot common.Hash, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, batchL2Data, localExitRoot, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []byte, common.Hash, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, batchL2Data, localExitRoot, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStateMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewStateMock creates a new instance of StateMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStateMock(t mockConstructorTestingTNewStateMock) *StateMock {
	mock := &StateMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
