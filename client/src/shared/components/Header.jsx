import React from 'react';

export default class Header extends React.Component {
  render() {
    return (
      <header id="home">
        <img className="home logo" src="/assets/images/docsoc-sf-logo.svg" width="120"/>
        <h1>
          DoCSoc<br/>Sponsor Portal
        </h1>
      </header>
    );
  }
}
