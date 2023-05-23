const express = require('express');
const app = express();
const morgan = require('morgan');
const routes = require('./routes/api');
const env = require('dotenv').config();
const errorHandler = require('./middleware/errorHandler');
require("express-async-errors");

//Configuraciones
app.set('port', process.env.PORT || 50054);
app.set('json spaces', 2);

//Middleware
app.use(morgan('dev'));
app.use(express.urlencoded({extended:false}));
app.use(express.json());

//Routes
app.use(routes);

app.use((req, res) => {
    res.status(404);
    res.json({
        status: 404,
        errorCode: 'RESOURCE_NOT_FOUND'
    }); 
});

app.use(errorHandler);

//Iniciando el servidor
app.listen(app.get('port'),()=>{
    console.log(`Server listening on port ${app.get('port')}`);
});

