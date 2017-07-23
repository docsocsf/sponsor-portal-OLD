import React from 'react';
import ReactDOM from 'react-dom';
import '../style/students.scss';
import Students from 'Views/Students';

const defaultConfig = {
  method: 'GET',
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  },
  credentials: 'same-origin',
};

const fetchWithConfig = (url, configOverrides) => {
  const config = Object.assign({}, defaultConfig, configOverrides);

  return fetch(url, config);
}

const getJWTHeader = async (hdr) => {
  const token = await getToken();

  const headers = hdr || new Headers();
  headers.append('Authorization', `Bearer ${token}`);
  return headers;
}

const getToken = async () => {
  const response = await fetch("/students/auth/jwt/token", {
    credentials: 'same-origin',
  });

  return response.text();
}

const fetchUser = async () => {
  const headers = await getJWTHeader();
  const resp = await fetchWithConfig('/students/api/user', { headers })
  return resp.json()
}

ReactDOM.render(<Students fetchUser={fetchUser}/>, document.getElementById("main"));
