var maxDepth = function (root) {
  if (root.children.length < 1) return 0;
  return 1 + max(root.children.map(e=>maxDepth(e)))
};

const max = (arr) => {
  var res = arr[0]
  arr.forEach(e => {
    if (e > res) {
      res = e
    }
  })
  return res
}

const res = maxDepth({
  "$id": "1",
  "children": [{
    "$id": "2",
    "children": [{
      "$id": "5",
      "children": [],
      "val": 5
    }, {
      "$id": "6",
      "children": [],
      "val": 6
    }],
    "val": 3
  }, {
    "$id": "3",
    "children": [],
    "val": 2
  }, {
    "$id": "4",
    "children": [],
    "val": 4
  }],
  "val": 1
})
console.log(res);