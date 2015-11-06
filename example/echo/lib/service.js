/*
 * Yuri
 * Copyright (c) 2015 Yieldbot, Inc. - All rights reserved.
 */

/* jslint node: true */
'use strict';

var seneca = require('seneca');

// Configuration
var APP_PORT = process.env.APP_PORT || '3030';

// Service
var service = seneca()
  .use(require('./index'))
  .listen({port: APP_PORT})
  .listen({type: 'tcp', port: '3031'});

// NATS
if(process.env.NATS_ENABLED) {
  service.use(require('seneca-nats-transport')).listen({type:'nats'});
}
