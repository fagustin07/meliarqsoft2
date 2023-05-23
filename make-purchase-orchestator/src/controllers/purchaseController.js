const { findCustomer, findSeller } = require('../clients/usersClient');
const { findProduct, updateProduct } = require('../clients/productsClient');
const PurchaseDTO = require('../dtos/purchaseDTO');
const PurchaseDataDTO = require('../dtos/purchaseDataDTO');

const createPurchase = (req, res, next) => {
    const purchaseDTO = new PurchaseDTO(req.body);
    purchaseDTO.validate();

    return processPurchase(purchaseDTO);
}

const processPurchase = (purchaseDTO) => {
    return getCustomer(purchaseDTO.customer_id)
        .then(customer => {
            return getProduct(purchaseDTO.product_id)
                .then(product => {
                    return getSeller(product.id_seller)
                        .then(seller => {
                            return reduceStock(product, purchaseDTO.units)
                                .then(newProd => {
                                    const purchaseDataDTO = new PurchaseDataDTO(customer, product, seller, purchaseDTO.units);
                                    return newProd;//postPurchase(purchaseDataDTO);
                                })

                        })
                })
        })
}

const getCustomer = (customer_id) => {
    return findCustomer(customer_id)
        .then(customer => {
            return customer;
        })
} 

const getProduct = (product_id) => {
    return findProduct(product_id)
        .then(product => {
            return product
        })
}

const getSeller = (seller_id) => {
    return findSeller(seller_id)
        .then(seller => {
            return seller;
        })
}

const reduceStock = (product, units) => {
    product.stock =- units;
    return updateProduct(product)
        .then(product => {
            return product;
        })

}

module.exports = { createPurchase };
