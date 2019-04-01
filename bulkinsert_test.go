package gbatis_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mallbook/gbatis"

	"github.com/mallbook/gbatis/uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBulkInsert(t *testing.T) {
	Convey("Test bulk insert 10", t, func() {
		Convey("Create table t_mall", func() {
			_, err := gbatis.Execute("/mapper/mall.createMallTable")
			So(err, ShouldBeNil)
		})

		Convey("Bulk insert 10", func() {
			rows := make([]interface{}, 0)
			for i := 0; i < 10; i++ {
				row := make([]interface{}, 0)
				u := uuid.New()
				row = append(row, u.String(), "a mall", "a.png", "hi a mall")
				rows = append(rows, row)
			}

			r, err := gbatis.BulkInsert("/mapper/mall.insertMall", rows)
			So(err, ShouldBeNil)
			ra, err := r.RowsAffected()
			So(err, ShouldBeNil)
			So(ra, ShouldEqual, 10)
		})

		Convey("Drop table t_mall", func() {
			_, err := gbatis.Execute("/mapper/mall.dropMallTable")
			So(err, ShouldBeNil)
		})
	})

	Convey("Test bulk insert 100", t, func() {
		Convey("Create table t_mall", func() {
			_, err := gbatis.Execute("/mapper/mall.createMallTable")
			So(err, ShouldBeNil)
		})

		Convey("Bulk insert 100", func() {
			rows := make([]interface{}, 0)
			for i := 0; i < 100; i++ {
				row := make([]interface{}, 0)
				u := uuid.New()
				row = append(row, u.String(), "a mall", "a.png", "hi a mall")
				rows = append(rows, row)
			}

			r, err := gbatis.BulkInsert("/mapper/mall.insertMall", rows)
			So(err, ShouldBeNil)
			ra, err := r.RowsAffected()
			So(err, ShouldBeNil)
			So(ra, ShouldEqual, 100)
		})

		Convey("Drop table t_mall", func() {
			_, err := gbatis.Execute("/mapper/mall.dropMallTable")
			So(err, ShouldBeNil)
		})
	})

	Convey("Test bulk insert 1000", t, func() {
		Convey("Create table t_mall", func() {
			_, err := gbatis.Execute("/mapper/mall.createMallTable")
			So(err, ShouldBeNil)
		})

		Convey("Bulk insert 1000", func() {
			rows := make([]interface{}, 0)
			for i := 0; i < 1000; i++ {
				row := make([]interface{}, 0)
				u := uuid.New()
				row = append(row, u.String(), "a mall", "a.png", "hi a mall")
				rows = append(rows, row)
			}

			r, err := gbatis.BulkInsert("/mapper/mall.insertMall", rows)
			So(err, ShouldBeNil)
			ra, err := r.RowsAffected()
			So(err, ShouldBeNil)
			So(ra, ShouldEqual, 1000)
		})

		Convey("Drop table t_mall", func() {
			_, err := gbatis.Execute("/mapper/mall.dropMallTable")
			So(err, ShouldBeNil)
		})
	})

	Convey("Test bulk insert 10000", t, func() {
		Convey("Create table t_mall", func() {
			_, err := gbatis.Execute("/mapper/mall.createMallTable")
			So(err, ShouldBeNil)
		})

		Convey("Bulk insert 10000", func() {
			rows := make([]interface{}, 0)
			for i := 0; i < 10000; i++ {
				row := make([]interface{}, 0)
				u := uuid.New()
				row = append(row, u.String(), "a mall", "a.png", "hi a mall")
				rows = append(rows, row)
			}

			r, err := gbatis.BulkInsert("/mapper/mall.insertMall", rows)
			So(err, ShouldBeNil)
			ra, err := r.RowsAffected()
			So(err, ShouldBeNil)
			So(ra, ShouldEqual, 10000)
		})

		Convey("Drop table t_mall", func() {
			_, err := gbatis.Execute("/mapper/mall.dropMallTable")
			So(err, ShouldBeNil)
		})
	})
}
