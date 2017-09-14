import request from 'superagent';

const defaultConfig = {
  method: 'GET',
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  },
  credentials: 'same-origin',
  type: 'json'
};

const fetchWithConfig = (url, configOverrides) => {
  const config = {...defaultConfig, ...configOverrides};

  if (config.method === "POST") {
    console.log("POST", config.data)
    let req = request(config.method, url)
      .type(config.type)
      .send(config.data)
      .set(config.headers)
      .withCredentials(config.credentials);
    console.log(req);
    return req;
  }

  return request(config.method, url)
    .set(config.headers)
    .withCredentials(config.credentials);
}

export default fetchWithConfig;
