package batchmgr

import (
	"testing"

	"github.com/nulayer/chainstore/memstore"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBatchManager_SizeLimit(t *testing.T) {
	Convey("BatchManager SizeLimit", t, func() {
		mstore := memstore.New(1024 * 1024) // 1MB memory store
		batch := New(BatchRules{
			Size: 5,
		})
		batch.Attach(mstore)

		batch.Put("john", []byte{1, 2, 3})
		batch.Put("will", []byte{4})
		batch.Put("heather", []byte{5, 6, 7})
		batch.Put("mike", []byte{8, 9})

		val, err := mstore.Get("john") // shouldnt run a batch put yet
		So(err, ShouldEqual, nil)
		So(val, ShouldEqual, nil)
	})
}
