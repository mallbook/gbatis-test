package gbatis_test

import (
	"testing"

	"github.com/mallbook/gbatis/uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUUID(t *testing.T) {
	u := uuid.New()
	id1 := u.String()
	// fmt.Printf("id1=%s\n", id1)
	Convey("Test UUID", t, func() {
		Convey("Check id1 equal id2", func() {
			id2 := u.String()
			// fmt.Printf("id1=%s\n", id1)
			// fmt.Printf("id2=%s\n", id2)
			So(id1, ShouldEqual, id2)
		})

		Convey("Check id1 equal id3", func() {
			id3 := u.String()
			// fmt.Printf("id1=%s\n", id1)
			// fmt.Printf("id3=%s\n", id3)
			So(id1, ShouldEqual, id3)
		})

		Convey("Check id1 equal id4", func() {
			id4 := u.String()
			// fmt.Printf("id1=%s\n", id1)
			// fmt.Printf("id4=%s\n", id4)
			So(id1, ShouldEqual, id4)
		})
	})
}
