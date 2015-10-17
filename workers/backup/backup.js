'use strict';

var MilkCocoa = require('milkcocoa');
var APP_ID = process.env.MILK_COCOA_ID;
var milkcocoa = new MilkCocoa(APP_ID);
var redis = require('redis');

module.exports = {
  getData: function(key, callback){
    var redisClient = redis.createClient();
    redisClient.get(key, function(err, data){
      redisClient.quit();
      callback(err, data);
    });
  },

  storeData: function(ds, id, data, callback){
    var ds = milkcocoa.dataStore(ds);
    ds.set(id, data, function(err, data){
      callback(err, data);
    });
  }
};
