package memorystorage

import (
	"context"
	"testing"
	"time"

	"github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	ctx := context.Background()

	t.Run("add", func(t *testing.T) {
		stor := New()
		ev := storage.Event{
			ID:    uuid.NewString(),
			Title: "test",
		}
		err := stor.Add(ctx, ev)
		require.NoError(t, err)
		require.NotEqual(t, "", ev.ID)
	})

	t.Run("delete", func(t *testing.T) {
		stor := New()
		ev := storage.Event{
			ID:    uuid.NewString(),
			Title: "test",
		}
		err := stor.Add(ctx, ev)
		require.NoError(t, err)
		require.NotEqual(t, "", ev.ID)

		err = stor.Delete(ctx, ev)
		require.NoError(t, err)
	})

	t.Run("update", func(t *testing.T) {
		stor := New()
		ev := storage.Event{
			ID:    uuid.NewString(),
			Title: "test",
		}
		err := stor.Add(ctx, ev)
		require.NoError(t, err)
		require.NotEqual(t, "", ev.ID)

		updatedEv := storage.Event{
			ID:    ev.ID,
			Title: "updated",
		}
		err = stor.Update(ctx, updatedEv)
		require.NoError(t, err)
	})

	t.Run("get list", func(t *testing.T) {
		stor := New()
		ev := storage.Event{
			ID:        uuid.NewString(),
			Title:     "time.Minute",
			EventDate: time.Now().Add(time.Minute),
		}
		se := storage.Event{
			ID:        uuid.NewString(),
			Title:     "time.Minute * 2",
			EventDate: time.Now().Add(time.Minute * 2),
		}

		err := stor.Add(ctx, ev)
		require.NoError(t, err)

		err = stor.Add(ctx, se)
		require.NoError(t, err)

		res, err := stor.ListDayEvents(ctx, time.Now())
		require.NoError(t, err)
		require.Equal(t, 2, len(res))
		require.Equal(t, res, []storage.Event{ev, se})
	})
}
