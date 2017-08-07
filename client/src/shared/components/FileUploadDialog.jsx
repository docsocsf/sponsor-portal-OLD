import React from 'react';
import Dropzone from 'react-dropzone';
import mime from 'mime-db';
import Progress from 'Components/Progress';

const initialState = {
  loading: false,
  percent: 0,
  files: [],
  errors: [],
};

export default class FileUploadDialog extends React.Component {
  constructor() {
    super()

    this.state = initialState;
    this.getState = this.getState.bind(this);
    this.fileUpload = this.fileUpload.bind(this);
  }

  async fileUpload(accepted, rejected){
    if (accepted.length >= 1) {
      try {
        this.setState({loading: true})
        let data = await this.props.onUpload(accepted, percent =>
          this.setState({percent}));
        this.setState({...initialState, files: accepted})
      } catch (e) {
        this.setState({...initialState, errors: [...this.state.errors, e]})
      }
    } else if (rejected.length >= 0) {
      let types = mime[this.props.accept].extensions.map((ext, i) =>
        <span key={i} className="extension">{ i > 0 && ", "}{ext}</span>
      );
      let e = new Error(`File must be of type: ${types}`)
      this.setState({...initialState, errors: [...this.state.errors, e]})
    }
  }

  getState() {
    let {files, errors, loading} = this.state;
    if (loading) {
      return "loading";
    } else if (files.length > 0) {
      return "accept";
    } else if (errors.length > 0) {
      return "reject"
    } else {
      return "empty"
    }
  }

  render() {
    let { accept, className, multiple } = this.props;
    let { files, errors, loading, percent } = this.state;
    let state = this.getState();

    return (
          <Dropzone
            accept={accept}
            multiple={multiple}
            onDrop={this.fileUpload}
            className={`${state} ${className}`}
            ref={(node) => { this.dropzoneRef = node }}
          >
            <SuccessState files={files} loading={loading}/>
            <ErrorState errors={errors} mime={accept}/>
            <LoadingState loading={loading} percent={percent} />
            <EmptyState files={files} loading={loading} open={e => this.dropzoneRef.open()} />
          </Dropzone>
    );
  }
}

function SuccessState(props) {
  if (!props.files.length > 0 || props.loading) {
    return null;
  }

  return (
    <p>
      {props.files.map((file, i) =>
        <span key={i}>{ i > 0 && (<br/>)}{file.name}</span>
      )}
    </p>
  );
}

function ErrorState(props) {
  if (!props.errors.length > 0) {
    return null;
  }

  let errors = props.errors.map((e, i) =>
    <span key={i} className="error-message">{ i > 0 && (<br/>)}{e.message}</span>
  );

  return (
    <p>
      {errors}
    </p>
  );
}

function LoadingState(props) {
  if (!props.loading) {
    return null;
  }

  return (
    <div>
      <p>
        Loading...
      </p>
      <Progress percent={props.percent} />
    </div>
  );
}

function EmptyState(props) {
  if (props.files.length > 0 || props.loading) {
    return null;
  }

  return (
    <p>
      Drag a file here or <a href="#" onClick={e => e.preventDefault()}>browse</a>.
    </p>
  );
}
