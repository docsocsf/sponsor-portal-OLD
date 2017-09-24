import React from 'react';
import Header from 'Components/Header';
import md5 from 'md5';
import request from 'superagent';
import {Redirect} from 'react-router';

export default class Login extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      email: '',
      password: '',
      redirect: undefined
    };
  }

  onSubmit = async (event) => {
    event.preventDefault()
    let { email, password } = this.state;
    password = md5(password)
    try {
      let resp = await request
        .post('/sponsors/auth/login')
        .type('form')
        .send({ email })
        .send({ password })
      let redirect = new URL(resp.xhr.responseURL).pathname;
      this.setState({redirect});
    } catch (e) {
      console.log(e);
      this.setState({error: "The email and/or password you entered was incorrect"});
    }
  }

  handleChange = (event) => this.setState({ [event.target.name]: event.target.value })

  render() {
    let {redirect, error, email, password } = this.state;
    if (!!redirect) {
      return (<Redirect to={redirect} />);
    }

    return (
      <div>
        <Header />
        <div id="login-form">
          { !!error &&
              <h4 className="error">{error}</h4>
          }
          <form onSubmit={this.onSubmit}>
            <input
              type="text"
              name="email"
              placeholder="Email"
              value={email}
              onChange={this.handleChange}/>
            <input
              type="password"
              name="password"
              placeholder="Password"
              value={password}
              onChange={this.handleChange}/>
            <input type="submit" value="Login" />
          </form>
        </div>
      </div>
    );
  }
}
