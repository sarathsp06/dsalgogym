const assert = require('chai').assert;
const isEveryCharacterUnique = require('./1.1-is_unique').isEveryCharacterUniquefunction;

describe('Testing 1.1-is_unique', _ => {
  it('Input "avonte" must return true', function () {
    assert.isTrue(isEveryCharacterUnique('avonte'));
  });

  it('Input "Droid" must return true', function () {
    assert.isTrue(isEveryCharacterUnique('Droid'));
  });

  it('Input "apple" must return false', function () {
    assert.isFalse(isEveryCharacterUnique('apple'));
  });
});