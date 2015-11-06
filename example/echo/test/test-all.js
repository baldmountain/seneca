/*
 * Yuri
 * Copyright (c) 2015 Yieldbot, Inc. - All rights reserved.
 */

/* jslint node: true, expr: true */
/* global describe: false, it: false */
'use strict';

var expect = require('chai').expect;

// Tests

var plugin = 'echo',
    seneca = require('seneca')({log: 'silent', strict: {result: false}})
      .use(require('../lib/index'));

/*
 * This is a utility method to make sure mocha done() is called even if an exception is thrown
 */
var testIt = function testIt(func, done) {
  var err;
  try { func(); } catch (e) { err = e; } finally { done(err); }
};

describe(plugin, function() {

  it('should verify echo cmd', function(done) {
    var action = {role: plugin, cmd: 'echo', 'msg':'hello'};
    seneca.act(action, function(err, res) {
      testIt(function() {
        expect(err).to.not.exist;
        expect(action.role).to.equal(res.role);
        expect(action.cmd).to.equal(res.cmd);
        expect(action.msg).to.equal(res.msg);
      }, done);
    });
  });

});
