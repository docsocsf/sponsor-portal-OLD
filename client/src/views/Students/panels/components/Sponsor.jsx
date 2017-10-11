import React from 'react';
import Icon from 'Components/Icon';

const prefix = "/assets/images/sponsors"

export default class Sponsor extends React.Component {
  render() {
    const { tier, info: { name, url, logo, description, email, apply, bespoke } } = this.props;
    return (
      <section className="sponsor-card">
        { bespoke && tier && <Icon type={`${tier} bespoke`} /> }
        <div className="summary">
          <img src={`${prefix}/${logo}`} alt={name} className="logo" />
          <div className="links">
            <span>
              <Icon type="link" />
              <a href={url} className="website">Website</a>
            </span>
            { email && (
              <span>
                <Icon type="email" />
                <a href={`mailto:${email}`} className="apply">Apply</a>
              </span>
            )}
            { apply && (
              <span>
                <Icon type="link" />
                <a href={apply} className="apply">Apply</a>
              </span>
            )}
          </div>
        </div>
        <main>
          {description && description.split('\n').map((item, key) => {
            return <p key={key}>{item}</p>
          })}
      </main>
      </section>
    );
  }
}
