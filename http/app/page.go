package app

const pageHTML = `
<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>
<p>My first paragraph.</p>

  <form>
  <label for="amount">Amount</label>
	<label for="desc">Description</label>
    <input type="text" id="desc" name="desc"><br>

	<label for="date">Date</label>
    <input type="text" id="date" name="date"><br>

	<label for="category">Category (dropdown)</label>
    <input type="dropdown" id="category" name="category"><br>

	<label for="type">Transaction Type (radio: income, expense, transfer, investment)</label>
    <input type="radio" id="type" name="type"><br>

	<label for="is_debit">IsDebit (checkbox)</label>
    <input type="checkbox" id="is_debit" name="is_debit"><br>

	<label for="account">Account (dropdown)</label>
    <input type="dropdown" id="account" name="account"><br>
  </form>

</body>
</html>
`
