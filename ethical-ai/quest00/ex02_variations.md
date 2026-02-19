## 1. Modified palindrome function to:
*Ignore space and punctuation.
*Be case-insensitive.
*Return the position where the string stops being a palindrome(if not one)

---

### First attempt Before AI assistance

```py
def checkPalindrome(s: str):
    #normalize string
    normalized_chars = [] 
    
    #keep only letters and numbers
    for char in s: 
        if char.isalnum():
            normalized_chars += char.lower()

    left = 0
    right = len(normalized_chars) - 1
    
    #compare left and right
    while left < right:
        if normalized_chars[left] != normalized_chars[right]:
            
            # Return the index where mismatch occurs (from normalized string perspective)
            mismatch = (f"Not a palindrome, mismatch at Index: {left} and {right}")
            return mismatch # returns a message and the point of mismatch (left, right)
    
        left += 1
        right -= 1
        
    # Step 3: No mismatch found(is a palindrome)
    return ("Is a palindrome")

# Test cases
print(checkPalindrome("racecar"))  # (is a palindrome)
print(checkPalindrome("hello"))    # (Not a palindrome, mismatch at: the left and the right index)
print(checkPalindrome("A man a plan a canal Panama"))  #(is a palindrome)

```
---

## 2. After My first attempt, I ask AI:
> "I modified my palindrome function to handle more cases.
Did I miss anything? Can it be more efficient?"


AI pointed out my list building was inefficient because I used += which loops internally. I changed it to append. It also suggested returning data instead of sentences which makes sense for real programs.

## Updated code after AI review:

```py

def checkPalindrome(s: str):
    #normalize string
    normalized_chars = [] 
    
    #keep only letters, numbers and change uppercase to lowercase
    for char in s: 
        if char.isalnum():
            normalized_chars.append(char.lower())

    left = 0
    right = len(normalized_chars) - 1
    
    #compare left and right
    while left < right:
        if normalized_chars[left] != normalized_chars[right]:
            
        
            return (False, left, right) # returns false and the mismatch point
    
        left += 1
        right -= 1
        
    # Step 3: No mismatch found(Return True)
    return True

# Test cases
print(checkPalindrome("racecar"))  # True(is a palindrome)
print(checkPalindrome("hello"))    # False(the left and the right index)
print(checkPalindrome("A man a plan a canal Panama"))  #True(is a palindrome)

```

---

## 3. Reflection on what AI added that I didn't consider initially.

### What I corrected after review

**inefficient List construction**
Using += on a list repeatedly causes repeated internal copying.
The proper way is using append()

** 2. Return Type Design:**
I learn't that returning sentences is not a good practice for reusable code. 
A functions should return data, not explanations. 
Better Options:
- True / False
- (False, left, right) for mismatch
- None for success