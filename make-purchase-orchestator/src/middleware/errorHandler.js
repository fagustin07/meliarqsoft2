const errors = require('../exceptions/APIErrors');

function errorHandler(err, req, res, next) {
    console.log(err);
    if (err instanceof errors.InvalidPurchase) {
        res.status(400).json({ error: err.name, message: err.message});
    }
    else if (err instanceof errors.InvalidUnits) {
        res.status(400).json({ error: err.name, message: err.message});
    }
    else if (err instanceof errors.ResponseError) {
        res.status(400).json({ error: err.name, message: err.message});
    }
    else if (err instanceof errors.PurchaseNotCreated) {
        res.status(400).json({ error: err.name, message: err.message});
    }
    else {
        res.status(500).json({ error: 'Ha ocurrido un error en el servidor.' });
    }
}

module.exports = errorHandler;
