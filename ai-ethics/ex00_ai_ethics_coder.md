# *Exercise 00: Learning to use AI ethically as a coder*
---
# Part A: Self-assessment

**How have you used AI for coding so far?**

I have use AI basically a research engine. I also use AI for checking and validating my code, I use AI to break down codes for clarity and deeper understand of the logic (why) behind every syntax. I sometime also ask AI for assistance when stuck, especially dealling with task I have no basic knowledge of it concepts.

**Do you ask AI for solutions before trying yourself?**

No, I usually try out solutions myself before asking AI

**Can you explain code you've submitted without AI's help?**

Yes, I can explain any code i submitted without AI's help.

**What would happen if AI was suddenly unavailable during an exam or interview?**

They is no cause for alarm, because going for an exam or interview i go prepared not with the intention of using AI.

## 2. Identify your current pattern: Which learner are you now?

> Learner B: "AI is my learning amplifier"

## 3. Write a brief paragraph: Where are you now, and where do you want to be?

- I'm currently use AI as a guide, support, teacher for validations, research, concept breakdown and explanations.
- I want to become a learner that use AI ethically and more effective for learning and understanding concepts in depth.

# Part B: Palindrome solution + reflection 
Challenge: Implement a simple function to check if a string is a palindrome.

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
```

### 3. Test with examples like "racecar", "hello", and "A man a plan a canal Panama".

**Test code**
print(checkPalindrome("racecar")) 
print(checkPalindrome("hello")) 
print(checkPalindrome("A man a plan a canal Panama"))

***Output:***
```graphml

True
False
False

```

##  Step 2: Strategic AI use After you have a working solution, ask AI:

**1. Time complexity**

O(n) time — one full pass to build the reversed string + one pass to compare

O(n) extra space — slicing creates a second copy of the string

**2. Missing edge cases**

- Case sensitivity 

- Spaces

- Punctuation

- Unicode normalization, Non-string inputs 

- Definition decision: should empty string count as palindrome? (usually True)


**3. Alternatives & trade-offs**

- Two-pointer approach

- Generator comparison

- Recursive 

**4. Performance on very long strings**

- my version is memory-bound (needs ~2× string size in RAM)

- Can slow down or crash for huge inputs due to allocation/copying

- Two-pointer version remains stable because it uses constant memory

## Step 3: Reflection

**What did you learn by struggling first?**

What I understood before help

I focused only on the direct definition:
A palindrome reads the same forward and backward.

I implemented the most straightforward comparison.

**How is your understanding different than if you'd just asked for the solution?**

What changed after thinking deeper

I now understand the difference between:

solving the problem

solving real-world input

Real input is messy (spaces, symbols, case differences).

**Can you now implement similar functions (reverse a string, find duplicates) without AI?**

Yes I can now write similar functions without help.

**What mental model did you build?**

I learned to stop just trying to get the correct answer, and start thinking about how much work the computer is doing to get that answer.

Before, I would use the easiest trick that works.
Now I think:

- Do I really need to check everything?

- Can I stop early if I already know it’s wrong?

I moved from getting the answer to getting the answer efficiently.


# Part C: Testing Your Understanding

Without using AI, complete these variations:

Ignore spaces and punctuation

Make it case-insensitive

Return the position where the string stops being a palindrome (if not a palindrome)

```py
def check_palindrome(s):
    n_s = []

    #normalize string
    for char in s:
            if char.isalnum():
                n_s.append(char.lower())

    left = 0
    right = len(n_s) -1

    #compare two pointer
    while left < right:
        if n_s(left) != n_s(right):
            return (False, left, right)
        left += 1
        right -= 1

    return True
```
# Part D: Personal fairness contract 

**I will use AI when:**

- After I've attempted a problem for at least 20 minutes

- To understand why my solution works/doesn't

- To explore alternatives after I have a working solution

- To understand advance concepts

**I will NOT use AI when:**

- I haven't tried the problem myself

- I'm taking an assessment or test

- I need to build fundamentals

- To teach others

**I know I'm using AI fairly when:**

- I can explain my code without looking at AI's response

- I could solve similar problems without AI

- I feel more confident in my abilities

- I could teach others 

Victor Ejeh
20/02/2026

# Part E: Real-World Scenario Analysis

**Interview: "Explain how you'd implement a caching system." If you always relied on AI, can you answer?**
If i have always relied on AI to implement a caching system, i will not be able to explain how it was done because i didn't really do it myself.

**Production bug at 2 AM: AI is unavailable. Can you debug code you don't fully understand?**
No i wont be able to debug the code because didn't fully understand it.

**New tech with little documentation: If you never learned to read docs and experiment, what happens?**
I will definatly get to learn how to read documention.


Using AI fairly will help me Better implement and understand my code, Improve my proplem solving skills. I will be irreplaceable.

# Part F: Building Irreplaceable Skills

# Irreplaceable Skills

Rate yourself from **1 (very weak)** to **5 (strong)** and fill in the improvement plan.

| Skill                    | Description                               | Rating (/5) | Improvement Plan                                                                            |
| ------------------------ | ----------------------------------------- | ----------- | ------------------------------------------------------------------------------------------- |
| Problem Decomposition    | Breaking down problems logically          | 3           | Before coding, write step‑by‑step algorithm in plain English and convert it into functions  |
| Systems Thinking         | Understanding how components interact     | 2           | Draw diagrams of program flow (input → processing → output → storage) before implementation |
| Critical Evaluation      | Knowing when code is wrong or inefficient | 2           | After solving, rewrite solution a second time aiming for fewer conditions and clearer logic |
| Debugging Mindset        | Investigating unexpected behavior         | 3           | Use print/log tracing to follow execution instead of immediately rewriting code             |
| Conceptual Understanding | Knowing WHY, not just HOW                 | 3           | Read documentation and explain concepts aloud without looking at code                       |

---

## Lowest‑Rated Skill

**Chosen Skill:** Systems Thinking

### 3 Actions I Will Take This Week (without AI)

1. For every exercise, draw a diagram showing data flow between functions before coding.
2. Build a small CLI tool that reads a file, processes it, and outputs a result, documenting each component’s responsibility.
3. Manually trace program execution using logs and write a written explanation of how data moves through the system.

