[Challenge 08 - Check Permutations](https://tutorialedge.net/challenges/go/check-permutations/)

👋 Welcome Gophers! In this Go challenge, you are going to be implementing a function that takes in two string values and checks to see if they are permutations of one another.

# Example
If I have 2 strings, "abc" and "cba", when I pass these strings into the function, it should return true as these two strings are permutations of each other.

> Hints
You can iterate through all the characters in a string using the range keyword in a for loop

```bash
// prints out the position and the rune
for pos, char := range str1 {
  fmt.Printf("%d: %c\n", pos, char)
}
```
Start off by building up a map of these rune values to the number of occurrences in one for loop and then work from there.

## Question
How can we optimize this function so that it is not performing unnecessary calculations?

```bash
A
We can check the length of each string at the start of the function and return false if they differ
```

```bash
B
We can use maps to efficiently lookup perviously encountered characters
```

```bash
C
All of the Above
```

> Correct Answer: We can implement all of the above checks to ensure that the function only does what it has to before returning the correct answer

# See the Solution
Feel free to have a look at the forum discussion thread for this challenge and contribute with your own solutions here 
- [Challenge 08 - Checking Permutations](https://discuss.tutorialedge.net/t/challenge-08-checking-permutations/25/2)


# Further Reading:
If you enjoyed this challenge, you may also enjoy some of the other challenges on this site:
- [Sorting By Price](https://tutorialedge.net/challenges/go/sort-by-price/)
- [Satisfying Interfaces](https://tutorialedge.net/challenges/go/interfaces/)