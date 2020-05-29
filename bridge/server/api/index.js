const express = require('express');
const axios = require('axios');

const router = express.Router();

module.exports = (params) => {
  const { apiUrl, apiToken } = params;
  const version = process.env.VERSION;

  router.get('/', async (req, res, next) => {
    try {
      return res.json({ version, apiUrl });
    } catch (err) {
      return next(err);
    }
  });

  router.all('*', async (req, res, next) => {
    try {
      const result = await axios({
        method: req.method,
        url: `${apiUrl}${req.url}`,
        data: req.params,
        headers: {
          'x-token': apiToken,
          'content-type': 'application/json'
        }
      });
      return res.json(result.data);
    } catch (err) {
      return next(err);
    }
  });

  return router;
};
