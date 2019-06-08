# [819. Most Common Word](https://leetcode.com/problems/most-common-word/)

## 2019/06/08

### 题目 💗[easy]

Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words. It is guaranteed there is at least one word that isn't banned, and that the answer is unique.

给定一个段落和一个列表单词,返回这个段落里面重复次数最长的单词

Words in the list of banned words are given in lowercase, and free of punctuation. Words in the paragraph are not case sensitive. The answer is in lowercase.

Example:

```bash
Input:
paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
banned = ["hit"]
Output: "ball"
Explanation:
"hit" occurs 3 times, but it is a banned word.
"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph.
Note that words in the paragraph are not case sensitive,
that punctuation is ignored (even if adjacent to words, such as "ball,"),
and that "hit" isn't the answer even though it occurs more because it is banned.
```

Note:

1. 1 <= paragraph.length <= 1000.
2. 0 <= banned.length <= 100.
3. 1 <= banned[i].length <= 10.
4. The answer is unique, and written in lowercase (even if its occurrences in paragraph may have uppercase symbols, and even if it is a proper noun.)
5. paragraph only consists of letters, spaces, or the punctuation symbols !?',;.
6. There are no hyphens or hyphenated words.
7. Words only consist of letters, never apostrophes or other punctuation symbols.
