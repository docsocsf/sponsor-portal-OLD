import React from 'react';
import FileUploadDialog from 'Components/FileUploadDialog';
import request from 'superagent';
import getJWTHeader from '../../jwt';

export default class StudentProfile extends React.Component {
  constructor() {
    super()

    this.state = {}
    this.getUser = this.getUser.bind(this);
    this.getCV = this.getCV.bind(this);
    this.uploadCV = this.uploadCV.bind(this);
  }

  async componentDidMount() {
    await this.getUser()
    this.getCV()
  }

  async getUser() {
    try {
      const user = await this.props.fetchUser()
      this.setState({user})
    } catch (e) {
      console.log("fetch user", e)
    }
  }

  async getCV() {
    try {
      const cv = await this.props.fetchCV()
      if (!cv) return
      this.setState({cv, upload: false})
      this.fileRef.updateFile([cv])
    } catch (e) {
      console.log("fetch cv", e)
    }
  }

  async uploadCV(files, progress) {
    if (files.length > 1) {
      throw new Error("Expecting exactly 1 CV")
    }

    let headers = await getJWTHeader();

    try {
      let data = await request
        .post('/students/api/cv')
        .set(headers)
        .attach('cv', files[0]).
        on('progress', event => {
          progress(event.percent);
        });

      this.setState({upload: false})
    } catch (e) {
      console.log(e)
      console.log("inspect", fs)
      throw new Error("Failed to upload file")
    }
  }

  render() {
    let {user, cv, upload} = this.state;
    return (
      <div>
        <header id="home">
          <h1>
            Hello, {user ? user.name : "Student"}!
          </h1>
        </header>
        <section id="cv">
          <h2>{ cv && !upload ? "Your CV" : "Upload CV"}</h2>
          <FileUploadDialog
            accept="application/pdf"
            className="cv"
            multiple={false}
            onUpload={this.uploadCV}
            ref={n => this.fileRef = n}
          />
        </section>
      </div>
    );
  }
}
