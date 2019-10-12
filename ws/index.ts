import * as socketio from "socket.io";
import * as express from "express";
import { Server } from "http";
import * as bodyParser from "body-parser";

const app = express();
const http = new Server(app);
const io = socketio(http);

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());

app.get("/ws/hc", (req, res) => {
  res.send("ok");
});

app.post("/ws/notify", (req, res) => {
  io.of("/ws").emit("message", req.body);
  res.send("ok");
});

io.of("/ws").on("connection", () => {
  console.log("connection");
});

http.listen(3001, function() {
  console.log("server listening. Port:" + 3001);
});
