package boltdb

import (
	"testing"

	. "github.com/nulayer/chainstore"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBoltdbStore(t *testing.T) {
	var store Store
	var err error

	store, err = NewStore(TempDir()+"/test.db", "test")
	defer store.Close()
	if err != nil {
		t.Error(err)
	}

	Convey("Boltdb Open", t, func() {

		Convey("Put a bunch of objects", func() {
			e1 := store.Put("hi", []byte{1, 2, 3})
			e2 := store.Put("bye", []byte{4, 5, 6})
			So(e1, ShouldEqual, nil)
			So(e2, ShouldEqual, nil)
		})

		Convey("Get those objects", func() {
			v1, _ := store.Get("hi")
			v2, _ := store.Get("bye")
			So(v1, ShouldResemble, []byte{1, 2, 3})
			So(v2, ShouldResemble, []byte{4, 5, 6})
		})

		Convey("Delete those objects", func() {
			e1 := store.Del("hi")
			e2 := store.Del("bye")
			So(e1, ShouldEqual, nil)
			So(e2, ShouldEqual, nil)

			v, _ := store.Get("hi")
			So(len(v), ShouldEqual, 0)
		})

		Convey("Disallow invalid keys", func() {
			err = store.Put("test!!!", []byte{1})
			So(err, ShouldEqual, ErrInvalidKey)
		})

	})
}
