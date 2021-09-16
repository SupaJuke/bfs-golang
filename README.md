# bfs-golang
This is my personal project of developing a bfs alrogithm using golang. The purpose of this project is for me to self-teach myself a new language (Golang) while also not be rusty in theory—the bfs algorithm.

## Things to do
1. <s>Have the program be able to execute on a graph represented by an adjacent list in a text format (.txt)</s> DONE!

## Text format
The first line (the header) must consist of only the value of each node (separated by a space), e.g. "A B C D E" or "Hello World I Love Myself"
Both of the examples would result in 5 nodes, one being "A", "B", "C", "D", and "E" while the other one being "Hello", "World", "I", "Love", and "Myself".

The following line will be in the following format: "[node]:[other outnodes, separated by a space]"

For example, "A:B C" and "World:I Love Hello" are valid while "A : B C" is not.
