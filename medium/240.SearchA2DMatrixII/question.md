[Search a 2D Matrix II](https://leetcode.com/problems/search-a-2d-matrix-ii/)

# Description
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

* Integers in each row are sorted in ascending from left to right.
* Integers in each column are sorted in ascending from top to bottom.

## Examples
Consider the following matrix:

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

Given target = 5, return true.

Given target = 20, return false.

## Tip
We can use divide-and-conquer algorithm to solve the problem. Illustrate the strategy in the following picture.
![divide-and-conquer illustration](./search_2d_matrix.png)
