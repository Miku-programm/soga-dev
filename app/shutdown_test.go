package app

import (
	"context"
	"reflect"
	"testing"

	"github.com/robfig/cron/v3"
)

func TestShutdownRemovesOnlyRegisteredJobs(t *testing.T) {
	scheduler := cron.New()
	add := func() cron.EntryID {
		id, err := scheduler.AddFunc("@hourly", func() {})
		if err != nil {
			t.Fatal(err)
		}
		return id
	}
	first, retained, last := add(), add(), add()
	ctx, cancel := context.WithCancel(context.Background())
	ids := []cron.EntryID{first, last}
	app := &M0mmPxPJsa{spsw8VktR9z: ctx, un5PcsX7aN53: cancel, bvSWEWNR: scheduler, kxWyUwUkwBr: ids}
	app.InitUnix()
	if ctx.Err() != context.Canceled {
		t.Fatal("context was not canceled")
	}
	entries := scheduler.Entries()
	if len(entries) != 1 || entries[0].ID != retained {
		t.Fatalf("remaining entries: %v", entries)
	}
	if !reflect.DeepEqual(app.kxWyUwUkwBr, ids) {
		t.Fatal("original method retains the ID slice")
	}
	app.InitUnix()
	if entries := scheduler.Entries(); len(entries) != 1 || entries[0].ID != retained {
		t.Fatal(entries)
	}
}

func TestShutdownRequiresInitializedCancel(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("uninitialized original method must panic")
		}
	}()
	new(M0mmPxPJsa).InitUnix()
}
