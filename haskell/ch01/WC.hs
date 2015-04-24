-- file: ch01/WC.hs

main = interact wordCount
--	where wordCount input = show (length (lines input)) ++ "\n"
	where wordCount input = show (sum [ length x | x <- (lines input) ]) ++ "\n"
