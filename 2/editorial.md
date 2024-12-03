# Day 2: Red-Nosed Reports

## Part 1

You are given $N$ reports (one per line). Each report $R_i$ has $L_i$ levels which are represented as integers. A report is found safe if:
* The levels are either all increasing or decreasing.
* Any two adjacent levels differ by **at least** 1 and **at most** 3. 

- `7 6 4 2 1` is **safe** because it satisfies both criteria.
- `1 2 7 8 9` is **unsafe** because $7 - 2 = 5$ which is greater than $3$. 
- `1 3 2 4 5` is **unsafe** because the array is neither all increasing nor all decreasing. 

We have to count how many reports are safe.

### Solution

Honestly, I feel the solution is quite straightforward - simply verify whether the two criteria are met. 

Let's consider report $R$. 

For the first criterion, we can check each adjacent elements, namely $R_{j}$ and $R_{j+1}$. If at some $j$, $R_{j} < R_{j+1}$ and at another $j$, $R_{j} > R_{j+1}$, then this criterion is violated.

For the second criterion, we calculate the difference between the two adjacent levels. If $\left|R_j - R_{j+1}\right| < 1$ or $\left|R_j - R_{j+1}\right| > 3$, then this criterion is violated.

If the two criteria is not violated, we can let this report be safe.

## Part 2

Similar to **Part 1** but this time, they are more lenient in considering a report as safe. If by removing a level from the report it becomes a safe report (as defined in Part 1), it is considered safe.

- `1 3 2 4 5` is now **safe** because you can remove the second element (3) and the report becomes `1 2 4 5` which is safe.

### Hint
Which element to delete?

### Solution
Let's define the function to check if a report is safe in Part 1 as `CheckSafe()`. For Part 2, just do what the problem tells you to. Try all possibilities of removing an element (i.e. try removing the first and do `CheckSafe()`, then try removing the second and so on). 

