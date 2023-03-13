# Prime Generator

**Input**: Given n

**Output**: Print all primis smaller or equal than n [otherwise n-1]

## Resolution
### 1. Brute force:  Naive Approach

Iterate from `2` to `n - 1` (or `n` in this case) and check each number for primeness. If any number in this range divides n, then it is not a prime number.

**Cycle Optimization**: iterate to `n/2`

**Time Complexity**: O(N)

**Auxiliary Space**: O(1)

-

**Naive approach (recursive)**: Recursion can also be used to check if a number between `2` to `n – 1` divides n. If we find any number that divides, we return false.

**Time Complexity**: O(N)

**Auxiliary Space**: O(N) if we consider the recursion stack. Otherwise, it is O(1).

### 2. Efficient Approach

Iterate through all numbers from 2 to `sqrt(n)` and for every number check if it divides n.

**Time Complexity**: O(sqrt(n))

**Auxiliary space**: O(1)


### 3. Better way: Sieve of Eratosthenes - Simple Sieve

**Cons:** When n is large

* An array size of Theta(n) may not fit in memory
* It is not cached friendly even for slighty bigger n
* The algorithm traverse the array without locality of reference

**Time Complexity**: O(n * log(log(n)))

**Auxiliary space**: O(n)

### 4. More efficiency: Segmented Sieve

The idea is divide the range `[0..n-1]` (`[0..n]`) in different segments and compute primes in all segments one by one.

=> In this case, the algorithm uses Simple Sieve to find primes smaller than or equal to sqrt(n)

**Pass:**

1.  `prime_root[]` : get output from `simpleSieve(sqrt(n))`
2. We need all primes in the range `[0..n-1]` (`[0..n]`). We divide this range into different segments such that the size of every segment is at most `sqrt(n)`
3. For every segment `[low_bound...high_bound]`
	* Create an array `mark[high-low+1]`. Here we need only O(x) space where x is a number of elements in a given range.
	* Iterate through all primes found in step 1, for every prime, mark its multiple in the given range `[low_bound...high_bound]`


In Simple Sieve, we needed O(n) space which may not be feasible for large n. Here we need O(√n) space and we process smaller ranges at a time (locality of reference)

**Time Complexity**: O(n * ln(sqrt(n)))
**Ausiliary Space**: O(sqrt(n))


Note that time complexity (or a number of operations) by Segmented Sieve is the same as Simple Sieve. It has advantages for large ‘n’ as it has better locality of reference thus allowing better caching by the CPU and also requires less memory space.