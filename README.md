# my clean architecture

## 追記：transactionの書き方
```
func (us *purchaseService) UpdateCoinRecordTx(ctx context.Context, u *model.User) error {
        _, err := us.Tx.DoInTx(ctx, us.updateCoinRecord(u))
        if err != nil {
                return utils.NewInternalServerError(err)
        }
        return nil
}
```