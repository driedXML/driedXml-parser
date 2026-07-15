package utils_test

import (
	"testing"

	"github.com/driedxml/parser/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	stack := utils.NewStack()

	require.Truef(t, stack.IsEmpty(), "stack is not empty")
	stack.Push("5")

	require.False(t, stack.IsEmpty(), "stack is empty")

	top, err := stack.Peek()
	require.NoErrorf(t, err, "stack.Top() error")
	require.Equalf(t, "5", top, "stack.Top() error")

	stack.Push("6")
	top, err = stack.Peek()
	require.NoErrorf(t, err, "stack.Top() error")
	require.Equalf(t, "6", top, "stack.Top() error")

	top, err = stack.Pop()
	require.NoErrorf(t, err, "stack.Top() error")
	require.Equalf(t, "6", top, "stack.Top() error")

	top, err = stack.Peek()
	require.NoErrorf(t, err, "stack.Top() error")
	require.Equalf(t, "5", top, "stack.Top() error")

	top, err = stack.Pop()
	require.NoErrorf(t, err, "stack.Top() error")
	require.Equalf(t, "5", top, "stack.Top() error")
	require.Truef(t, stack.IsEmpty(), "stack is not empty")

	top, err = stack.Pop()
	require.Errorf(t, err, "stack.Pop() error")
	require.Nil(t, top, "stack.Pop() error")
}
