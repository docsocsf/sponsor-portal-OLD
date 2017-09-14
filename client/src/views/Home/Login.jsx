import React from 'react';
import Header from 'Components/Header';
import md5 from 'md5';
import fetchWithConfig from '../../fetch';

export default class Login extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      email: '',
      password: '',
    };

    this.handleChange = this.handleChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  async onSubmit(event) {
    event.preventDefault()
    let { email, password } = this.state;
    password = md5(password)
    try {
      console.log("login", email, password);
      let resp = await fetchWithConfig("/sponsors/auth/login",
        { method: "POST",
          data: {email, password},
          type: 'form',
          headers:
          {'Content-Type': 'application/x-www-form-urlencoded' }
        });
      console.log(resp)
    } catch (e) {
      this.setState({error: "The email and/or password you entered was incorrect"});
    }
  }

  handleChange(event) {
    this.setState({
      [event.target.name]: event.target.value
    })
  }

  render() {
    return (
      <div>
        <Header />
        <div id="login-form">
          { !!this.state.error &&
          <h4 className="error">{this.state.error}</h4> }
          <form onSubmit={this.onSubmit}>
            <input
              type="text"
              name="email"
              placeholder="Email"
              value={this.state.email}
              onChange={this.handleChange}/>
            <input
              type="password"
              name="password"
              placeholder="Password"
              value={this.state.password}
              onChange={this.handleChange}/>
            <input type="submit" value="Login" />
          </form>
        </div>
      </div>
    );
  }
}
