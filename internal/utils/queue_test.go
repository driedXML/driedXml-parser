package utils_test

import (
	"testing"

	"github.com/driedxml/parser/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	queue := utils.NewQueue()

	require.Truef(t, queue.IsEmpty(), "queue is not empty")
	queue.Enqueue("5")

	require.False(t, queue.IsEmpty(), "queue is empty")

	head, err := queue.Peek()
	require.NoErrorf(t, err, "queue.Peek() error")
	require.Equalf(t, "5", head, "queue.Peek() error")

	queue.Enqueue("6")
	head, err = queue.Peek()
	require.NoErrorf(t, err, "queue.Peek() error")
	require.Equalf(t, "5", head, "queue.Peek() error")

	head, err = queue.Dequeue()
	require.NoErrorf(t, err, "queue.Dequeue() error")
	require.Equalf(t, "5", head, "queue.Dequeue() error")
	require.False(t, queue.IsEmpty(), "queue is empty")

	head, err = queue.Peek()
	require.NoErrorf(t, err, "queue.Peek() error")
	require.Equalf(t, "6", head, "queue.Peek() error")

	head, err = queue.Dequeue()
	require.NoErrorf(t, err, "queue.Dequeue() error")
	require.Equalf(t, "6", head, "queue.Dequeue() error")
	require.Truef(t, queue.IsEmpty(), "queue is empty")
}
