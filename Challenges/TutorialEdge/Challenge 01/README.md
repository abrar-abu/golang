[Challenge 01 - Type Assertions in Go](https://tutorialedge.net/challenges/go/type-assertions/)

In this challenge, we are going to become familiar with the concept of Type Assertions in Go!

# Preface
If you are new to the language, then type assertions are a concept that can sometimes trip you up and appear a little tricky at first, but after overcoming the syntax it becomes far easier to understand.

Through using type assertions, we can retrieve the dynamic value of an interface. For example:

```bash
var myName interface{} = "Elliot"

name := myName.(string)
fmt.Println(name)
```

In this example, we have an interface which has a dynamic value of “Elliot”. We can then use a type assertion to retrieve this dynamic value and use the value just like we would any other string value in Go.

# Challenge
In this challenge, we are going to define a function that is called GetDeveloper which will take in 2 interface{} arguments.

Within this function, you will have to declare a new Developer instance and use type assertion to populate the values correctly before then returning this new Developer instance.

# See the Solution
Feel free to have a look at the forum discussion thread for this challenge and contribute with your own solutions here - [Challenge 01 - Type Assertions](https://discuss.tutorialedge.net/t/challenge-01-type-assertions/18/2)

Further Reading \
[Go Interfaces Tutorial](https://tutorialedge.net/golang/go-interfaces-tutorial/)