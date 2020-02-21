package test

import (
	"fmt"
	_ "github.com/MobileCPX/PreDimoco/routers"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestF(t *testing.T) {
	if (false) && true {
		fmt.Println("ssss")
	}
	fmt.Println("le")
}

func TestMo(t *testing.T) {
	now := time.Now()
	aftef := now.AddDate(0, 0, 7)
	fmt.Println(now, aftef, now.After(aftef))
}

func TestMIs(t *testing.T) {
	msisdn := "11234567890"
	fmt.Println(msisdn[len(msisdn)-9:])
}
