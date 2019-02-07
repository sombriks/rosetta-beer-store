const path = require("path");

const HtmlWebpackPlugin = require("html-webpack-plugin");
const CleanWebpackPlugin = require("clean-webpack-plugin");

const webpack = require("webpack");

module.exports = {
  mode: process.env.NODE_ENV || "development",
  module: {
    rules: [
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"]
      },
      {
        test: /\.html$/,
        loader: "html-loader"
      },
      {
        test: /\.(png|svg|jpg|gif|woff|woff2|eot|ttf|otf)$/,
        use: ["file-loader"]
      },
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: ["babel-loader"]
      }
    ]
  },
  resolve: {
    extensions: ["*", ".js", ".jsx"]
  },
  entry: "./src/main.jsx",
  output: {
    filename: "build.js",
    path: path.resolve(__dirname, "dist")
  },
  devtool:
    process.env.NODE_ENV == "development" ? "inline-source-map" : undefined,
  devServer: {
    contentBase: "./dist",
    hot: true
  },
  plugins: [
    new CleanWebpackPlugin(["dist"]),
    new HtmlWebpackPlugin({
      template: "./index.html"
    }),
    new webpack.HotModuleReplacementPlugin()
  ]
};
