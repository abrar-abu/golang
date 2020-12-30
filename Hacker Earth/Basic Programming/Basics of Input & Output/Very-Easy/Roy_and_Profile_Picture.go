/*
Roy wants to change his profile picture on Facebook. Now Facebook has some restriction over the dimension of picture that we can upload.
Minimum dimension of the picture can be L x L, where L is the length of the side of square.

Now Roy has N photos of various dimensions.
Dimension of a photo is denoted as W x H
where W - width of the photo and H - Height of the photo

When any photo is uploaded following events may occur:

[1] If any of the width or height is less than L, user is prompted to upload another one. Print "UPLOAD ANOTHER" in this case.
[2] If width and height, both are large enough and
(a) if the photo is already square then it is accepted. Print "ACCEPTED" in this case.
(b) else user is prompted to crop it. Print "CROP IT" in this case.

(quotes are only for clarification)

Given L, N, W and H as input, print appropriate text as output.

Input:
First line contains L.
Second line contains N, number of photos.
Following N lines each contains two space separated integers W and H.

Output:
Print appropriate text for each photo in a new line.

Constraints:
1 <= L,W,H <= 10000
1 <= N <= 1000

SAMPLE INPUT 	
180
3
640 480
120 300
180 180

SAMPLE OUTPUT 	
CROP IT
UPLOAD ANOTHER
ACCEPTED

*/

package main

import (
	"fmt"
)

func main() {
	var Size, Sample int
	fmt.Scan(&Size)
	fmt.Scan(&Sample)
	array := make([][2]int, Sample)
	for i := 0; i < Sample; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&array[i][j])
		}
	}
	for i := 0; i < Sample; i++ {
	    if (array[i][0] < Size || array[i][1] < Size){
	        fmt.Println("UPLOAD ANOTHER")
	    } else {
	        if( array[i][0] == array[i][1] && (array[i][0]>=Size  && array[i][1]>=Size)){
					fmt.Println("ACCEPTED")
	        } else {
				fmt.Println("CROP IT")	            
	        }
	        
	    }
	}
}
