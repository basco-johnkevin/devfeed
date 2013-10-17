// Generated by CoffeeScript 1.6.3
(function() {
  var __bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; };

  define(["devfeed", "golem"], function(Devfeed, Golem) {
    Devfeed.module("Entities", function(Entities, Devfeed, Backbone, Marionette, $, _) {
      var API, websocket;
      Entities.WebSocket = (function() {
        function WebSocket() {
          this.message = __bind(this.message, this);
          this.open = __bind(this.open, this);
          this.WsConn = new Golem.Connection("ws://" + CONFIG.baseUrl + "/ws", CONFIG.wsDebug);
          this.WsConn.on("open", this.open);
          this.WsConn.on("message", this.message);
        }

        WebSocket.prototype.open = function() {
          var userSession;
          userSession = Devfeed.request("user:session");
          return this.WsConn.emit("init", {
            user_id: userSession.get("id")
          });
        };

        WebSocket.prototype.message = function(data) {
          return console.log(data);
        };

        return WebSocket;

      })();
      websocket = null;
      API = {
        createWebSocket: function() {
          if (!websocket) {
            return websocket = new Entities.WebSocket();
          }
        },
        getWebSocket: function() {
          return websocket;
        }
      };
      Devfeed.on("loggedin", function() {
        return API.createWebSocket();
      });
      Devfeed.reqres.setHandler("websocket:entity", function() {
        return API.getWebSocket();
      });
      return Devfeed.commands.setHandler("websocket:create", function() {
        return API.createWebSocket();
      });
    });
    return Devfeed.Entities.WebSocket;
  });

}).call(this);
