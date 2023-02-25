package logic_test

import (
	"github/austinmorales/moneyme/logic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateSymbol(t *testing.T) {
	require.NotNil(t, logic.StockPrice("goog"), "shouldn't be nil")
}
