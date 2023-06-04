main = print $ sum $ filter (even) $ calculateFibList (<= 4_000_000) 0 1 []


calculateFibList cond prev2 prev1 list
    | cond cur = cur : calculateFibList cond prev1 cur list
    | otherwise = list
    where cur = prev2 + prev1