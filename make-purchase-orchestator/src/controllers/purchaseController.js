const { findCustomer, findSeller } = require('../clients/usersClient');
const { findProduct, updateProduct } = require('../clients/productsClient');
const { postPurchase } = require('../clients/purchasesClient')
const PurchaseDTO = require('../dtos/purchaseDTO');
const PurchaseDataDTO = require('../dtos/purchaseDataDTO');
const errors = require('../exceptions/APIErrors');


const createPurchase = (req, res, next) => {
    const purchaseDTO = new PurchaseDTO(req.body);
    purchaseDTO.validate();

    return processPurchase(purchaseDTO, next);
}

const processPurchase = (purchaseDTO, next) => {
    return getCustomer(purchaseDTO.customer_id)
        .then(customer => {
            return getProduct(purchaseDTO.product_id)
                .then(product => {
                    return getSeller(product.id_seller)
                        .then(seller => {
                            return reduceStock(product, purchaseDTO.units)
                                .then(newProd => {
                                    const purchaseDataDTO = new PurchaseDataDTO(customer, newProd, seller, purchaseDTO.units);
                                    return postPurchase(purchaseDataDTO)
                                        .then(response => {
                                            return response
                                        })
                                        .catch(err => {
                                            // Si hubo un error, actualizo el producto con su stock original
                                            return updateProduct(product)
                                                .then(res => {throw err;})
                                        });
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
    let newProd = { ...product };
    newProd.stock = product.stock - units;
    return updateProduct(newProd)
        .then(res => {
            return newProd;
        })
}

module.exports = { createPurchase };
