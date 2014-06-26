package batchmgr

import (
	"time"

	"github.com/nulayer/chainstore"
	"github.com/oleiade/lane"
)

type item struct {
	key string
	val []byte
}

type BatchManager struct {
	*chainstore.DefaultManager

	numLimit  int64
	currNum   int64
	sizeLimit int64
	currSize  int64
	timeout   time.Duration
	batch     *lane.Queue

	Error error
}

type BatchRules struct {
	Num     int64
	Size    int64
	Timeout time.Duration
}

// currently only batch PUTs
// TODO: batch GETs
func New(rules BatchRules) *BatchManager {
	manager := &BatchManager{
		batch: lane.NewQueue(),
	}

	if rules.Num != 0 {
		manager.num = rules.Num
	}
	if rules.Size != 0 {
		manager.size = rules.Size
	}
	if rules.Timeout != 0 {
		manager.timeout = rules.Timeout
		ticker := time.NewTicker(manager.timeout)
		go func() {
			for {
				<-ticker.C // wait for timeout
				err := manager.flush()
				if err != nil {
					manager.Error = err
				}
			}
		}()
	}

	return manager
}

func (b *BatchManager) flush() (err error) {
	if b.batch.Head() != nil {
		for i := b.batch.Dequeue(); i != nil; {
			op := i.(item)
			err := b.Chain.Put(op.key, op.val)
			if err != nil {
				return err
			}
		}
	}
}

func (b *BatchManager) Open() (err error) { return }

func (b *BatchManager) Close() (err error) { return }

func (b *BatchManager) Put(key string, obj []byte) (err error) {
	if !chainstore.IsValidKey(key) {
		return chainstore.ErrInvalidKey
	}

	op := item{key, obj}
	b.batch.Enqueue(op)
	b.currNum++
	b.currSize += len(obj)

	if b.currNum >= b.numLimit || b.currSize >= b.sizeLimit {
		err = b.flush()
	}

	return
}

func (b *BatchManager) Get(key string) (obj []byte, err error) {
	return b.Chain.Get(key)
}

func (b *BatchManager) Del(key string) (err error) {
	return b.Chain.Del(key)
}
