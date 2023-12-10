
### Problem Statement

You are in charge of the cake for a child's birthday. You have decided the cake will have one candle for each year of their total age. They will only be able to blow out the tallest of the candles. Count how many candles are tallest.

Write a function `birthdayCakeCandles(candles)` that takes in a list of candle heights and returns the number of candles that are tallest.

#### Function Description

The function `birthdayCakeCandles` has the following parameter:

- `candles` : an integer array representing the heights of candles

#### Returns

- an integer representing the number of candles that are tallest

#### Input Format

The first line contains a single integer, `n`, the size of `candles`.
The second line contains `n` space-separated integers, where each integer describes the height of a candle.

#### Constraints

- \( 1 \leq n \leq 10^5 \)
- \( 1 \leq \text{candles}[i] \leq 10^7 \)

#### Sample Input

```
4
3 2 1 3
```

#### Sample Output

```
2
```

#### Explanation

Candle heights are `[3, 2, 1, 3]`. The tallest candles are `3` units, and there are `2` of them.
