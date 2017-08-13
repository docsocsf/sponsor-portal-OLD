import request from 'superagent';

const defaultConfig = {
  method: 'GET',
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  },
  credentials: 'same-origin',
};

const fetchWithConfig = (url, configOverrides) => {
  const config = {...defaultConfig, ...configOverrides};

  var req = request(config.method, url)
    .set(config.headers)
    .withCredentials(config.credentials);
  return req;
}

export default fetchWithConfig;
