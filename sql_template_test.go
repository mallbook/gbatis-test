package gbatis_test

import (
	"testing"

	"github.com/mallbook/gbatis/uuid"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mallbook/gbatis"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSQLTemplate(t *testing.T) {
	id := uuid.New()
	Convey("Test template sql", t, func() {
		Convey("Create table t_mall", func() {
			data := make(map[string]string)
			data["tableName"] = "t_mall2"
			_, err := gbatis.Template("/mapper/mall2.createMallTable", data).Execute()
			So(err, ShouldBeNil)
		})

		Convey("Insert data into t_mall", func() {
			name := "KK Mall"
			avatar := "kkmall.png"
			story := "Hi KK Mall"
			// fmt.Printf("Insert -> id = %s\n", id)
			data := make(map[string]string)
			data["tableName"] = "t_mall2"
			r, err := gbatis.Template("/mapper/mall2.insertMall", data).Execute(id, name, avatar, story)
			So(err, ShouldBeNil)
			ra, err := r.RowsAffected()
			So(err, ShouldBeNil)
			So(ra, ShouldEqual, 1)
		})

		Convey("Update mall data", func() {
			name := "MH Mall"
			// fmt.Printf("Update -> id=%s\n", id)
			data := make(map[string]string)
			data["tableName"] = "t_mall2"
			r, err := gbatis.Template("/mapper/mall2.updateMall", data).Execute(name, id)
			So(err, ShouldBeNil)
			ra, err := r.RowsAffected()
			So(err, ShouldBeNil)
			So(ra, ShouldEqual, 1)
		})

		Convey("Select mall data", func() {
			// fmt.Printf("Select -> id=%s\n", id)
			data := make(map[string]string)
			data["tableName"] = "t_mall2"
			row, err := gbatis.Template("/mapper/mall2.selectMall", data).SelectRow(id)
			So(err, ShouldBeNil)
			So(row, ShouldNotBeNil)
			var mallID string
			var name string
			var avatar string
			var createdAt int32
			var updatedAt int32
			var story string
			err = row.Scan(&mallID, &name, &avatar, &createdAt, &updatedAt, &story)
			So(err, ShouldBeNil)
			So(mallID, ShouldEqual, id.String())
			So(name, ShouldEqual, "MH Mall")
			So(avatar, ShouldEqual, "kkmall.png")
			So(story, ShouldEqual, "Hi KK Mall")
		})

		Convey("Delete mall data", func() {
			// fmt.Printf("Delete -> id=%s", id)
			data := make(map[string]string)
			data["tableName"] = "t_mall2"
			r, err := gbatis.Template("/mapper/mall2.deleteMall", data).Execute(id)
			So(err, ShouldBeNil)
			ra, err := r.RowsAffected()
			So(err, ShouldBeNil)
			So(ra, ShouldEqual, 1)
		})

		Convey("Drop table t_mall", func() {
			data := make(map[string]string)
			data["tableName"] = "t_mall2"
			_, err := gbatis.Template("/mapper/mall2.dropMallTable", data).Execute()
			So(err, ShouldBeNil)
		})
	})
}
