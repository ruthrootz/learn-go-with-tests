package pointersAndErrors

import (
  "testing"
)

func TestWallet(t *testing.T) {

  t.Run("deposit", func(t *testing.T) {
    wallet := Wallet{}
    wallet.Deposit(Bitcoin(10))
    got := wallet.Balance()
    want := Bitcoin(10)

    if got != want {
      t.Errorf("got %s want %s", got, want)
    }
  })

  t.Run("withdraw", func(t *testing.T) {
    wallet := Wallet{balance: Bitcoin(20)}
    wallet.Withdraw(Bitcoin(10))
    got := wallet.Balance()
    want := Bitcoin(10)

    if got != want {
      t.Errorf("got %s want %s", got, want)
    }
  })

  t.Run("withdraw insufficient funds", func(t *testing.T) {
    startingBalance := Bitcoin(20)
    wallet := Wallet{startingBalance}
    err := wallet.Withdraw(Bitcoin(100))

    if err == nil {
      t.Error("wanted an error but didn't get one")
    }
  })

}

