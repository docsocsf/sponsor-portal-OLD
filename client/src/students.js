import React from 'react';
import ReactDOM from 'react-dom';
import '../style/students.scss';
import Students from 'Views/Students';
import fetchWithConfig from './fetch';
import memoize from 'promise-memoize';

const getJWTHeader = async (hdr) => {
  const token = await getToken();

  const headers = hdr || new Headers();
  headers.append('Authorization', `Bearer ${token}`);
  return headers;
}

const getToken = memoize(async () => {
  const response = await fetchWithConfig("/students/auth/jwt/token");

  return response.text();
}, {maxAge: 60000});

const fetchUser = async () => {
  const headers = await getJWTHeader();
  const resp = await fetchWithConfig('/students/api/user', { headers })
  return resp.json()
}

ReactDOM.render(<Students fetchUser={fetchUser}/>, document.getElementById("main"));
