package todo

import "testing"

func TestNewTask(t *testing.T) {
	want := &Task{
		info:   "Buy Milk",
		status: false,
	}

	got := NewTask("Buy Milk", false)
	if got.info != want.info {
		t.Errorf("NewTask() => got.info(%v) != want.info(%v)", got.info, want.info)
	}

	if got.status != want.status {
		t.Errorf("NewTask() => got.status(%v) != want.status(%v)", got.status, want.status)
	}

	if len(got.id) == 0 {
		t.Error("NewTask() => got.id isn't initiliazed")
	}
}
