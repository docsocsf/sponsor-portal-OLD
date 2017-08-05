import React from 'react';
import FileUploadDialog from 'Components/FileUploadDialog';

export default class StudentProfile extends React.Component {
  constructor() {
    super()

    this.state = {}
    this.getUser = this.getUser.bind(this);
  }

  componentDidMount() {
    this.getUser()
  }

  async getUser() {
    try {
      const user = await this.props.fetchUser()
      this.setState({user})
    } catch (e) {
      console.log("fetch user", e)
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
        <section id="cv">
          <h2>Upload CV</h2>
          <FileUploadDialog accept="application/pdf" className="cv" multiple={false}/>
        </section>
      </div>
    );
  }
}
