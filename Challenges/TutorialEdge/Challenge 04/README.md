[Challenge 04 - Checking for Duplicates](https://tutorialedge.net/challenges/go/checking-for-duplicates/)

In this challenge, we are going to be looking at how you can effectively filter out the duplicate entries from a slice in Go.

The task will be to implement the FilterDuplicates function so that it returns a slice of type string which contains only the unique names of developers retrieved from the inputted list.

# Example:

```bash
// input
[]Developers{
  Developer{Name: "Elliot"},
  Developer{Name: "Alan"},
  Developer{Name: "Jennifer"},
  Developer{Name: "Graham"},
  Developer{Name: "Paul"},
  Developer{Name: "Alan"},
}

// output
[]string{
  "Elliot",
  "Alan",
  "Jennifer",
  "Graham",
  "Paul",
}

```

> Hints
You may wish to use a map in your function in order to check if elements have already been seen by our function.

# See the Solution
Feel free to have a look at the forum discussion thread for this challenge and contribute with your own solutions here 
- [Challenge 04 - Filtering Duplicates](https://discuss.tutorialedge.net/t/challenge-04-filtering-duplicates/21)

# Further Reading:
If you like this challenge then you may also enjoy some of the other challenges on the site:
- [Sorting Flights by Price](https://tutorialedge.net/challenges/go/sort-by-price/)
- [Satisfying Interfaces](https://tutorialedge.net/challenges/go/interfaces/)
