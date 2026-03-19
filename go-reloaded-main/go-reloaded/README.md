# go-reloaded

🗺 go-reloaded Program Flow (Word-Level Approach)

```
Start
  |
  v
Read input file into a string
  |
  v
Split string into words (word-level) → words[]
  |
  v
FOR each word in words[]:
  |
  +--> Is word an instruction like (up), (low), (cap)?
  |       |
  |       +--> Yes: 
  |       |     Modify previous 1 or N words according to instruction
  |       |     Remove/clean instruction from words[]
  |       |     Continue to next word
  |       |
  |       +--> No: move to next check
  |
  +--> Is word a = ? 
  |       |
  |       +--> Yes:
  |       |     Look at next word
  |       |     If next word starts with vowel/h → change "a" to "an"
  |       |     Continue
  |       |
  |       +--> No: continue
  |
  +--> Is word a number followed by (hex) or (bin)?
          |
          +--> Yes:
          |     Convert number to decimal
          |     Remove (hex)/(bin)
          |     Continue
          |
          +--> No: move to next word
  |
  v
After loop, join words[] back into a string
  |
  v
Run second pass to fix punctuation and quotes:
  - Remove space before ., ,, !, ?, :, ;
  - Add one space after them
  - Handle ... and !? properly
  - Fix quotes 'word' spacing
  |
  v
Write the transformed string to output file
  |
  v
End

```


🔑 Key Points in This Flowchart

Two passes:

First pass = word-level transformations (instructions, grammar, numbers)

Second pass = character-level cleanup (punctuation, quotes)

Word-level first:

Makes backward modifications easy

Keeps logic readable

Controlled index access:

Always check previous and next exist

Prevent “index out of range” errors

Remove instruction words after using them:

Keeps the array clean for further processing

Hybrid approach:

Word-level for logic

Character-level for formatting

🏁 Beginner-Friendly Mental Model

Think of your program like a teacher reading a sentence word by word:

Sees a number → maybe convert it

Sees (up) → shout previous word(s) in uppercase

Sees "a" → checks if it should become "an"

Ignores normal words → moves on

After finishing sentence → tidy up punctuation


Refined Pseudocode for go-reloaded (Word-Level Approach)

```
START

1️⃣ Read the input file into a string called "text"

2️⃣ Split "text" into words array:
   words = split text by spaces

3️⃣ FOR each index i in words:
     currentWord = words[i]

     // Check for transformations
     IF currentWord contains "(up)" or "(up, N)":
         - Find number of words to modify (default 1 if no N)
         - Modify previous N words to UPPERCASE
         - Remove/clean instruction token(s) from words[]

     ELSE IF currentWord contains "(low)" or "(low, N)":
         - Same as above, but convert to lowercase

     ELSE IF currentWord contains "(cap)" or "(cap, N)":
         - Same as above, but capitalize each word

     ELSE IF currentWord contains "(hex)":
         - Convert previous word from hexadecimal to decimal
         - Remove "(hex)"

     ELSE IF currentWord contains "(bin)":
         - Convert previous word from binary to decimal
         - Remove "(bin)"

     ELSE IF currentWord == "a":
         - Look at next word
         - If next word starts with vowel or 'h' → change "a" to "an"

     ELSE:
         - Do nothing, move to next word

4️⃣ After processing all words:
     - Join words array back into a single string → transformedText

5️⃣ Second pass (character-level) to fix punctuation:
     FOR each character in transformedText:
         - Remove space before . , ! ? : ;
         - Ensure space after punctuation
         - Handle ... and !? correctly
         - Fix single quotes: ' word ' → 'word'

6️⃣ Write the cleaned transformedText to output file

END
```
✅ Key Notes

Two passes:

First → word-level modifications (instructions, grammar, numbers)

Second → character-level cleanup (punctuation, quotes)

Controlled access:

Always check i > 0 for previous word

Always check i + 1 < len(words) for next word

Instruction removal:

After applying instruction, remove it from array so you don’t accidentally process it again

Scalability:

Adding new rules later (like (rev) for reversing a word) is easy


🧠 PART 1 — Big Picture Architecture

Your program has 5 major stages:

