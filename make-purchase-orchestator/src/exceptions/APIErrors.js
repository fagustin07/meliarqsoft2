class APIError extends Error {
    constructor(name, statusCode, message = null) {
        super(message || name);
        this.name = name;
        this.status = statusCode;
        this.message = message;
    }
}

class InvalidPurchase extends APIError {
    constructor() {
        super('Invalid Purchase', 400, 'INVALID_INPUT_DATA');
    }  
}

class InvalidUnits extends APIError {
    constructor() {
        super('Units cannot be less than 1', 400, 'INVALID_INPUT_DATA');
    }  
}

class ResponseError extends APIError {
    constructor(status, message) {
        super(message, status, 'ERROR');
    }
}

module.exports = {
    InvalidPurchase,
    InvalidUnits,
    ResponseError
};