package bank

import (
	"fmt"
	"github.com/appliedgocourses/bank"
	"os"
	"strconv"
	"strings"
)

const CREATE = "create"
const LIST = "list"
const UPDATE = "update"
const TRANSFER = "transfer"
const HISTORY = "history"

func Run() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	bank.Load()

	var cmd, acc1, acc2, amount, err = parse()

	if err != nil {
		fmt.Println(err)
		usage()
		return
	}

	switch cmd {

	case CREATE:
		if err := create(acc1); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("account %s create success\n", acc1)
		}

	case LIST:
		list()

	case UPDATE:
		if newAmount, err := update(acc1, amount); err != nil {
			fmt.Println(err)
			return
		} else {
			op := "withdrawal"
			if amount > 0 {
				op = "deposit"
			}
			fmt.Printf("account %s %s success new amount = %d\n", acc1, op, newAmount)
		}

	case TRANSFER:
		if fromAmount, toAmount, err := transfer(acc1, acc2, amount); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("account transferof %d from %s to %s success new from_amount = %d to_amount = %d\n", amount, acc1, acc2, fromAmount, toAmount)
		}

	case HISTORY:
		if err := history(acc1); err != nil {
			fmt.Println(err)
			return
		}

	}

	defer func() {
		if err := bank.Save(); err != nil {
			fmt.Printf("save panic: %s\n", err)
		}
	}()

}

func parse() (cmd string, name string, name2 string, amount int, err error) {
	cmd = strings.ToLower(os.Args[1])

	switch cmd {
	case CREATE:
		if len(os.Args) != 3 {
			err = fmt.Errorf("create wrong args")
			return cmd, "", "", 0, err
		}
		name = os.Args[2]

	case LIST:
		if len(os.Args) != 2 {
			err = fmt.Errorf("list wrong args")
			return cmd, "", "", 0, err
		}

	case UPDATE:
		if len(os.Args) != 4 {
			err = fmt.Errorf("update wrong args")
			return cmd, "", "", 0, err
		}
		name = os.Args[2]
		if amount, err = strconv.Atoi(os.Args[3]); err != nil {
			err = fmt.Errorf("update atoi amount: %s", err)
			return cmd, "", "", 0, err
		}

	case TRANSFER:
		if len(os.Args) != 5 {
			err = fmt.Errorf("transfer wrong args")
			return cmd, "", "", 0, err
		}
		name = os.Args[2]
		name2 = os.Args[3]
		if amount, err = strconv.Atoi(os.Args[4]); err != nil {
			err = fmt.Errorf("transfer atoi amount: %s", err)
			return cmd, "", "", 0, err
		}

	case HISTORY:
		if len(os.Args) != 3 {
			err = fmt.Errorf("history wrong args")
			return cmd, "", "", 0, err
		}
		name = os.Args[2]

	default:
		err = fmt.Errorf("bank unknown command %s", cmd)
		return cmd, "", "", 0, err
	}

	return cmd, name, name2, amount, err
}

func create(name string) error {
	a := bank.NewAccount(name)
	fmt.Printf("created account %s\n", bank.Name(a))
	return nil
}

func list() {
	dump := bank.ListAccounts()
	fmt.Printf("list of accounts\n%s\n", dump)
}

func history(name string) error {

	a, e := bank.GetAccount(name)

	if e != nil {
		return fmt.Errorf("history could not find account %s: %s", name, e)
	}

	f := bank.History(a)

	for ;; {
		amt, bal, more := f()
		op := "wth"
		if amt > 0 {
			op = "dps"
		}
		fmt.Printf("%s of %d			bal = %d\n", op, amt, bal)

		if more == false {
			return nil
		}
	}

	return nil
}

func update(name string, amount int) (newAmount int, err error) {

	if amount == 0 {
		return 0, fmt.Errorf("update amount is 0, nothing to do")
	}

	a, e := bank.GetAccount(name)

	if e != nil {
		return 0, fmt.Errorf("update could not find account %s: %s", name, e)
	}

	if amount > 0 {
		newAmount, e := bank.Deposit(a, amount)
		if e != nil {
			return 0, fmt.Errorf("update could not deposit %d into %s: %s", amount, name, e)
		}
		return newAmount, nil
	} else {
		withdrawAmount := amount * -1
		newAmount, e := bank.Withdraw(a, withdrawAmount)
		if e != nil {
			return 0, fmt.Errorf("update could not withdraw %d into %s: %s", withdrawAmount, name, e)
		}
		return newAmount, nil
	}
}

func transfer(name string, name2 string, amount int) (fromAmount int, toAmount int, err error) {

	if amount <= 0 {
		return 0, 0, nil
	}

	a1, e := bank.GetAccount(name)
	if e != nil {
		return 0, 0, fmt.Errorf("transfer could not find from_account %s: %s", name, e)
	}

	a2, e := bank.GetAccount(name2)
	if e != nil {
		return 0, 0, fmt.Errorf("transfer could not find to_account %s: %s", name2, e)
	}

	fromAmount, toAmount, e = bank.Transfer(a1, a2, amount)
	if e != nil {
		return 0, 0, fmt.Errorf("transfer failed from %s to %s for amount %d: %s", name, name2, amount, e)
	}
	return fromAmount, toAmount, nil

}

func usage() {
	fmt.Println(`Usage:
bank create <name>                     Create an account.
bank list                              List all accounts.
bank update <name> <amount>            Deposit or withdraw money.
bank transfer <name> <name> <amount>   Transfer money between two accounts.
bank history <name>                    Show an account's transaction history.
`)
}
