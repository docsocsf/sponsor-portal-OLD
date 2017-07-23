import React from 'react';

export default class StudentProfile extends React.Component {
  constructor() {
    super()

    this.state = {}
  }

  async componentDidMount() {
    try {
      const user = await this.props.fetchUser()
      this.setState({user})
    } catch (e) {
      console.log(e)
    }
  }

  render() {
    return (
      <div>
        <header id="home">
          <h1>
            Hello, {this.state.user ? this.state.user.name : "Student"}!
          </h1>
        </header>
      </div>
    );
  }
}
