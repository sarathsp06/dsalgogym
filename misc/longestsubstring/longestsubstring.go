package longestsubstring
/**
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
    myMap := map[byte]int{}
    var subStart int
    var currentMax int
    bArr := []byte(s)
    for i, char := range bArr {
        if j, ok := myMap[char]; ok {
        if currentMax < i - subStart + 1 {
            currentMax = len(bArr[subStart:i])
        }    
        if j >= subStart {
                subStart = j + 1   
            }
        }
        myMap[char] = i
    }
    if len(bArr[subStart:]) > currentMax {
        return len(bArr[subStart:])
    }
    return currentMax
}