import React from 'react';
import FileUploadDialog from 'Components/FileUploadDialog';
import request from 'superagent';
import config from './mock-config';
import mocker from 'superagent-mock';

const logger = log => console.log("mock", log);
const mock = mocker(request, config, logger);

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

  async uploadCV(files, progress) {
    if (files.length > 1) {
      throw new Error("Expecting exactly 1 CV")
    }

    console.log(config[0].post.toString());

    let data = await request
      .post('/cv')
      .attach('cv', files[0]).
      on('progress', event => {
        progress(event.percent);
      });
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
          <FileUploadDialog
            accept="application/pdf"
            className="cv"
            multiple={false}
            onUpload={this.uploadCV}
          />
        </section>
      </div>
    );
  }
}
