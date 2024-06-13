package types

import "fmt"

type Bill struct {
	ID          string
	Name        string
	Description string
	Amount      string
	Installment int64
}

func (b *Bill) IsInstallmentHealthy() error {

	if b.Installment > 128 {
		return fmt.Errorf("installment doesnt be increase 128")
	}

	return nil

}
