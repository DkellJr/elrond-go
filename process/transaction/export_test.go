package transaction

import (
	"math/big"

	"github.com/ElrondNetwork/elrond-go/data/state"
	"github.com/ElrondNetwork/elrond-go/data/transaction"
	"github.com/ElrondNetwork/elrond-go/process"
)

type TxProcessor *txProcessor

func (txProc *txProcessor) GetAccounts(adrSrc, adrDst []byte,
) (acntSrc, acntDst state.UserAccountHandler, err error) {
	return txProc.getAccounts(adrSrc, adrDst)
}

func (txProc *txProcessor) CheckTxValues(tx *transaction.Transaction, acntSnd, acntDst state.UserAccountHandler) error {
	return txProc.checkTxValues(tx, acntSnd, acntDst)
}

func (txProc *txProcessor) IncreaseNonce(acntSrc state.UserAccountHandler) {
	acntSrc.IncreaseNonce(1)
}

func (txProc *txProcessor) ProcessTxFee(
	tx *transaction.Transaction,
	acntSnd, acntDst state.UserAccountHandler,
	cost *big.Int,
) (*big.Int, error) {
	return txProc.processTxFee(tx, acntSnd, acntDst, cost)
}

func (inTx *InterceptedTransaction) SetWhitelistHandler(handler process.WhiteListHandler) {
	inTx.whiteListerVerifiedTxs = handler
}

func (txProc *txProcessor) GetUserTxCost(
	userTx *transaction.Transaction,
	userTxHash []byte,
	userTxType process.TransactionType,
) *big.Int {
	return txProc.getUserTxCost(userTx, userTxHash, userTxType)
}

func (txProc *baseTxProcessor) IsCrossTxFromMe(adrSrc, adrDst []byte) bool {
	return txProc.isCrossTxFromMe(adrSrc, adrDst)
}

func (txProc *txProcessor) SetPenalizedTooMuchGasEnableEpoch(epoch uint32) {
	txProc.penalizedTooMuchGasEnableEpoch = epoch
}
