# go-reloaded
### overview
go-reloaded is a text editor, apply specific transformations to the text, and write the modified text to another file. The modifications include converting numbers between different bases, changing the case of words, adjusting punctuation, and modifying articles. This README provides instructions on how to use the tool and details the transformations it supports.

### usage
it expects two arguments: the input file and the output file.
+ `go run . inputFile outputFile`
inputFile contains the text need modifications
OutputFile contains the modified text

## example
Input (input.txt):
1E (hex) files were added.
It has been 10 (bin) years since the event.
Ready, set, go (up) !
I should stop SHOUTING (low)
Welcome to the Brooklyn bridge (cap)
I was sitting over there ,and then BAMM !!
I was thinking ... You were right
I am exactly how they describe me: ' awesome '
As Elton John said: ' I am the most well-known homosexual in the world '
There it was. A amazing rock!
```console 
./go-reloaded input.txt output.txt
```
Output (output.txt):
30 files were added.
It has been 2 years since the event.
Ready, set, GO!
I should stop shouting
Welcome to the Brooklyn Bridge
I was sitting over there, and then BAMM!!
I was thinking... You were right
I am exactly how they describe me: 'awesome'
As Elton John said: 'I am the most well-known homosexual in the world'
There it was. An amazing rock!
### Contributing

Feel free to submit issues and enhancement requests.
### License

This project is licensed under the MIT License - see the LICENSE file for details.
### Contact

For further information, you can reach out to the project maintainer at yhrouk@example.com.
