define(
    "BansakuMain",
    [
    "BansakuVM"
    ],
    function(BansakuVM){
      ion.sound({
        sounds: [
        {
          name: "bansaku"
        }],
        path: "../sound/",
        preload: true,
        multiplay: true,
        volume: 0.6
      });
      var ws = new WebSocket("ws://localhost:60000/ws");
      var bansakuVM = new BansakuVM(ws);
      ko.applyBindings(bansakuVM);
    }
);
