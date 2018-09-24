var loc = window.location;
var uri = 'ws:';

uri += '//' + loc.host;
uri += loc.pathname + 'ws';
ws = new WebSocket(uri);

ws.onopen = function() {
  console.log('Подключаемся к вёбсокету');
};
ws.onmessage = function(evt) {
  var out = document.getElementById('search');
  out.innerHTML += evt.data;
};

setInterval(function() {
  ws.send('Hello, server!');
}, 1000);
