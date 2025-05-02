package tf

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/Snikimonkd/tf/assert"
	opts "github.com/Snikimonkd/tf/options"
)

type User struct {
	CreatedAt time.Time
}

func Test_User(t *testing.T) {
	t.Parallel()
	berlin, err := time.LoadLocation("Europe/Berlin")
	assert.Nil(t, err)

	now := time.Now()
	u1 := User{
		CreatedAt: now.In(berlin),
	}
	u2 := User{
		CreatedAt: now.In(time.UTC),
	}
	assert.Equal(t, u1, u2)
}

func Test_Slices(t *testing.T) {
	t.Parallel()
	want := []int{1, 2, 3}
	got := []int{3, 2, 1}

	assert.Equal(
		t,
		want,
		got,
		cmpopts.SortSlices(func(x, y int) bool {
			return x < y
		}),
	)
}

func Test_SlicesV2(t *testing.T) {
	t.Parallel()
	want := []int{1, 2, 3}
	got := []int{3, 2, 1}

	assert.Equal(t, want, got, opts.SortSlicesOpt[int]())
}
