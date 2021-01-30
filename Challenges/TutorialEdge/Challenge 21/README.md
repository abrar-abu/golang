[Challenge 21 - JSON and Stock Dividends](https://tutorialedge.net/challenges/go/working-with-json/)

👋 Welcome Gophers! In this Go challenge, you are going to be working with JSON data and implementing a function that takes in a JSON string full of stock quotes and returns the one with the highest dividend return.

You will need to implement the HighestDividend function so that it takes in a JSON string. You will then need to somehow convert this JSON string into something you can traverse in Go using the encoding/json package.

Finally, you will need to return the string ticker for the stock that has the highest dividend.

# Example
```bash
[
  {"ticker": "APPL", "dividend": 0.5},
  {"ticker": "MSFT", "dividend": 0.2}
]

## HighestDividend returns "APPL"
```
# See the Solution
Feel free to have a look at the forum discussion thread for this challenge and contribute with your own solutions here 
- [Challenge 21 - Working with JSON](https://discuss.tutorialedge.net/t/challenge-21-json-and-stock-dividends/63)

# Further Reading:
If you enjoyed this challenge, you may also enjoy some of these other challenges:
- [Sorting Flights by Price](https://tutorialedge.net/challenges/go/sort-by-price/)
- [Satisfying Interfaces](https://tutorialedge.net/challenges/go/interfaces/)