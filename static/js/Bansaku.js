define(
      "Bansaku",
      [],
      function(){
        function Bansaku(model){
          if(model !== undefined){
            console.log("set observable model");
            console.log(model.count);
            this.count = ko.observable(model.count);
          }else{
            this.count = ko.observable(0);
          }
        }
        return Bansaku;
      }
    );
