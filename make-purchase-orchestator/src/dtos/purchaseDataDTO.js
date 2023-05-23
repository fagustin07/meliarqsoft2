const errors = require('../exceptions/APIErrors');

class PurchaseDataDTO {
    constructor(customer, product, seller, units) {
        this.buyer_email = customer.email,
        this.buyer_name = customer.name,
        this.id_user = customer.id,
        this.id_product = product.id,
        this.product_name = product.name,
        this.seller_name = seller.businessName,
        this.seller_email = seller.email,
        this.units = units,
        this.total = product.price * units
    }
}

module.exports = PurchaseDataDTO;