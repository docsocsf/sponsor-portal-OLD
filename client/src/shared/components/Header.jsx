import React from 'react';

export default class Header extends React.Component {
  render() {
    const { sponsor, name } = this.props;
    const isSponsor = !!sponsor;
    const sponsorBanner = (<span className="sponsor"> &times; {sponsor}</span>);

    return (
      <header id="home">
        <div className="container">
          <div className="logos">
            <img className="home logo" src="/assets/images/docsoc-sf-logo.svg" width="30"/>
            { isSponsor && sponsorBanner }
          </div>
          <div className="menu">
            { !!name && (<span>Hey, {name}!</span>)}
          </div>
        </div>
      </header>
    );
  }
}
