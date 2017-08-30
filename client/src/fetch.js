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

  return fetch(url, config);
}

export default fetchWithConfig;
