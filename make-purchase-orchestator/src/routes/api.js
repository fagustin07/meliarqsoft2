const { createPurchase } = require('../controllers/purchaseController');
const { Router } = require('express');
const router = Router();
const errors = require('../exceptions/APIErrors');

// Raiz
router.get('/', (req, res) => {
    res.json({"status": "working"})
})

// Crear compra
router.post('/create', (req, res, next) => {
    createPurchase(req, res, next)
        .then(response => { res.json(response); })
        .catch(err => {next(err)});
});


module.exports = router;