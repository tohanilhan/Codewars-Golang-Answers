# Explosive Sum
How many ways can you make the sum of a number?
From wikipedia: https://en.wikipedia.org/wiki/Partition_(number_theory)

> In number theory and combinatorics, a partition of a positive integer n, also called an integer partition, is a way of writing n as a sum of positive integers. Two sums that differ only in the order of their summands are considered the same partition. If order matters, the sum becomes a composition. For example, 4 can be partitioned in five distinct ways:

```
4
3 + 1
2 + 2
2 + 1 + 1
1 + 1 + 1 + 1
```
## Examples

### Basic
```
ExpSum(1) // 1
ExpSum(2) // 2 -> 1+1 , 2
ExpSum(3) // 3 -> 1+1+1, 1+2, 3
ExpSum(4) // 5 -> 1+1+1+1, 1+1+2, 1+3, 2+2, 4
ExpSum(5) // 7 -> 1+1+1+1+1, 1+1+1+2, 1+1+3, 1+2+2, 1+4, 5, 2+3

ExpSum(10) // 42
```
### Explosive
```
ExpSum(50) // 204226
ExpSum(80) // 15796476
ExpSum(100) // 190569292
```
See [here](http://www.numericana.com/data/partition.htm) for more examples.