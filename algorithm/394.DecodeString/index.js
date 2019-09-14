/**
 * @param {string} s
 * @return {string}
 */
var decodeString = function (s) {
  console.log(s);
  if (s.split("").indexOf('[') == -1) return s;
  let res=""
  let startIndex, endIndex = 0;
  // startIndex = s.split("").indexOf('[');
  // endIndex = s.split("").lastIndexOf(']');
  s.split("").forEach((e,i) => {
    if (e==)
    number = s.match(/\d+(?=\[)/g)
    console.log(startIndex, endIndex, number);
    let a = s.substring(-1, startIndex - number.length)
    let b = decodeString(s.substring(startIndex + 1, endIndex))
    let c = s.substring(endIndex + 1)
    return a + b.repeat(number) + c
  })




  return res
};


console.log(
  // decodeString("100[leetcode]"),
  decodeString("3[a]2[bc]"),
)