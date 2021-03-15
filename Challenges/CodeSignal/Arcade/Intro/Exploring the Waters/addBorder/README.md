`Easy`	`Codewriting` 	`300`

Given a rectangular matrix of characters, add a border of asterisks(*) to it.

## Example

- For

``` bash
picture = ["abc",
           "ded"]
```

the output should be

``` bash
addBorder(picture) = ["*****",
                      "*abc*",
                      "*ded*",
                      "*****"]
```

## Input/Output

- [execution time limit] 4 seconds (go)

- [input] array.string picture

    A non-empty array of non-empty equal-length strings.

    Guaranteed constraints:
    `1 ≤ picture.length ≤ 100`,
    `1 ≤ picture[i].length ≤ 100`.

- [output] array.string

    The same matrix of characters, framed with a border of asterisks of width 1.

## [Go] Syntax Tips

``` go
// Prints help message to the console
// Returns a string
func helloWorld(name string) string {
    fmt.Printf("This prints to the console when you Run Tests");
    return "Hello, " + name;
}
```
