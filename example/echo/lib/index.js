/*
 * Yuri
 * Copyright (c) 2015 Yieldbot, Inc. - All rights reserved.
 */

/* jslint node: true */
'use strict';

module.exports = function(options) {

  var seneca = this,
      plugin = 'echo';

  /*
   * REMOVE THIS COMMENT AFTER READ IT
   * Here we set options to something. Normally you'd merge some defaults into the
   * Options passed in, but we'll leave it up to you to decide how to do that.
   * Example: `options = seneca.util.deepextend({foo:bar}, options);`
   */
  options = options || {};

  // Add action patterns

  /*
   * REMOVE THIS COMMENT AFTER READ IT
   * Note: if you return something that isn't a plain JavaScript object, Seneca will
   * print a warning unless you initialize seneca with the `{ strict: { result: false } }` option.
   */
  // Echo
  seneca.add({role: plugin, cmd: 'echo'}, function(msg, done) {
    return done(null, msg);
  });

  // Return
  return {
    name: plugin
  };

};
