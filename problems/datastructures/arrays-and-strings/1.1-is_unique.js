/**
* Question:
* Implement an algorithm to determine if a string has all unique characters.  What if you can not use additional data structures.
 */

function isEveryCharacterUnique (string) {
  var specialString = ''
  for (var i of string) {
    if (specialString.includes(i)) return false
    specialString += i
  }
  return true
}

exports.isEveryCharacterUniquefunction = isEveryCharacterUnique
// console.log(isEveryCharacterUnique('avonte'))
// console.log(isEveryCharacterUnique('Droid'))
// console.log(isEveryCharacterUnique('apple'))
