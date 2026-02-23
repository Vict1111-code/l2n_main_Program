def check_palindrome(s):
    n_s = []

    for char in s:
            if char.isalnum():
                n_s.append(char.lower())

    left = 0
    right = len(n_s) -1

    while left < right:
        if n_s[left] != n_s[right]:
            return (False, left, right)
        left += 1
        right -= 1

    return True

print(check_palindrome("race ,.,car"))