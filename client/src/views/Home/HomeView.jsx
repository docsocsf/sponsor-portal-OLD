import React from 'react';

export default class HomeView extends React.Component {
  render() {
    return (
      <div>
        <header id="home">
          <img className="home logo" src="/images/docsoc-sf-logo.svg" width="120"/>
          <h1>
            DoCSoc<br/ >Sponsor Portal
          </h1>
        </header>
        <p>
          Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod
          tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At
          vero eos et accusam et justo duo dolores et ea rebum.
        </p>
        <p>
          Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod
          tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At
          vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren,
          no sea takimata sanctus est Lorem ipsum dolor sit amet.
        </p>
        <div id="login">
          <button className="student">
            <a href="/students/auth/login">
              Login as a Student
            </a>
          </button>
          <button className="sponsor">
            Login as a DoCSoc sponsor
          </button>
        </div>
      </div>
    );
  }
}
