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
    b := BetData{}
    b.LoadFromJSON(context.Background())
    fmt.Println(b)
    assert.NotEmpty(t, b)
}

func TestBetData_WriteToJSON(t *testing.T) {
    b := BetData{}
    b.LoadFromJSON(context.Background())
    originStep := b.RunBet.Step
    b.RunBet.Step = originStep + 1
    r := b.WriteToJSON(context.Background())
    assert.True(t, r)
    b.LoadFromJSON(context.Background())
    assert.Equal(t, originStep+1, b.RunBet.Step)
    b.RunBet.Step = originStep
    b.WriteToJSON(context.Background())
}
