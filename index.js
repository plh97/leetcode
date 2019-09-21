/**
 * @param {string} path
 * @return {string}
 */
var simplifyPath = function (path) {
  let arr = path.split("/").filter(e => e.length > 0);
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] == ".") {
      arr.splice(i, 1)
      i = -1
    } else if (arr[i] == "..") {
      if (i == 0) {
        arr.splice(i, 1)
      } else {
        arr.splice(i - 1, 2)
      }
      i = -1
    }
  }
  return "/" + arr.join('/')
};

console.log(
  "ans->     " +
  simplifyPath("/../"),
  simplifyPath("/a/./b/../../c/"),
  simplifyPath("/a/../../b/../c//.//"),
)