const path = require('path');
const paths = require('./paths')

const { CleanWebpackPlugin } = require('clean-webpack-plugin')
const CopyWebpackPlugin = require('copy-webpack-plugin')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const VueLoaderPlugin = require('vue-loader/lib/plugin')

module.exports = {
    // Where webpack looks to start building the bundle
    entry: [ 'whatwg-fetch', paths.src + '/index.js'],

    resolve: {
        extensions: [ '.js', '.vue' ],
        alias: {
            'components': path.resolve(__dirname, '../src/components/'),
            'view': path.resolve(__dirname, '../src/view/'),
            'images': path.resolve(__dirname, '../src/images/'),
            'styles': path.resolve(__dirname, '../src/styles/'),
            'config' : path.resolve(__dirname, '../src/config/'),
            'store' : path.resolve(__dirname, '../src/store/'),
            'router' : path.resolve(__dirname, '../src/router/'),
            'libs' : path.resolve(__dirname, '../src/libs/'),
        }
    },

    // Where webpack outputs the assets and bundles
    output: {
        path: paths.build,
        filename: '[name].bundle.js',
        publicPath: '/',
    },


    // Customize the webpack build process
    plugins: [

        // Vue plugin for the magic
        new VueLoaderPlugin(),

        // Removes/cleans build folders and unused assets when rebuilding
        // new CleanWebpackPlugin(),

        // Copies files from target to destination folder
        new CopyWebpackPlugin({
            patterns: [
                {
                    from: paths.public,
                    to: 'assets',
                    globOptions: {
                        ignore: ['*.DS_Store'],
                    },
                    noErrorOnMissing: true
                },
            ],
        }),

        // Generates an HTML file from a template
        // Generates deprecation warning: https://github.com/jantimon/html-webpack-plugin/issues/1501
        new HtmlWebpackPlugin({
            title: 'webpack Boilerplate',
            favicon: paths.src + '/images/favicon.png',
            template: paths.src + '/template.html', // template file
            filename: 'index.html', // output file
        }),
    ],

    // Determine how modules within the project are treated
    module: {
        rules: [
            // JavaScript: Use Babel to transpile JavaScript files
            {test: /\.vue$/, loader: 'vue-loader' },
            {test: /\.js$/, exclude: /node_modules/, use: ['babel-loader']},

            // Styles: Inject CSS into the head with source maps
            {
                test: /\.(less|css)$/,
                use: [
                    // Note: Only style-loader works for me !!!
                    // 'vue-style-loader',
                    'style-loader',
                    {loader: 'css-loader', options: {sourceMap: true, importLoaders: 1}},
                    {loader: 'postcss-loader', options: {sourceMap: true}},
                    {loader: 'less-loader', options: {sourceMap: true}},
                ],
            },

            // Images: Copy image files to build folder
            {test: /\.(?:ico|gif|png|jpg|jpeg)$/i, type: 'asset/resource'},

            // Fonts and SVGs: Inline files
            {test: /\.(woff(2)?|eot|ttf|otf|svg|)$/, type: 'asset/inline'},
        ],
    },
}