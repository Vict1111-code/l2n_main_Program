# Step 1 - Do it yourself

### 1. Write pseudocode for a function that checks if a string is a palindrome.

** Palindrome logic(pseudocode) **
```

function checkPalindrome(string)


	normalized string (lowercase, Ignore space and Punctuation)

	reverse string equal to input or test string[::-1]

	IF reverse_string equal input or test string:
		RETURN TRUE
	ELSE:
		RETURN FALSE	
end

```

### 2. Python Implementation.

Before AI Assistance 

```py
def checkPalindrome(s: str):
	#Reverse the input string
	reverse_s = s[::-1]
	
	#compare the input string and reverse string
	if reverse_s == s: #checks if the reverse string is equal to the original string
		return True #Is a palindrome 
	else:
		return False #Not a palindrome
#Test code
print(checkPalindrome("racecar")) 
print(checkPalindrome("hello")) 
print(checkPalindrome("A man a plan a canal Panama"))
```

### 3. Test with examples like "racecar", "hello", and "A man a plan a canal Panama".

#### Output:
```graphml

True
False
False

```



# Step 2 - Use AI to learn


## 1. Time Complexity

> AI explained that reversing takes O(n) and comparing also takes O(n). So overall it's linear time. It also creates another string, so memory is O(n). I didn’t think about memory before that.

## 2. Edge Cases

| Case           | Example                         | Problem                |
| -------------- | ------------------------------- | ---------------------- |
| Capitalization | `"RaceCar"`                     | case mismatch           |
| Spaces         | `"A man a plan a canal Panama"` | extra characters        |
| Punctuation    | `"madam, I'm Adam"`             | symbols break equality |
| Empty string   | `""`                            | works accidentally     |
| Single char    | `"a"`                           | works                  |
| Unicode        | `"Àbba"`                        | lowercase alone not enough   |

## 3. Better Approach:
It is more efficient to use two pointer techniques instead of reverse the whole string at once it helps reduce memory usage and Complexity

---

# Step 3 - Reflection

## What did you learn from solving it before asking AI?

What I understood before help

I focused only on the direct definition:
A palindrome reads the same forward and backward.

I implemented the most straightforward comparison.

## How is your understanding different now?

What changed after thinking deeper

I now understand the difference between:

solving the problem

solving real-world input

Real input is messy (spaces, symbols, case differences).

## Could you now write similar functions (e.g., reverse a string) without help?

Yes I can now write similar functions without help.