const express = require('express');
const axios = require('axios');

const router = express.Router();

module.exports = (params) => {
  const { apiUrl, apiToken } = params;

  router.get('/', async (req, res, next) => {
    try {
      return res.json({
        version: process.env.VERSION,
        apiUrl: process.env.API_URL
      });
    } catch (err) {
      return next(err);
    }
  });

  router.all('*', async (req, res, next) => {
    try {
      let method = req.method;
      let url = `${apiUrl}${req.url}`;
      let data = req.params;
      let headers = {
        'x-token': apiToken,
        'content-type': 'application/json'
      };
      const result = await axios({ method, url, data, headers });
      return res.json(result);
    } catch (err) {
      return next(err);
    }
  });

  return router;
};
