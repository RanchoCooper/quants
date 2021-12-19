package bet

import (
    "context"
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2021/12/3
 */

func TestBetData_LoadFromJSON(t *testing.T) {
    gridBet := NewGridBet()
    gridBet.Grid.LoadFromJSON(context.Background())
    fmt.Println(gridBet)
    assert.NotEmpty(t, gridBet)
}

func TestBetData_WriteToJSON(t *testing.T) {
    gridBet := NewGridBet()
    gridBet.Grid.LoadFromJSON(context.Background())
    originStep := gridBet.Grid.Step
    gridBet.Grid.Step = originStep + 1
    r := gridBet.Grid.WriteToJSON(context.Background())
    assert.True(t, r)
    gridBet.Grid.LoadFromJSON(context.Background())
    assert.Equal(t, originStep+1, gridBet.Grid.Step)
    gridBet.Grid.Step = originStep
    gridBet.Grid.WriteToJSON(context.Background())
}
