package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error, but didn't get one")
		}

		if got != want {
			t.Errorf("got %q, but want %q", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		if got != nil {
			t.Error("got an error, but didn't want one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}
		got := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(20))
		assertNoError(t, got)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(30))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}