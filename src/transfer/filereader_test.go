package transfer

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		var d TransData
		d = TransFile{}
		f, _ := os.OpenFile("test", os.O_RDONLY, 0)
		fw, _ := os.OpenFile("wtest", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		err := d.Read(f)
		err = d.Write(fw)
		So(err, ShouldEqual, nil)
	})
}
