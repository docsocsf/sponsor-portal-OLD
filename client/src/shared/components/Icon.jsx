import React from 'react';

export default class Icon extends React.Component {
  render() {
    const { type } = this.props;
    return (
      <span className={`${type} icon`}></span>
    );
  }
}
