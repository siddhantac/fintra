## Brainstorming
* add transactions
* add budget per category
	* should be able to calculate remaining budget available
* move money between categories
* "savings" is a special category, it would be the bank balance as well as a category
* add `account` field later

## Phase 0
- [ ] bank account
    - [ ] starting balance
    - [ ] current balance
- [ ] 2 categories + savings
- [ ] overview of the whole system (status report)
    - [x]show a table with all transactions
    - [ ]AND, final balance of the account
- [ ] in-memory storage
- [ ] API methods:
	- [x] create new transaction
	- [x] get all transactions

## Phase 1
* file based storage (json)
* API methods:
	* set starting balance
	* assign budget per category
	* re-assign category budget by moving money around
* system status should show balance based on starting balance & all transactions since then