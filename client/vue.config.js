module.exports = {
    devServer: {
        // This will forward any request that does not match a static file to localhost:3000
        proxy: 'https://api.nbmfscafdcasc.xyz'
    }
}