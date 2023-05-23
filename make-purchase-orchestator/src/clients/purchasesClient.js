const fetch = require("node-fetch");
const { checkStatus } = require('../middleware/responseStatus');
const PurchaseDataDTO = require('../dtos/purchaseDataDTO');

const postPurchase = (purchaseDataDTO) => {
    return fetch(process.env.PURCHASE_URL + '/purchases', {
        method: 'POST',
        body: JSON.stringify(purchaseDataDTO),
        headers: { "Content-Type": "application/json" }
        }
    ).then(function(response) {
        checkStatus(response);
        return response.json();
    })
}

module.exports = postPurchase;