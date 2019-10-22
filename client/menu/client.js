const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listDishes: () => client.get('/dishes'),
        createOrder: (ids, table_number) => client.post('/dishes', { ids, table_number })
    }

};

module.exports = { Client };
