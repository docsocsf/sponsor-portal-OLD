import request from 'request-promise-native';

const defaultConfig = {
  method: 'GET',
  baseUrl: 'http://localhost:8080',
  headers: {
    Accept: 'application/json',
  }
};

const fetchWithConfig = (url, configOverrides) => {
  const config = {...defaultConfig, ...configOverrides, url};

  return request(config)
}

export default fetchWithConfig;
