package CbO_7wp3B

import (
	"reflect"
	"testing"

	"github.com/robfig/cron/v3"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/k_qeiAYoxa"
)

type schedulerFixture struct{ spec string }

func (j schedulerFixture) Cron() string { return j.spec }
func (j schedulerFixture) Run()         {}

func TestSchedulerRTTIAndActualDI(t *testing.T) {
	rt := reflect.TypeFor[M0mmPxPJsa]()
	if rt.Size() != 80 {
		t.Fatal(rt.Size())
	}
	for i, want := range []uintptr{0, 16, 24, 32, 56} {
		if rt.Field(i).Offset != want {
			t.Fatal(rt.Field(i))
		}
	}
	input := reflect.TypeFor[C7gpSMD8r]()
	if input.Size() != 32 || input.Field(0).Name != "WdkJetA" || !input.Field(0).Anonymous || input.Field(2).Tag.Get("group") != "JOBS" {
		t.Fatal(input)
	}
	if err := cNT24rG4F.DependencyContainer().Invoke(func(app *M0mmPxPJsa, c *cron.Cron) {
		if app.bvSWEWNR != c || app.spsw8VktR9z != nil || app.un5PcsX7aN53 != nil || app.kxWyUwUkwBr != nil {
			t.Fatal("constructor state")
		}
		if c.Location().String() != "Asia/Shanghai" {
			t.Fatal(c.Location())
		}
		if _, err := c.AddJob("* * * * *", schedulerFixture{}); err == nil {
			t.Fatal("original parser requires seconds")
		}
		id, err := c.AddJob("0 0 0 1 1 *", schedulerFixture{})
		if err != nil {
			t.Fatal(err)
		}
		c.Remove(id)
	}); err != nil {
		t.Fatal(err)
	}
}

func TestJobRegistrationPreservesPartialAndPriorEntries(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	app := lKahKJ0kn(C7gpSMD8r{XwSaXtyHwQR: c, E7fCaMMT: []k_qeiAYoxa.YNfyrQAMx{schedulerFixture{"@every 1h"}, schedulerFixture{"invalid"}, schedulerFixture{"@every 2h"}}})
	if err := app.vVeTh4DwS(); err == nil {
		t.Fatal("invalid schedule accepted")
	}
	if len(app.kxWyUwUkwBr) != 1 || len(c.Entries()) != 1 {
		t.Fatal("partial successful job was removed")
	}
	old := app.kxWyUwUkwBr[0]
	app.jicpAs = app.jicpAs[:1]
	if err := app.vVeTh4DwS(); err != nil {
		t.Fatal(err)
	}
	if len(app.kxWyUwUkwBr) != 1 || app.kxWyUwUkwBr[0] == old || len(c.Entries()) != 2 {
		t.Fatal("previous jobs unexpectedly removed")
	}
	app.jicpAs = nil
	if err := app.vVeTh4DwS(); err != nil {
		t.Fatal(err)
	}
	if app.kxWyUwUkwBr == nil || len(app.kxWyUwUkwBr) != 0 || len(c.Entries()) != 2 {
		t.Fatal("empty registration behavior")
	}
}
