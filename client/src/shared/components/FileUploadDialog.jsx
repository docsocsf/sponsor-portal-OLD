import React from 'react';
import Dropzone from 'react-dropzone';
import mime from 'mime-db';

export default class FileUploadDialog extends React.Component {
  constructor() {
    super()

    this.state = {}
    this.getState = this.getState.bind(this);
    this.fileUpload = this.fileUpload.bind(this);
  }

  fileUpload(accepted, rejected){
    console.log("pls be here")
    if (accepted.length >= 1) {
      this.setState({file: accepted[0], error: undefined})
    } else if (rejected.length >= 0) {
      this.setState({file: undefined, error: rejected[0]})
    }
  }

  getState() {
    if (this.state.file) {
      return "accept"
    } else if (this.state.error) {
      return "reject"
    } else {
      return "empty"
    }
  }

  render() {
    let { accept, className } = this.props;
    let { file, error } = this.state;
    let state = this.getState();

    return (
          <Dropzone
            {...this.props}
            onDrop={this.fileUpload}
            className={`${state} ${className}`}
            ref={(node) => { this.dropzoneRef = node }}
          >
            <SuccessState file={file}/>
            <ErrorState error={error} mime={accept}/>
            <EmptyState file={file} open={e => this.dropzoneRef.open()} />
          </Dropzone>
    );
  }
}

function SuccessState(props) {
  if (!props.file) {
    return null;
  }
  return (<span>{props.file.name}</span>);
}

function ErrorState(props) {
  if (!props.error) {
    return null;
  }

  let types = mime[props.mime].extensions.map((ext, i) =>
    <span key={i} className="extension">{ i > 0 && ", "}{ext}</span>
  );

  return (
    <div>
      <p className="error-message">
        File must be of type: {types}.
      </p>
    </div>
  );
}

function EmptyState(props) {
  if (props.file) {
    return null;
  }

  return (
    <p>
      Drag a file here or <a href="#" onClick={e => e.preventDefault()}>browse</a>.
    </p>
  );
}
