main = print $ maximum
    [ product 
    | x <- [100..999]
    , y <- [100..999]
    , let product = x * y
    , isPalindrome product
    ]

isPalindrome :: Integer -> Bool
isPalindrome x = str == reversed
  where
    str = show x
    reversed = reverse str
