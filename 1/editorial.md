# Day 1: Historian Hysteria

## Part 1

#### TDLR
Find the sum of differences of the $i$-th smallest element in the first array $A$ and the second array $B$ for all $1 \le i \le N$.

#### Solution
Sort the two arrays. This way, we can easily get the pairs of the $i$-th smallest elements. Simply iterate through $1$ to $N$ and add the differences.
Formally, let $A^s$ and $B^s$ be the sorted arrays of $A$ and $B$ respectively. Calculate $\sum^{N}_{i=1}{\left|A^s_i - B^s_i\right|}$.

## Part 2

#### TLDR
For each element in array $A$, find the number of its occurrence in array $B$ (i.e. how many times $A_i$ appears in $B$) and multiply it by $A_i$. Add them up for all $i$. In other words, the value of $$\sum^N_{i=1}{A_i \times cnt(A_i)}$$.

#### Solution
The only challenge is to quickly count the occurrence of each element. For this, we can use a hashmap $M$ such that $M[x] = cnt(x)$ how many times x appears in B. Then do the usual iterate, multiply, and add.
