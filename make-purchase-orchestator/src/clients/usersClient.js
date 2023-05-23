const fetch = require("node-fetch");
const { checkStatus } = require('../middleware/responseStatus');

const findCustomer = (customer_id) => {
    return fetch(process.env.USER_URL + '/customers/' + customer_id, {
        method: 'GET'
        }
    ).then(function(response) {
        checkStatus(response);
        return response.json();
    })
}

const findSeller = (seller_id) => {
    return fetch(process.env.USER_URL + '/sellers/' + seller_id, {
        method: 'GET'
        }
    ).then(function(response) {
        checkStatus(response);
        return response.json();
    })
}


module.exports = { findCustomer, findSeller };