1. Read file
2. Transform words
3. Join words
4. Fix punctuation
5. Write file

Each stage requires specific concepts.

Let’s break them down.

📚 PART 2 — Concepts You Must Understand Clearly

I will group them by importance.

🔹 A. File Handling (Foundation Layer)

You must understand:

1️⃣ How to read a file

Concept:

Files are read as bytes

You convert bytes → string

You must know:

os.ReadFile

How to handle errors

How to get arguments from terminal (os.Args)

You should be comfortable answering:

What is []byte?

What is string?

How do you convert between them?

🔹 B. Strings in Go (CRITICAL)

This project is 70% string manipulation.

You must understand clearly:

1️⃣ What a string is in Go

Immutable

Sequence of bytes

2️⃣ String vs rune

You don’t need deep Unicode knowledge — but you must know:

for range over string gives runes

text[i] gives byte

3️⃣ strings package functions

You must deeply understand:

strings.Fields() → splits by whitespace

strings.Join() → joins slice into string

strings.Contains() → check substring

strings.HasPrefix()

strings.HasSuffix()

strings.Trim()

strings.Replace()

strings.ToUpper()

strings.ToLower()

If you don’t master these, you’ll struggle.

🔹 C. Slices (VERY IMPORTANT)

Your words will be:

[]string

You must understand:

1️⃣ Indexing
words[i]
words[i-1]
words[i+1]
2️⃣ Bounds checking

Before using:

words[i-1]

You must ensure:

i > 0

Before:

words[i+1]

You must ensure:

i+1 < len(words)
3️⃣ Removing elements from slice

This is critical.

If you remove an element while looping:

Index shifts

You might skip words

You must understand how slice re-slicing works conceptually.

🔹 D. Loops & Control Flow

You will use:

for i := 0; i < len(words); i++ {

You must understand:

How index changes

What happens when you modify slice inside loop

When to decrement i

When to continue

🔹 E. Base Conversion (hex & bin)

You must understand:

What is base?

Decimal = base 10

Binary = base 2

Hex = base 16

Example:

10 (bin) → binary 10 → decimal 2
1E (hex) → hex 1E → decimal 30

You must know:

How to parse string in specific base

How to convert int → string

Conceptually:

string → int (base 2 or 16)
int → string (base 10)

This is handled by strconv.

You must understand:

What ParseInt does

What base argument means

🔹 F. Algorithmic Thinking (Most Important)

You must learn to think like this:

1️⃣ When you see (up)

How many words back?

Modify them

Remove instruction

2️⃣ When you see (up, 3)

Extract number

Loop backward N times

Modify safely

This requires:

Backward iteration logic

Safety checking

🔹 G. Character-Level Processing (Punctuation Phase)

After joining words:

You must scan:

for i := 0; i < len(text); i++ {

You must understand:

Detect punctuation characters

Remove space before them

Ensure space after them

Handle special cases: ..., !?

Fix quotes '

This part requires:

Careful index control

Possibly building a new string

🧠 PART 3 — Step-by-Step Implementation Order

DO NOT BUILD EVERYTHING AT ONCE.

Follow this order:

✅ Step 1 — File reading & writing only

Goal:

Read input

Write it unchanged to output

No transformation yet.

If this works → your I/O is correct.

✅ Step 2 — Split and rejoin

Read file

strings.Fields

strings.Join

Write output

Make sure spacing behaves as expected.

✅ Step 3 — Implement ONLY (up) for one word

Ignore (up, n)
Ignore others.

Make this work first.

✅ Step 4 — Extend to (up, n)

Learn:

Extract number

Loop backward

Remove two instruction tokens

✅ Step 5 — Implement (low) and (cap)

Reuse logic.

✅ Step 6 — Implement (hex) and (bin)

Learn:

Base conversion

Error handling

✅ Step 7 — Implement a → an

Forward look logic.

✅ Step 8 — Punctuation cleanup phase

Separate function.

🧠 PART 4 — Common Beginner Mistakes

Be careful about:

Accessing words[i-1] when i == 0

Removing slice element and forgetting index shifts

Not cleaning instruction tokens properly

Forgetting to handle (up, 20) when there are fewer words

Messing up punctuation loop by skipping characters
