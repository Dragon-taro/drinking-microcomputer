import socketio from "socket.io";
import express from "express";
import { createServer } from "http";
import * as bodyParser from "body-parser";

const app = express();
const http = createServer(app);
const io = socketio(http, { serveClient: false });

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());

app.get("/ws/hc", (req, res) => {
  res.send("ok");
});

app.post("/ws/notify", (req, res) => {
  io.emit("message", req.body);
  res.send("ok");
});

io.on("connection", () => {
  console.log("connection");
});

http.listen(3001, function() {
  console.log("server listening. Port:" + 3001);
});
