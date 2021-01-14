[Challenge 05 - Implementing a Stack in Go](https://tutorialedge.net/challenges/go/implementing-a-stack/)

👋 Welcome Gophers! In this challenge, we are going to be implementing some of the basic functionality of the Stack data structure in Go.

This is going to be the first of a number of data-structure questions which may come in handy if you are about to go in for an interview!

We`ll be carrying on the theme of flying from the previous challenge here and implementing 3 crucial methods needed to support a basic implementation of a stack.


# Push
The first challenge will be to implement the Push function of our Stack interface.

This method will take in a Flight and push the flight onto the top of our Items stack.

# Peek
The second part of this challenge will be implementing the Peek function.

This method will allow us to view what item is at the top of our stack but not modify the underlying stack values.

# Pop
The third and final part of this challenge will be implementing the Pop function.

This method will allow us to pop an element off the top of our Items stack and return to us the top flight.

## Question
Pointer receivers, denoted by the (s *Stack), on our methods allow us to modify the value to which the receiver points

```sh 
A
True
```
```bash
B
False
```
> Correct Answer: True - not using pointer receivers here would mean that calling pop and push would not update the underlying value of the stack

# See the Solution
Feel free to have a look at the forum discussion thread for this challenge and contribute with your own solutions here 
- [Challenge 05 - Implementing a Stack](https://discuss.tutorialedge.net/t/challenge-05-implementing-a-stack/22)


# Further Reading:
If you enjoyed this challenge, you may also enjoy some of the other challenges on this site:
- [Sorting By Price](https://tutorialedge.net/challenges/go/sort-by-price/)
