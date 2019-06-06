# KMP Algorithm

## Description
[KMP](https://www.cnblogs.com/yjiyjige/p/3263858.html) is a string matching algorithm, it's time complexity is O(m+n). 
The key of KMP algorithm is to use the information after matching failure to minimize the number of matches between pattern string and main string in order to achieve fast matching. The concrete implementation is to implement a next () function, which contains the local matching information of pattern string.

Give one string, get the pattern array(int[]) which stores backtracking locations

## Example
* Example1
Input: "ab"
Output: [-1,0]

* Example2
Input: "aa"
Output: [-1,-1]

* Example3
Input: "abcd"
Output: [-1,0,0,0]

* Example4
Input: "abcabcd"
Output: [-1,0,0,-1,0,0,3]