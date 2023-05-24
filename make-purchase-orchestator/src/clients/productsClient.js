const fetch = require("node-fetch");
const { checkStatus } = require('../middleware/responseStatus');

const findProduct = (product_id) => {
    return fetch(process.env.PRODUCT_URL + '/products/' + product_id, {
        method: 'GET'
        }
    ).then(function(response) {
        checkStatus(response);
        return response.json();
    })
}

const updateProduct = (product) => {
    return fetch(process.env.PRODUCT_URL + '/products/' + product.id, {
        method: 'PUT',
        body: JSON.stringify(product),
        headers: { "Content-Type": "application/json" }
        }
    ).then(function(response) {
        checkStatus(response);
        return response;
    })
}

module.exports = { findProduct, updateProduct };