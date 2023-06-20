package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeTransactionList interface {
	GetAllTransactions() []ETradeTransaction
	GetTransactionById(transactionID int64) ETradeTransaction
}

type eTradeTransactionList struct {
	transactions []ETradeTransaction
}

const (
	// The transaction list response JSON looks like this:
	// {
	//   "TransactionListResponse": {
	//     "Transaction": [
	//       {
	//         <transaction info>
	//       }
	//     ]
	//   }
	// }

	// transactionsSliceResponsePath is the path to a slice of transactions.
	transactionsSliceResponsePath = "transactionListResponse.transaction"
)

func CreateETradeTransactionList(transactionListResponseMap jsonmap.JsonMap) (ETradeTransactionList, error) {
	transactionsSlice, err := transactionListResponseMap.GetSliceOfMapsAtPath(transactionsSliceResponsePath)
	if err != nil {
		return nil, err
	}
	allTransactions := make([]ETradeTransaction, 0, len(transactionsSlice))
	for _, transactionInfoMap := range transactionsSlice {
		transaction, err := CreateETradeTransaction(transactionInfoMap)
		if err != nil {
			return nil, err
		}
		allTransactions = append(allTransactions, transaction)
	}
	return &eTradeTransactionList{transactions: allTransactions}, nil
}

func (e *eTradeTransactionList) GetAllTransactions() []ETradeTransaction {
	return e.transactions
}

func (e *eTradeTransactionList) GetTransactionById(transactionID int64) ETradeTransaction {
	for _, transaction := range e.transactions {
		if transaction.GetId() == transactionID {
			return transaction
		}
	}
	return nil
}
