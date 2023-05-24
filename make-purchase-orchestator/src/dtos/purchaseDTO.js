const errors = require('../exceptions/APIErrors');

class PurchaseDTO {
    constructor(body) {
        this.customer_id = body.customer_id;
        this.product_id = body.product_id;
        this.units = body.units;
    }

    validate() {
        if(!this.customer_id || !this.product_id || this.units == undefined) {
            throw new errors.InvalidPurchase();
        }

        if(this.units < 1) {
            throw new errors.InvalidUnits();
        }
    }
}

module.exports = PurchaseDTO;