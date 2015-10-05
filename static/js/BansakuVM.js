define(
    "BansakuVM",
    [
      "Bansaku"
    ],
    function(Bansaku){
      function BansakuVM(ws){
        var that = this;
        this.bansaku = ko.observable(new Bansaku());
        this.send = function(){
         var model = {
          "bansaku": "tsukita"
         } 
         ws.send($.toJSON(model));
         ion.sound.play("bansaku");
        };

        ws.onmessage = function(e){
          var model = $.evalJSON(e.data);
          that.bansaku(model);
        };
      }
      return BansakuVM;
    }
    );
