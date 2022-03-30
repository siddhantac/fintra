An expense manager application with zero-based budgeting.

**Under construction**

# Roadmap

- [x] create transactions
- [x] list transactions, get transaction by id
- [x] account and category validation (simple hashmap based validation)
- [x] data persistence using boltDB
- [ ] (WIP) home page UI
- [ ] transfers - (2 transactions in DB)
- [ ] report
    - [ ] filter transactions by category
    - [ ] filter transactions by account
    - [ ] filter transactions by date range
    - [ ] account balance
        - [ ] requires starting balance
        - [ ] current balance maybe cached in boltDB (using db transactions)

---

## Brainstorming
* add transactions
* add budget per category
	* should be able to calculate remaining budget available
* move money between categories
* "savings" is a special category, it would be the bank balance as well as a category
* add `account` field later

## Phase 0
- [x] bank account
    - [x] starting balance
    - [x] current balance
- [ ] 2 categories + savings
- [x] overview of the whole system (status report)
    - [x]show a table with all transactions
* [x] in-memory storage for transactions
- [x] API methods:
	- [x] create new transaction
	- [x] get all transactions

## Phase 1
- [x] setup proper architecture
    - [x] service layer
    - [x] http layer (WIP)
    - [x] wire everything together in main
- [x] http server
-   - [x] add go-chi
-   - [x] add logs
- [ ] repository layer
    - [x] transactions repo
      - [x] GET 
      - [x] Create
      - [ ] Delete
    - [x] account repo (cols: txn id, txn amt, running balance)
      - [x] GET 
      - [x] Create
      - [ ] Delete
    - [ ] categories repo (required by "assign budget per category")
      - [ ] GET methods. (POST method to create can be done later)

* repository/ -> accountRepo, transactionsRepo, categoriesRepo

_NOTE: consider using this for auto-rebuild https://github.com/cosmtrek/air_

## Phase 2 - simple transactions
* persistent storage
    * file based storage (json)
    * sqlite?
* API methods to  return accounts and categories
* API methods to add new ones (maybe)
* account and category validation when creating transaction
* system overview: see list of transactions
* simple report: expenditure by category and by account, account balance

## Phase 3 - zero based budget
* API methods:
	* set starting balance
	* assign budget per category 
	* re-assign category budget by moving money around
* system status should show balance based on starting balance & all transactions since then

