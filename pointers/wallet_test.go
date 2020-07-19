package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(Bitcoin(3))

		want := Bitcoin(7)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Overdrawn", func(t *testing.T) {
		initialBalance := Bitcoin(10)
		wallet := Wallet{initialBalance}
		err := wallet.Withdraw(Bitcoin(30))

		assertBalance(t, wallet, initialBalance)
		assertError(t, err, ErrBalanceInsufficient)
	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.balance
	if got != want {
		t.Errorf("\ngot: %s\nwant: %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Error("Expected an error but didn't get one")
	}

	if got.Error() != want.Error() {
		t.Errorf("Expected %s but got %s", want.Error(), got.Error())
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Didn't expect error but got: %s", err.Error())
	}
}
