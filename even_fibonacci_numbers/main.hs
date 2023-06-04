main = print $ sum $ takeWhile (< 4_000_000) $ filter (even) $ map (fib) [1,2..]

fib 1 = 1
fib 2 = 2
fib x = fib (x - 1) + fib (x - 2)