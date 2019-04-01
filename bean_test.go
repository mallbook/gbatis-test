package gbatis_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mallbook/gbatis"
	"github.com/mallbook/gbatis/bean"
	"github.com/mallbook/gbatis/uuid"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMultiBean(t *testing.T) {
	Convey("Test multiple beans", t, func() {
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

		Convey("Select all datas", func() {
			malls, err := gbatis.Select("/mapper/mall.selectAllMalls")
			So(err, ShouldBeNil)
			So(len(malls), ShouldEqual, 10)
		})

		Convey("Drop table t_mall", func() {
			_, err := gbatis.Execute("/mapper/mall.dropMallTable")
			So(err, ShouldBeNil)
		})
	})
}

func TestOneBean(t *testing.T) {
	id := uuid.New()
	Convey("Test one bean", t, func() {
		Convey("Create table t_mall", func() {
			_, err := gbatis.Execute("/mapper/mall.createMallTable")
			So(err, ShouldBeNil)
		})

		Convey("Insert data into t_mall", func() {
			name := "KK Mall"
			avatar := "kkmall.png"
			story := "Hi KK Mall"
			// fmt.Printf("Insert -> id = %s\n", id)
			r, err := gbatis.Execute("/mapper/mall.insertMall", id, name, avatar, story)
			So(err, ShouldBeNil)
			ra, err := r.RowsAffected()
			So(err, ShouldBeNil)
			So(ra, ShouldEqual, 1)
		})

		Convey("Select one object", func() {
			mall, err := gbatis.SelectOne("/mapper/mall.selectAllMalls")
			So(err, ShouldBeNil)
			So(mall, ShouldNotBeNil)
			m, ok := mall.(bean.Mall)
			So(ok, ShouldBeTrue)
			So(m.ID, ShouldEqual, id.String())
			So(m.Name, ShouldEqual, "KK Mall")
			So(m.Avatar, ShouldEqual, "kkmall.png")
			So(m.Story, ShouldEqual, "Hi KK Mall")
		})

		Convey("Drop table t_mall", func() {
			_, err := gbatis.Execute("/mapper/mall.dropMallTable")
			So(err, ShouldBeNil)
		})
	})
}
