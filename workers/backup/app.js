var backup = require('./backup.js');

backup.getData("count", function(err, count){
  var object = {
    count: count
  }
  backup.storeData("count", "count", object, function(err, data){
    process.exit();
  });
});

