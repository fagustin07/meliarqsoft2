const errors = require('../exceptions/APIErrors')

const checkStatus = (response) => {
    if(response.status.toString()[0] != '2') {
        throw new errors.ResponseError(response.status, response.statusText);
    }
}

module.exports = { checkStatus };