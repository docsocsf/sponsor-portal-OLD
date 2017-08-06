import React from 'react';

export default class Progress extends React.Component {
  render() {
    const { percent } = this.props;
    return (
      <progress max="100" value={percent}>
        <div className="progress-bar">
          <span style={{width: percent + "%"}}>Loading...</span>
        </div>
      </progress>
    );
  }
}
