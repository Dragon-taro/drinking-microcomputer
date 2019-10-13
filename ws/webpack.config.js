const path = require("path");

module.exports = {
  entry: ["./index.ts"],
  output: {
    path: path.join(__dirname, "ws-server"),
    filename: "server.bundle.js"
  },
  resolve: {
    modules: ["node_modules"],
    extensions: [".ts", ".js"]
  },
  module: {
    rules: [
      {
        test: /\.ts$/,
        exclude: /node_modules/,
        use: [
          {
            loader: "ts-loader"
          }
        ]
      }
    ]
  },
  target: "node"
};
