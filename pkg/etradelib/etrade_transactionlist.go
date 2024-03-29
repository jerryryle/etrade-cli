package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeTransactionList interface {
	GetAllTransactions() []ETradeTransaction
	GetTransactionById(transactionID string) ETradeTransaction
	NextPage() string
	AddPage(responseMap jsonmap.JsonMap) error
	AddPageFromResponse(response []byte) error
	AsJsonMap() jsonmap.JsonMap
}

type eTradeTransactionList struct {
	transactions []ETradeTransaction
	nextPage     string
}

const (
	// The AsJsonMap() map looks like this:
	// {
	//   "transactions": [
	//     {
	//       <transaction info>
	//     }
	//   ]
	// }

	// TransactionListTransactionsPath is the path to a slice of transactions.
	TransactionListTransactionsPath = ".transactions"
)

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

	// transactionsListTransactionsResponsePath is the path to a slice of transactions.
	transactionsListTransactionsResponsePath = ".transactionListResponse.transaction"

	// transactionsListMarkerResponsePath is the path to the next page number string
	transactionsListMarkerResponsePath = ".transactionListResponse.marker"
)

func CreateETradeTransactionListFromResponse(response []byte) (
	ETradeTransactionList, error,
) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeTransactionList(responseMap)
}

func CreateETradeTransactionList(responseMap jsonmap.JsonMap) (ETradeTransactionList, error) {
	// Create a new orderList with everything initialized to its zero value.
	transactionList := eTradeTransactionList{
		transactions: []ETradeTransaction{},
		nextPage:     "",
	}
	err := transactionList.AddPage(responseMap)
	if err != nil {
		return nil, err
	}
	return &transactionList, nil
}

func (e *eTradeTransactionList) GetAllTransactions() []ETradeTransaction {
	return e.transactions
}

func (e *eTradeTransactionList) GetTransactionById(transactionID string) ETradeTransaction {
	for _, transaction := range e.transactions {
		if transaction.GetId() == transactionID {
			return transaction
		}
	}
	return nil
}

func (e *eTradeTransactionList) NextPage() string {
	return e.nextPage
}

func (e *eTradeTransactionList) AddPageFromResponse(response []byte) error {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return err
	}
	return e.AddPage(responseMap)
}

func (e *eTradeTransactionList) AddPage(responseMap jsonmap.JsonMap) error {
	transactionsSlice, err := responseMap.GetSliceOfMapsAtPath(transactionsListTransactionsResponsePath)
	if err != nil {
		return err
	}

	// the marker key only appears if there are more pages, so ignore any
	// error and accept a possibly-zero int.
	nextPage, _ := responseMap.GetStringAtPath(transactionsListMarkerResponsePath)

	allTransactions := make([]ETradeTransaction, 0, len(transactionsSlice))
	for _, transactionJsonMap := range transactionsSlice {
		transaction, err := CreateETradeTransaction(transactionJsonMap)
		if err != nil {
			return err
		}
		allTransactions = append(allTransactions, transaction)
	}
	e.transactions = append(e.transactions, allTransactions...)
	e.nextPage = nextPage
	return nil
}

func (e *eTradeTransactionList) AsJsonMap() jsonmap.JsonMap {
	transactionSlice := make(jsonmap.JsonSlice, 0, len(e.transactions))
	for _, transaction := range e.transactions {
		transactionSlice = append(transactionSlice, transaction.AsJsonMap())
	}
	var transactionListMap = jsonmap.JsonMap{}
	err := transactionListMap.SetSliceAtPath(TransactionListTransactionsPath, transactionSlice)
	if err != nil {
		panic(err)
	}
	return transactionListMap
}
