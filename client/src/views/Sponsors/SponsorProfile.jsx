import React from 'react';
import ReactTable from 'react-table';
import 'react-table/react-table.css';
import {fetchCVs, downloadCV} from './actions';
import Header from 'Components/Header';

class Download extends React.Component {
  download = async () => {
    const {id} = this.props;
    try {
      await downloadCV(id)
    } catch (e) {
      console.log("download cvs", e)
    }
  }

  render = () => <button onClick={this.download}>Download</button>
}

const columns = [
  {
    Header: 'Student\'s Name',
    accessor: 'name',
  },
  {
    Header: '',
    accessor: 'id',
    Cell: row => <Download id={row.value} />,
    width: 120
  },
];

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
      const cvs = await fetchCVs()
      if (!cvs) return
      console.log(cvs)
      this.setState({cvs})
    } catch (e) {
      console.log("fetch cvs", e)
    }
  }

  render() {
    let {cvs} = this.state;
    return (
      <div>
        <Header logout="/auth/sponsors/logout" />
        <div id="main">
          <h1>Student CVs</h1>
          <ReactTable
            className="-striped"
            data={cvs}
            columns={columns}
            showPagination={cvs.length > 10}
            showPageSizeOptions={false}
          />
        </div>
      </div>
    );
  }
}
