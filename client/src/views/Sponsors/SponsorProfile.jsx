import React from 'react';
import ReactTable from 'react-table';
import 'react-table/react-table.css';
import {fetchCVs, downloadCV} from './actions';

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
    Header: 'CV',
    accessor: 'name',
  },
  {
    Header: 'Download',
    accessor: 'id',
    Cell: row => <Download id={row.value} />
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
      this.setState({cvs})
    } catch (e) {
      console.log("fetch cvs", e)
    }
  }

  render() {
    let {cvs} = this.state;
    return (
      <div>
        <h1>Hello, Sponsor</h1>
        <ReactTable
          data={cvs}
          columns={columns}
          showPagination={false}
          showPageSizeOptions={false}
        />
      </div>
    );
  }
}
