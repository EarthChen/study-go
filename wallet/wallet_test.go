package wallet

import "testing"

func TestWallet_Balance(t *testing.T) {
	type fields struct {
		balance Bitcoin
	}
	tests := []struct {
		name   string
		fields fields
		want   Bitcoin
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				balance: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Wallet{
				balance: tt.fields.balance,
			}
			if got := w.Balance(); got != tt.want {
				t.Errorf("Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWallet_Deposit(t *testing.T) {
	type fields struct {
		balance Bitcoin
	}
	type args struct {
		amount Bitcoin
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: fields{
				balance: 1,
			},
			args: args{
				amount: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Wallet{
				balance: tt.fields.balance,
			}
			w.Deposit(1)
		})
	}
}
