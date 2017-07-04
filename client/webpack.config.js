var path = require('path');
var webpack = require('webpack');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

const extractSass = new ExtractTextPlugin({
  filename: "[name].bundle.css",
  allChunks: true
});

module.exports = {
  entry: {
    index: "./src/index.js",
  },
  output: {
    path: path.resolve("dist/assets"),
    filename: "[name].bundle.js",
  },
  resolve: {
    alias: {
      Views: path.resolve(__dirname, 'src/views'),
      Components: path.resolve(__dirname, 'src/shared/components')
    },
    extensions: [".js", ".jsx"],
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
