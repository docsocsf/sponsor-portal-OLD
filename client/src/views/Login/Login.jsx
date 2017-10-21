import React from 'react';
import Header from 'Components/Header';
import md5 from 'md5';
import request from 'superagent';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';

export default class Login extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      email: '',
      password: '',
    };
  }

  onSubmit = async (event) => {
    event.preventDefault()
    try {
      let resp = await request
        .post('/auth/sponsors/login')
        .send(new FormData(document.getElementById("sponsor-login-form")))
      let redirect = new URL(resp.xhr.responseURL).pathname;
      window.location.replace(redirect)
    } catch (e) {
      this.setState({error: "The email and/or password you entered was incorrect."});
    }
  }

  handleChange = (event) => this.setState({ [event.target.name]: event.target.value })

  render() {
    let {error, email, password } = this.state;

    return (
      <div id="login-view">
        <div id="login">
          <img className="logo" src="/assets/images/docsoc-sf-logo.svg" width="60"/>
          <h1>
            DoCSoc<br/>Sponsor Portal
          </h1>
          <Tabs>
            <TabList>
              <Tab>Student Login</Tab>
              <Tab>Sponsor Login</Tab>
            </TabList>

            <TabPanel>
                <button>
                  <a href="/students">Login with Imperial</a>
                </button>
            </TabPanel>
            <TabPanel>
              <LoginForm
                onSubmit={this.onSubmit}
                handleChange={this.handleChange}
                id="sponsor-login-form"
                email={email}
                password={password}
                error={error}
              />
            </TabPanel>
          </Tabs>
        </div>
      </div>
    );
  }
}

const LoginForm = (props) => (
  <form onSubmit={props.onSubmit} id={props.id} className="stacked centered">
    { !!props.error &&
        <p className="error">{props.error}</p>
    }
    <input
      type="text"
      name="email"
      placeholder="Email"
      value={props.email}
      onChange={props.handleChange}/>
    <input
      type="password"
      name="password"
      placeholder="Password"
      value={props.password}
      onChange={props.handleChange}/>
    <input type="submit" value="Login" />
  </form>
);
