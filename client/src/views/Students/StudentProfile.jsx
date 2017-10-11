import React from 'react';
import FileUploadDialog from 'Components/FileUploadDialog';
import Header from 'Components/Header';
import SponsorPanel from './panels/SponsorPanel';
import request from 'superagent';
import { getJWTHeader } from '../../jwt';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import { fetchUser, fetchCV } from './actions';

export default class StudentProfile extends React.Component {
  constructor(props) {
    super(props)

    this.state = {}
  }

  componentDidMount = async () => {
    await this.getUser()
    await this.getCV()
  }

  getUser = async () => {
    try {
      const user = await fetchUser()
      this.setState({user})
    } catch (e) {
      console.log("fetch user", e)
    }
  }

  getCV = async () => {
    try {
      const cv = await fetchCV()
      if (!cv) return
      this.setState({cv, upload: false})
      this.fileRef.updateFile([cv])
    } catch (e) {
      console.log("fetch cv", e)
    }
  }

  uploadCV = async (files, progress) => {
    if (files.length > 1) {
      throw new Error("Expecting exactly 1 CV")
    }

    let headers = await getJWTHeader();

    try {
      let data = await request
        .post('/api/students/cv')
        .set(headers)
        .attach('cv', files[0]).
        on('progress', event => {
          progress(event.percent);
        });

      this.setState({upload: false})
    } catch (e) {
      console.log(e)
      throw new Error("Failed to upload file")
    }
  }

  render() {
    let {user, cv, upload} = this.state;

    return (
      <div>
        <Header name={user && user.name} logout="/auth/students/logout"/>
        <div className="student-page">
          <Tabs className="tabs underline">
            <TabList>
              <Tab>Profile</Tab>
              <Tab>Sponsors</Tab>
            </TabList>

            <TabPanel>
              <section id="profile">
                <div id="cv">
                  <h2>{ cv && !upload ? "Your CV" : "Upload CV"}</h2>
                  <FileUploadDialog
                    accept="application/pdf"
                    className="cv"
                    multiple={false}
                    files={cv ? [cv] : []}
                    onUpload={this.uploadCV}
                    ref={n => this.fileRef = n}
                  />
                </div>
              </section>
            </TabPanel>
            <TabPanel>
              <SponsorPanel />
            </TabPanel>
          </Tabs>
        </div>
      </div>
    );
  }
}
