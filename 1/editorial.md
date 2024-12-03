# Day 1: Historian Hysteria

## Part 1

#### TDLR
Find the sum of differences of the i-th smallest element in the first array and the second array for all 1 <= i <= n.

#### Solution
Sort the two arrays. This way, we can easily get the pairs of the i-th smallest elements. Simply iterate through 1 to n and add the differences.

## Part 2

#### TLDR
For each element in array A (A[i]), find the number of its occurrence in array B (i.e. how many times A[i] appears in B) and multiply it by A[i]. Add them up for all i.

#### Solution
The only challenge is to quickly count the occurrence of each element. For this, we can use a hashmap M such that M[x] = how many times x appears in B. Then do the usual iterate, multiply, and add.