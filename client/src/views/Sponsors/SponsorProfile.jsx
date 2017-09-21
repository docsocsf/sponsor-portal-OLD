import React from 'react';

export default class SponsorProfile extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      cvs: []
    }
  }

  componentDidMount = async () => {
    this.getCVs()
  }

  getCVs = async () => {
    try {
      const cvs = await this.props.fetchCVs()
      if (!cvs) return
      this.setState({cvs})
    } catch (e) {
      console.log("fetch cvs", e)
    }
  }

  render() {
    let {cvs} = this.state;
    cvs = cvs.map((cv, i) => <li key={i}>{cv.name}</li>)
    return (
      <div>
        <h1>Hello, Sponsor</h1>
        <ul>
          {cvs}
        </ul>
      </div>
    );
  }
}
