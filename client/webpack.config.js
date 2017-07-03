var path = require('path');
var webpack = require('webpack');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

const extractSass = new ExtractTextPlugin({
  filename: "[name].bundle.css",
  allChunks: true
});

module.exports = {
  entry: {
    main: "./src/index.js",
  },
  output: {
    path: path.resolve("dist/assets"),
    filename: "[name].bundle.js",
  },
  module: {
    rules: [
      {
        test: /\.jsx?/,
        exclude: /(node_modules)/,
        loader: ["babel-loader"]
      },
      {
        test: /\.scss$/,
        use: extractSass.extract({
          use: [{
            loader: 'css-loader'
          },
          {
            loader: 'sass-loader'
          }]
        })
      },
    ]
  },
  plugins: [
    extractSass,
  ]
};
